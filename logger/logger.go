package logger

import (
	"Library/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
)

func Init() error {
	p := config.AppSettings.LogParams

	// создаём папку, если нет
	if _, err := os.Stat(p.LogDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(p.LogDirectory, 0755); err != nil {
			return err
		}
	}

	// файловые ротаторы
	infoOut := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", p.LogDirectory, p.LogInfo),
		MaxSize:    p.MaxSizeMegabytes,
		MaxBackups: p.MaxBackups,
		MaxAge:     p.MaxAgeDays,
		Compress:   p.Compress,
		LocalTime:  p.LocalTime,
	}
	errorOut := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", p.LogDirectory, p.LogError),
		MaxSize:    p.MaxSizeMegabytes,
		MaxBackups: p.MaxBackups,
		MaxAge:     p.MaxAgeDays,
		Compress:   p.Compress,
		LocalTime:  p.LocalTime,
	}
	warnOut := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", p.LogDirectory, p.LogWarn),
		MaxSize:    p.MaxSizeMegabytes,
		MaxBackups: p.MaxBackups,
		MaxAge:     p.MaxAgeDays,
		Compress:   p.Compress,
		LocalTime:  p.LocalTime,
	}
	debugOut := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", p.LogDirectory, p.LogDebug),
		MaxSize:    p.MaxSizeMegabytes,
		MaxBackups: p.MaxBackups,
		MaxAge:     p.MaxAgeDays,
		Compress:   p.Compress,
		LocalTime:  p.LocalTime,
	}

	// создаём глобальные логгеры
	Info = log.New(io.MultiWriter(os.Stdout, infoOut), "INFO:  ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorOut, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(warnOut, "WARN:  ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(debugOut, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Gin-логи (по умолчанию INFO) пишем и в STDOUT, и в файл info
	gin.DefaultWriter = io.MultiWriter(os.Stdout, infoOut)
	return nil
}
