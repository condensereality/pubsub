package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/vardius/golog"
	"io"
	"time"

	"github.com/vardius/pubsub/v2/proto"
)

type server struct {
	proto.UnimplementedPubSubServer
	bus    Service
	logger golog.Logger
}

type timeout struct {
	topic string
	duration time.Duration
}

type message struct {
	Payload Payload
	Deadline *time.Duration
}

func (t timeout) Error() string {
	return fmt.Sprintf("subscriber to %s timed out after %s", t.topic, t.duration)
}

// newServer returns new pub/sub server object
func newServer(bus Service, logger golog.Logger) proto.PubSubServer {
	return &server{proto.UnimplementedPubSubServer{},bus, logger}
}

// Publish publishes message payload to all topic handlers
func (s *server) Publish(ctx context.Context, r *proto.PublishRequest) (*empty.Empty, error) {
	payload := r.GetPayload()
	s.logger.Debug(ctx, "Publish: %s %d bytes", r.GetTopic(), len(payload))

	s.bus.Publish(ctx, r.GetTopic(), &payload)

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	return new(empty.Empty), nil
}

// Subscribe subscribes to a topic,
// will unsubscribe when stream.Send returns error
func (s *server) Subscribe(r *proto.SubscribeRequest, stream proto.PubSub_SubscribeServer) error {
	handlerErrCh := make(chan error)
	handler := func(ctx context.Context, payload Payload) {
		var err error
		handlerCtx := stream.Context()
		s.logger.Debug(handlerCtx, "Subscribe: %s %d bytes", r.GetTopic(), len(*payload))

		sendErrCh := make(chan error)

		resp := &proto.SubscribeResponse{
			Payload: *payload,
		}

		sendTimeout := 1 * time.Second

		go func() {
			sendErrCh <- stream.Send(resp)
		}()

		select {
		case <-time.After(sendTimeout):
			err = timeout{duration: sendTimeout, topic: r.GetTopic()}
			s.logger.Debug(handlerCtx, "timeout")
		case <-stream.Context().Done():
			err = stream.Context().Err()
			s.logger.Debug(handlerCtx, fmt.Sprintf("context done: %s", err))
		case err = <-sendErrCh:
			if err == nil {
				s.logger.Debug(handlerCtx, "send completed successfully")
			} else {
				s.logger.Debug(handlerCtx, "send completed with error: %s", err.Error())
			}
		}

		if err != nil {
			handlerErrCh <- err
		}
	}

	s.logger.Info(stream.Context(), "Subscribe: %s", r.GetTopic())

	if err := s.bus.Subscribe(r.GetTopic(), handler); err != nil {
		return err
	}

	handlerErr := <-handlerErrCh

	if err := s.bus.Unsubscribe(r.GetTopic(), handler); err != nil {
		return err
	}

	if handlerErr == io.EOF {
		s.logger.Info(stream.Context(), "Unsubscribe: %s - Stream closed, no more input is available", r.GetTopic())

		return nil
	}

	s.logger.Info(stream.Context(), "Unsubscribe: %s - %s", r.GetTopic(), handlerErr.Error())

	return handlerErr
}
