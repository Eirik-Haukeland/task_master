package say

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strings"
)

func StoreValue(valueName string, valueData string) error {
	var err error
	if strings.TrimSpace(valueName) == "" {
		err = errors.New("valueName is empty")
		return err
	} else if strings.TrimSpace(valueData) == "" {
		err = errors.New("valueData is empty")
		return err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	err = client.SetNX(ctx, valueName, valueData, redis.KeepTTL).Err()

	return err
}

func GetValue(valueName string) (error, string) {
	var err error
	if strings.TrimSpace(valueName) == "" {
		err = errors.New("valueName is empty")
		return err, ""
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	val, err := client.Get(ctx, valueName).Result()
	if err != nil {
		return err, ""
	}

	return err, val
}

func Hello(thing string) {
	if thing == "" {
		thing = "world"
	}

	fmt.Printf("hello, %s\n", thing)
}
