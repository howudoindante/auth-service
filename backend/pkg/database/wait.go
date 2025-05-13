package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func waitForDB(dsn string, maxAttempts int) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := 0; i < maxAttempts; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		time.Sleep(2 * time.Second)
	}
	return nil, err
}
