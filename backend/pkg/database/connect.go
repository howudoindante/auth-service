package database

import (
	"auth/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabaseWithSchema(cfg config.Database) (*gorm.DB, error) {
	// Подключение к postgres без указания конкретной БД
	adminDSN := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s sslmode=%s dbname=postgres",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.SSLMode,
	)
	adminDB, err := waitForDB(adminDSN, cfg.ConnectingAttempts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect after %d attempts: %w", cfg.ConnectingAttempts, err)
	}

	// Создание БД, если не существует
	createDBQuery := fmt.Sprintf(`CREATE DATABASE "%s";`, cfg.Name)
	err = adminDB.Exec(fmt.Sprintf(`SELECT 1 FROM pg_database WHERE datname = '%s'`, cfg.Name)).Error
	if err == nil {
		if err := adminDB.Exec(createDBQuery).Error; err != nil {
			// Возможно БД уже есть, игнорируем ошибку
			log.Printf("warning: could not create DB %s: %v", cfg.Name, err)
		}
	}

	// Подключение к нужной БД
	dbDSN := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name, cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", cfg.Name, err)
	}

	if cfg.Scheme != "" {
		if err := db.Exec(fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS "%s"`, cfg.Scheme)).Error; err != nil {
			fmt.Printf("warning: schema creation error: %v", err)
		}
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto`)

	// Возвращаем подключение с search_path
	finalDSN := fmt.Sprintf(
		"%s search_path=%s",
		dbDSN, cfg.Scheme,
	)
	dbWithSchema, err := gorm.Open(postgres.Open(finalDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect with schema: %w", err)
	}

	return dbWithSchema, nil
}
