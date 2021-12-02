package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var clientMap map[string]*redis.Client

type Config struct {
	Name     string
	Host     string
	Port     int
	Password string
	DB       int
}

func Startup(configs []Config) {
	clientMap = make(map[string]*redis.Client, len(configs))
	for _, v := range configs {
		client := redis.NewClient(&redis.Options{
			Addr:     v.Host + ":" + strconv.Itoa(v.Port),
			Password: v.Password,
			DB:       v.DB,
		})
		client.AddHook(hook{})
		clientMap[v.Name] = client
	}
}

func Shutdown() error {
	for _, v := range clientMap {
		err := v.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func GetClient(name string, ctx context.Context) (*redis.Client, error) {
	if client, ok := clientMap[name]; ok {
		return client.WithContext(ctx), nil
	}
	return nil, errors.New("redis client not found")
}
