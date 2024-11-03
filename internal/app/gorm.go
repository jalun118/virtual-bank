package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var loggerDebug = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	},
)

var loggerProduction = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second * 5,
		Colorful:                  false,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
		LogLevel:                  logger.Info,
	},
)

func NewDatabase(viper *viper.Viper) *gorm.DB {
	USERNAME := viper.GetString("DATABASE_USERNAME")
	PASSWORD := viper.GetString("DATABASE_PASSWORD")
	HOST := viper.GetString("DATABASE_HOST")
	DATABASE := viper.GetString("DATABASE_NAME")
	PORT := viper.GetInt("DATABASE_PORT")
	IDLE_CONNECTION := viper.GetInt("DATABASE_IDLE_CONNECTION")
	MAX_CONNECTION := viper.GetInt("DATABASE_MAX_CONNECTION")
	MAX_LIFE_TIME_CONNECTION := viper.GetInt("DATABASE_MAX_LIFE_TIME_CONNECTION")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORT, DATABASE)

	gormConfig := &gorm.Config{}

	mode := viper.GetString("GIN_MODE")
	switch mode {
	case "debug":
		gormConfig = &gorm.Config{
			Logger: loggerDebug,
		}
	case "release":
		gormConfig = &gorm.Config{
			Logger: loggerProduction,
		}
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(IDLE_CONNECTION)
	connection.SetMaxOpenConns(MAX_CONNECTION)
	connection.SetConnMaxLifetime(time.Second * time.Duration(MAX_LIFE_TIME_CONNECTION))

	return db
}
