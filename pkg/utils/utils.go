package utils

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ReadJsonFile(path string, items interface{}) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteValue, &items); err != nil {
		return err
	}

	return nil
}

func Pagination(item interface{}, pageNumber int, db *gorm.DB, c context.Context, query interface{}, args ...interface{}) (int, error) {
	limit := 10
	offset := 0

	var totalCount int64
	if err := db.WithContext(c).Model(item).Where(query, args...).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	if pageNumber > totalPages || pageNumber <= 0 {
		return 0, errors.New("invalid page number")
	}

	// Check if pageNumber is provided and valid
	if pageNumber > 0 {
		offset = (pageNumber - 1) * limit
	}

	// Get items with pagination
	if err := db.WithContext(c).Limit(limit).Offset(offset).Where(query, args...).Find(item).Error; err != nil {
		return 0, err
	}
	return totalPages, nil
}
