package main

import (
	"context"
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

	for i := 0; i < 100; i++ {
		if err := client.Publish(ctx, "topic1", i).Err(); err != nil {
			log.Fatal(err)
		}
	}

}
