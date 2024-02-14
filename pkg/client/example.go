package client

import (
	"context"
	"log"
	"time"

	pb "github.com/converged-computing/rainbow/pkg/api/v1"
	"github.com/converged-computing/rainbow/pkg/types"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	anypb "google.golang.org/protobuf/types/known/anypb"
	ts "google.golang.org/protobuf/types/known/timestamppb"
	wrapper "google.golang.org/protobuf/types/known/wrapperspb"
)

// Stream sends a message to the server and receives the response.
func (c *RainbowClient) Stream(ctx context.Context, it types.MessageIterator) error {
	if it == nil {
		return errors.New("message provider is required")
	}

	if !c.Connected() {
		return errors.New("client is not connected")
	}

	// Stream example
	stream, err := c.service.Stream(ctx)
	if err != nil {
		return errors.Wrap(err, "could not create stream")
	}

	// Send messages to the stream
	for it.HasNext() {
		// check for context errors before each iteration
		if err := ctx.Err(); err != nil {
			return errors.Wrap(err, "context")
		}

		if err := stream.Context().Err(); err != nil {
			return errors.Wrap(err, "stream context")
		}

		// create message with wrapper
		a, err := anypb.New(wrapper.String(it.Next()))
		if err != nil {
			return errors.Wrap(err, "unable to create message")
		}

		// send the message
		if err := stream.Send(&pb.Request{
			Content: &pb.Content{
				Id:   uuid.New().String(),
				Data: a,
			},
			Sent: ts.Now(),
		}); err != nil {
			return errors.Wrap(err, "failed to send a message")
		}

		// capture the response
		reply, err := stream.Recv()
		if err != nil {
			return errors.Wrap(err, "error processing response")
		}

		log.Printf("stream response: %s", reply.GetProcessingDetails())
	}

	if err := stream.CloseSend(); err != nil {
		return errors.Wrap(err, "failed to close stream")
	}
	return nil
}

// Scalar sends a message to the server and returns the response.
func (c *RainbowClient) Serial(ctx context.Context, message string) (string, error) {
	if message == "" {
		return "", errors.New("message is required")
	}

	if !c.Connected() {
		return "", errors.New("client is not connected")
	}

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// create message with wrapper
	a, err := anypb.New(wrapper.String(message))
	if err != nil {
		return "", errors.Wrap(err, "unable to create message")
	}

	// Scalar example
	r, err := c.service.Serial(ctx, &pb.Request{
		Content: &pb.Content{
			Id:   uuid.New().String(),
			Data: a,
		},
		Sent: ts.Now(),
	})
	if err != nil {
		return "", errors.Wrap(err, "could not send scalar message")
	}

	return r.GetProcessingDetails(), nil
}