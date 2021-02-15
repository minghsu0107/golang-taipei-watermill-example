# Watermill Example
This is a tutorial of building event-driven application in Golang using [Watermill](https://github.com/ThreeDotsLabs/watermill), a fantastic Go library for working efficiently with message streams.
## Overview
![](./assets/overview.png)

We run a publisher in the background, publishing messages to the message broker on topic `incoming_topic` every 3 seconds. At the same time, `helloHandler` and `incomingTopicHandler` subscribe on topic `incoming_topic`; `outgoingTopicHandler` subscribes on `outgoing_topic`. Once `helloHandler` receives a message, it will then publish another greeting message to topic `outgoing_topic`, which will be recived by `outgoingTopicHandler`.

In `transport.go`, we decouple our service logics from the broker implementation. That is, we are able to switch to any message broker without changing existing code as long as the broker client implements `message.Publisher` and `message.Subscriber` interface.
## Usage
```bash
docker-compose up
```
We can switch to GoChannel Pub/Sub by simply setting `WATERMILL_PUBSUB_TYPE` to empty string.