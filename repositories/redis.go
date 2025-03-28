package repositories

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Cli *redis.Client
}

func NewRedisClient() (*RedisClient, error) {

	host := "localhost"
	if os.Getenv("REDIS_HOST") != "" {
		host = os.Getenv("REDIS_HOST")
	}
	port := "6379"
	if os.Getenv("REDIS_PORT") != "" {
		port = os.Getenv("REDIS_PORT")
	}
	db := 0
	dbValue, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err == nil {
		db = dbValue
	}

	url := fmt.Sprintf("%s:%s", host, port)

	return &RedisClient{
		Cli: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: "",
			DB:       db,
		}),
	}, nil
}

func (r *RedisClient) GetUser(ctx context.Context, key string) (*User, error) {
	value, err := r.Cli.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("failed to get value: %w", err)
	}
	if value == "" {
		return nil, nil
	}

	value = strings.Trim(string(value), "\"")

	decodedValue, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Base64 value: %w", err)
	}

	var user *User
	err = json.Unmarshal(decodedValue, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize JSON value: %w", err)
	}

	return user, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	return r.Cli.Set(ctx, key, valueBytes, 0).Err()
}

func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.Cli.Del(ctx, key).Err()
}

func (r *RedisClient) Ping(ctx context.Context) error {
	return r.Cli.Ping(ctx).Err()
}
