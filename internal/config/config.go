package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"Library/internal/models"
)

var AppSettings models.Configs

// ReadSettings загружает .env и internal/config/configs.json
func ReadSettings() error {
	// 1) Подгружаем переменные из .env
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env not found, делаем дальше по окружению")
	}

	// 2) Читаем секции log/auth/app из JSON
	f, err := os.Open("internal/config/configs.json")
	if err != nil {
		return errors.New("cannot open configs.json: " + err.Error())
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&AppSettings); err != nil {
		return errors.New("cannot parse configs.json: " + err.Error())
	}

	// 3) Перезаписываем PostgresParams из ENV
	AppSettings.PostgresParams = models.PostgresParams{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	// Дополнительно проверяем, что всё не пусто
	if AppSettings.PostgresParams.Host == "" ||
		AppSettings.PostgresParams.Port == "" ||
		AppSettings.PostgresParams.User == "" ||
		AppSettings.PostgresParams.Password == "" ||
		AppSettings.PostgresParams.Database == "" {
		return errors.New("missing one of DB_HOST/DB_PORT/DB_USER/DB_PASSWORD/DB_NAME in .env")
	}

	return nil
}
