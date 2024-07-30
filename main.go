package main

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func main() {
	storage := storage{}
	cachedStorage := &cachedStorage{
		storage: storage,
		client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}

	app := app{
		storage: cachedStorage,
	}

	p, err := app.getMyProfile(token{userID: 1})

	log.Info().Msgf("Got user by id 1: %v, %v", p, err)
	p, err = app.getMyProfile(token{userID: 1})

	log.Info().Msgf("Got user by id 1: %v, %v", p, err)
	p, err = app.getMyProfile(token{userID: 1})

	log.Info().Msgf("Got user by id 1: %v, %v", p, err)
}

func example() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	res, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal().Msgf("Pinging redis : %v", err)
	}
	log.Info().Msgf("Pinged: %v", res)

	res, err = client.Set(ctx, "name", "John", 0).Result()
	if err != nil {
		log.Fatal().Msgf("Setting name: %v", err)
	}

	log.Info().Msgf("Set: %v", res)

	res, err = client.Get(ctx, "name").Result()
	if err != nil {
		log.Fatal().Msgf("Getting name: %v", err)
	}

	log.Info().Msgf("Getting name: %v", res)
}
