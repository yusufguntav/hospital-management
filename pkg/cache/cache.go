package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/yusufguntav/hospital-management/pkg/config"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"gorm.io/gorm"
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
func Get(ctx context.Context, key string, data interface{}) error {
	result, err := client.WithContext(ctx).Get(key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(result), &data); err != nil {
		return err
	}

	return nil
}

// TODO: GetDistrcitsAndCities ve GetJobsAndTitles benzer yapıdalar fonksiyona dök
func GetDistrictsAndCities(c context.Context, db *gorm.DB) (*[]entities.District, *[]entities.City, error) {

	cacheDistricts := []entities.District{}
	Get(c, "districts", &cacheDistricts)

	if len(cacheDistricts) == 0 || cacheDistricts == nil {
		log.Println("Adding districts to cache")
		var districts []entities.District
		db.WithContext(c).Find(&districts)
		Set(c, "districts", districts, 0)
		cacheDistricts = districts
	}

	cacheCities := []entities.City{}
	Get(c, "cities", &cacheCities)

	if len(cacheCities) == 0 || cacheCities == nil {
		log.Println("Adding cities to cache")
		var cities []entities.City
		db.WithContext(c).Find(&cities)
		Set(c, "cities", cities, 0)
		cacheCities = cities
	}

	return &cacheDistricts, &cacheCities, nil
}

func GetJobs(c context.Context, db *gorm.DB) (*[]entities.Job, error) {

	cacheJobs := []entities.Job{}
	Get(c, "jobs", &cacheJobs)

	if len(cacheJobs) == 0 || cacheJobs == nil {
		log.Println("Adding jobs to cache")
		var jobs []entities.Job
		db.WithContext(c).Find(&jobs)
		Set(c, "jobs", jobs, 0)
		cacheJobs = jobs
	}

	return &cacheJobs, nil
}
func GetTitles(c context.Context, db *gorm.DB) (*[]entities.Title, error) {

	cacheTitles := []entities.Title{}
	Get(c, "titles", &cacheTitles)

	if len(cacheTitles) == 0 || cacheTitles == nil {
		log.Println("Adding cities to cache")
		var titles []entities.Title
		db.WithContext(c).Find(&titles)
		Set(c, "titles", titles, 0)
		cacheTitles = titles
	}

	return &cacheTitles, nil
}

func GetClinics(c context.Context, db *gorm.DB) (*[]entities.Clinic, error) {

	cacheClinics := []entities.Clinic{}
	Get(c, "clinics", &cacheClinics)

	if len(cacheClinics) == 0 || cacheClinics == nil {
		log.Println("Adding clinics to cache")
		var clinics []entities.Clinic
		db.WithContext(c).Find(&clinics)
		Set(c, "clinics", clinics, 0)
		cacheClinics = clinics
	}

	return &cacheClinics, nil
}
