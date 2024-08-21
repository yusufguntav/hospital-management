package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/yusufguntav/hospital-management/pkg/config"
)

var client *redis.Client

// TODO kodlar tek satıra indireilcek err handler
func InitRedis(redisConf config.Redis) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConf.Host + ":" + redisConf.Port,
		Password: redisConf.Pass,
	})
}

func IsExist(ctx context.Context, key string) bool {
	err := client.WithContext(ctx).Get(key).Err()
	if err != nil {
		log.Println("cache error: ", err)
		return false
	}
	return true

}

func Set(ctx context.Context, key string, list interface{}, ex int64) error {
	jsondata, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = client.WithContext(ctx).Set(key, jsondata, time.Duration(ex*int64(time.Second))).Err()
	if err != nil {
		return err
	}
	return nil
}
func Get(ctx context.Context, key string) (interface{}, error) {
	result, err := client.WithContext(ctx).Get(key).Result()
	if err != nil {
		return nil, err
	}

	var data interface{}
	if err := json.Unmarshal([]byte(result), &data); err != nil {
		return nil, err
	}

	return data, nil
}
