package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type cachedStorage struct {
	client  *redis.Client
	storage storage
}

func (s *cachedStorage) getUserByID(id int) (*User, error) {

	key := fmt.Sprintf("user: %v", id)
	ctx := context.Background()

	res, err := s.client.Get(ctx, key).Result()

	if err != nil && err.Error() == "redis: nil" { //todo refactor change to better check
		user, err := s.storage.getUserByID(id)

		if err != nil {
			return nil, err
		}

		v, err := json.Marshal(user)
		if err != nil {
			return nil, fmt.Errorf("marshaling for cache: %w", err)
		}

		res, err = s.client.SetEx(ctx, key, v, 15*time.Second).Result()

		if err != nil {
			log.Error().Msgf("Failed to cache : %v", err)
			return user, nil
		}

		return user, nil
	}

	if err != nil {
		return nil, fmt.Errorf("getting cached user: %w", err)
	}

	var u User
	if err := json.Unmarshal([]byte(res), &u); err != nil {
		return nil, fmt.Errorf("unmarshaling cached user: %w", err)

	}

	return &u, nil

}
