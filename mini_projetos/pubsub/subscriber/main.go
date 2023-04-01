package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
)

//const PROJECT = "jp-estudos"
//const MY_TOPIC = "my-topic"

const PROJECT = "mock-local"
const MY_TOPIC = "topic-local"


func main() {
	ctx := context.Background()

	//criar um cliente pubsub
	cl, err := pubsub.NewClient(ctx, PROJECT)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Printf("=> client criado\n")

	topic := getTopic(ctx, cl, MY_TOPIC)
	if topic == nil {
		log.Fatalf("Topic not found: %s", MY_TOPIC)
	}

	fmt.Printf("=> topic encontrado: %s \n", topic.String())
	//cria o subscriber
	 sub, err := createSubscriber(ctx, cl, "sub-local", topic)
	 if err != nil {
	 	log.Fatalf("Failed create subscriber: %s", err)
	 }

	//sub := cl.Subscription("my-topic-sub")
	fmt.Printf("=> subscriber: %s \n", sub.String())

	//realizar o pull das msg
	err = pullMgs(ctx, cl, sub, topic)
	if err != nil {
		log.Fatalf("Failed to pull msg: %v", err)
	}
}

func pullMgs(ctx context.Context, client *pubsub.Client, sub *pubsub.Subscription, topic *pubsub.Topic) error {
	var mx sync.Mutex
	received := 0

	cctx, cancel := context.WithCancel(ctx)
	err := sub.Receive(cctx, func(c context.Context, m *pubsub.Message) {
		mx.Lock()
		defer mx.Unlock()

		received++
		if received >= 10 {
			cancel()
			m.Nack()
			return
		}

		fmt.Printf("=> Message:\n\t %q\n", m.Data)
		m.Ack()
	})

	if err != nil {
		return err
	}
	return nil
}

func getTopic(ctx context.Context, c *pubsub.Client, topicName string) *pubsub.Topic {
	//busca o topico
	t := c.Topic(topicName)
	exist, err := t.Exists(ctx)

	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
		return nil
	}

	if exist {
		return t
	}
	return nil
}

func createSubscriber(ctx context.Context, c *pubsub.Client, name string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	sub, err := c.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	fmt.Printf("Created subscription: %v\n", sub)
	return sub, nil
}
