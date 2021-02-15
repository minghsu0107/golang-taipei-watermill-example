package main

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type helloHandler struct{}

type incomingTopicHandler struct{}

type outgoingTopicHandler struct{}

// Handler receives the published message and publish to my_outgoing topic
// if there is error in this handler or router fails to publish to my_outgoing, the msg will be NACKed
func (h helloHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	fmt.Printf(
		"\n> helloHandler received message: %s\n> %s\n> metadata: %v\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)

	msg = message.NewMessage(watermill.NewUUID(), []byte("greet from helloHandler"))
	return message.Messages{msg}, nil
}

// HandlerWithoutPublish receives and handles subscribe message without publishing it again
func (h incomingTopicHandler) HandlerWithoutPublish(msg *message.Message) error {
	fmt.Printf(
		"\n> incomingTopicHandler received message: %s\n> %s\n> metadata: %v\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}

// HandlerWithoutPublish receives and handles subscribe message without publishing it again
func (h outgoingTopicHandler) HandlerWithoutPublish(msg *message.Message) error {
	fmt.Printf(
		"\n> outgoingTopicHandler received message: %s\n> %s\n> metadata: %v\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}
