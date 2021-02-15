package main

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	stan "github.com/nats-io/stan.go"
)

var marshaler nats.GobMarshaler

// NewGoChannel returns a go channel pubsub
func NewGoChannel(logger watermill.LoggerAdapter) *gochannel.GoChannel {
	return gochannel.NewGoChannel(gochannel.Config{
		// If persistent is set to true, when subscriber subscribes to the topic,
		// it will receive all previously produced messages
		Persistent: true,
	}, logger)
}

// NewNATSPublisher returns a NATS publisher
func NewNATSPublisher(logger watermill.LoggerAdapter, clusterID, natsURL string) (message.Publisher, error) {
	return nats.NewStreamingPublisher(
		nats.StreamingPublisherConfig{
			ClusterID: clusterID,
			ClientID:  watermill.NewShortUUID(),
			StanOptions: []stan.Option{
				stan.NatsURL(natsURL),
			},
			Marshaler: marshaler,
		},
		logger,
	)
}

// NewNATSSubscriber returns a NATS subscriber
func NewNATSSubscriber(logger watermill.LoggerAdapter, clusterID, clientID, natsURL string) (message.Subscriber, error) {
	return nats.NewStreamingSubscriber(
		nats.StreamingSubscriberConfig{
			ClusterID: clusterID,
			ClientID:  clientID,
			StanOptions: []stan.Option{
				stan.NatsURL("nats://nats-streaming:4222"),
			},
			Unmarshaler: marshaler,
		},
		logger,
	)
}
