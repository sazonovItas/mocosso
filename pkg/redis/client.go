package cacheredis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	// pingTimeout is default timeout to ping redis server.
	pingTimeout = 5 * time.Second
)

func MustConnect(redisURL string) *redis.Client {
	cli, err := Connect(redisURL)
	if err != nil {
		panic(err)
	}

	return cli
}

func Connect(redisURL string) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	return ConnectContext(ctx, redisURL)
}

func ConnectContext(ctx context.Context, redisURL string) (*redis.Client, error) {
	const op = "pkg.redis.ConnectContext"

	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	cli := redis.NewClient(opts)
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return cli, nil
}
