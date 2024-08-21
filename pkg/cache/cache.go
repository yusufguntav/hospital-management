package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/yusufguntav/hospital-management/pkg/config"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/utils"
)

var client *redis.Client

func InitRedis(redisConf config.Redis) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConf.Host + ":" + redisConf.Port,
		Password: redisConf.Pass,
	})
}

func IsExist(ctx context.Context, key string) bool {
	if err := client.WithContext(ctx).Get(key).Err(); err != nil {
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
	if err := client.WithContext(ctx).Set(key, jsondata, time.Duration(ex*int64(time.Second))).Err(); err != nil {
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

func AddDistrictsAndCities(c context.Context) error {
	var districts []entities.District
	utils.ReadJsonFile("./pkg/data/districts.json", &districts)
	if err := Set(c, "districts", districts, 0); err != nil {
		return err
	}

	var cities []entities.City
	utils.ReadJsonFile("./pkg/data/city.json", &cities)
	if err := Set(c, "cities", cities, 0); err != nil {
		return err
	}

	return nil
}
