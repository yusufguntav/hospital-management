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
			&entities.Job{},
			&entities.Title{},
			&entities.Employee{},
			&entities.Clinic{},
			&entities.ClinicAndHospital{},
		)
	})
	addDatas()
}
func addDatas() {
	addDataIfNotExists(entities.City{}, "./pkg/data/city.json", "Adding cities to db")
	addDataIfNotExists(entities.District{}, "./pkg/data/districts.json", "Adding districts to db")
	addDataIfNotExists(entities.Job{}, "./pkg/data/job.json", "Adding jobs to db")
	addDataIfNotExists(entities.Title{}, "./pkg/data/titles.json", "Adding titles to db")
	addDataIfNotExists(entities.Clinic{}, "./pkg/data/clinic.json", "Adding clinics to db")
}

func addDataIfNotExists(modelType any, filePath string, logMessage string) {
	var count int64
	if err := DBClient().Model(modelType).Count(&count).Error; err != nil {
		log.Print("Error:", err)
	}
	if count == 0 {
		log.Println(logMessage)
		data := []interface{}{}
		if err := utils.ReadJsonFile(filePath, &data); err != nil {
			log.Print("Error:", err)
			return
		}
		for _, item := range data {
			if err := DBClient().Model(modelType).Create(item).Error; err != nil {
				log.Print("Error:", err)
			}

		}
	}
}

func DBClient() *gorm.DB {
	if db == nil {
		log.Panic("Postgres is not initialized. Call InitDB first.")
	}
	return db
}
