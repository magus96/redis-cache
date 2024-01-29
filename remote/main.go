package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:5432",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	sub := client.Subscribe(ctx, "topic1")
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", msg.Payload)
	}
}
