package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/yusufguntav/hospital-management/pkg/config"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	err         error
	client_once sync.Once
)

func InitDB(dbc config.Database) {
	client_once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbc.Host, dbc.Port, dbc.User, dbc.Pass, dbc.Name)
		db, err = gorm.Open(
			postgres.New(
				postgres.Config{
					DSN:                  dsn,
					PreferSimpleProtocol: true,
				},
			),
		)
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(
			&entities.Hospital{},
			&entities.User{},
			&entities.City{},
			&entities.District{},
		)
	})
	addDatas()
}
func addDatas() {
	var count int64
	DBClient().Model(&entities.City{}).Count(&count)
	if count == 0 {
		log.Println("Adding cities")
		cities := []entities.City{}
		if err := utils.ReadJsonFile("./pkg/data/city.json", &cities); err != nil {
			log.Print("Error:", err)
		}
		for _, city := range cities {
			DBClient().Create(&city)
		}
	}
	count = 0

	DBClient().Model(&entities.District{}).Count(&count)
	if count == 0 {
		log.Println("Adding districts")
		districts := []entities.District{}
		if err := utils.ReadJsonFile("./pkg/data/districts.json", &districts); err != nil {
			log.Print("Error:", err)
		}
		for _, district := range districts {
			DBClient().Create(&district)
		}
	}
}

func DBClient() *gorm.DB {
	if db == nil {
		log.Panic("Postgres is not initialized. Call InitDB first.")
	}
	return db
}
