package databases

import (
	"fmt"
	"log"
	"sync"
	"tutorial/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstance *postgresDatabase
	once                     sync.Once
)

func NewPostgresDatabase(config *config.Database) Database {

	once.Do(func() {
		dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

		db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Printf("Connected to database %s", config.DBName)

		postgresDatabaseInstance = &postgresDatabase{db}

		postgresDatabaseInstance.AutoMigrate()
	})
	return postgresDatabaseInstance
}

func (db *postgresDatabase) Connect() *gorm.DB {
	return postgresDatabaseInstance.DB
}
