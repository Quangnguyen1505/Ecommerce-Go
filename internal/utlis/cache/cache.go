package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ntquang/ecommerce/global"
	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {
	data, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		return err
	} else if err == redis.Nil {
		return fmt.Errorf("key %s not found", key)
	}

	// convert json to object
	if err = json.Unmarshal([]byte(data), &obj); err != nil {
		return fmt.Errorf("failed convert json to object")
	}
	return nil
}
