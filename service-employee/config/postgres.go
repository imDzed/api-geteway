package config

import (
	"sync"

	"github.com/gofiber/fiber/v2/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once       sync.Once
	postgresDB *gorm.DB
)

func NewPostgresDatabase() {
	once.Do(func() {
		dsn := "host=localhost user=admin password=admin dbname=service_employees port=5432 sslmode=disable"
		// dsn := "host=localhost user=admin password=admin dbname=service_employees port=9920 sslmode=disable TimeZone=Asia/Shanghai"

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Error("terjadi kesalahan, error:", err.Error())
		}

		postgresDB = db
	})
}

func GetPostgresDatabase() *gorm.DB {
	return postgresDB
}
