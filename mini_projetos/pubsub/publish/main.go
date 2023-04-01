package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

const PROJECT = "jp-estudos"
// const MY_TOPIC = "my-topic"

//const PROJECT = "mock-local"
const MY_TOPIC = "teste"

func main() {
	ctx := context.Background()

	// mensagem capturada como argumento
	msg := os.Args[1]

	// criar um client do pubsub
	cl, err := pubsub.NewClient(ctx, PROJECT)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// criando um topico
	topic := createTopicIfNotExist(ctx, cl, MY_TOPIC)

	//Publicando uma mensagem no topico
	publishNewMsg(ctx, cl, topic, msg)

}

func createTopicIfNotExist(ctx context.Context, c *pubsub.Client, topic string) *pubsub.Topic {
	// verifica se o topic já existe
	t := c.Topic(topic)
	exist, err := t.Exists(ctx)

	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}

	if exist {
		return t
	}

	//cria um novo topico
	t, err = c.CreateTopic(ctx, topic)
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}

	return t
}

func publishNewMsg(ctx context.Context, c *pubsub.Client, topic *pubsub.Topic, msg string) error {
	// realiza a publicação da mensagem
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := result.Get(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("=> Publish a message! ID: %v\n", id)
	return nil
}
