package db

import (
	"fmt"

	"Library/internal/models"
	"Library/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

// ConnectDB открывает соединение с PostgreSQL и логирует результат.
func ConnectDB(cfg models.PostgresParams) error {
	// Формируем DSN
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)
	logger.Info.Printf("ConnectDB: connecting to Postgres at %s:%s/%s", cfg.Host, cfg.Port, cfg.Database)

	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error.Printf("ConnectDB: failed to connect to Postgres: %v", err)
		return fmt.Errorf("failed to connect to Postgres: %w", err)
	}

	logger.Info.Println("ConnectDB: ✅ Connected to PostgreSQL")
	return nil
}

// CloseDB закрывает соединение с базой и логирует результат.
func CloseDB() error {
	if db == nil {
		logger.Warn.Println("CloseDB: warning: database connection is already nil")
		return nil
	}
	err := db.Close()
	if err != nil {
		logger.Error.Printf("CloseDB: error closing DB: %v", err)
		return err
	}
	logger.Info.Println("CloseDB: database connection closed")
	return nil
}

// GetDBConn возвращает текущее подключение к БД.
func GetDBConn() *sqlx.DB {
	if db == nil {
		logger.Warn.Println("GetDBConn: warning: returning nil DB connection")
	}
	return db
}
