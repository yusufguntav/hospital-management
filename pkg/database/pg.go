package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/yusufguntav/hospital-management/pkg/config"
	"github.com/yusufguntav/hospital-management/pkg/entities"
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
		)
	})
}

func DBClient() *gorm.DB {
	if db == nil {
		log.Panic("Postgres is not initialized. Call InitDB first.")
	}
	return db
}
