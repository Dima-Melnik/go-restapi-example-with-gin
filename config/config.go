package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Dima-Melnik/go-restapi-example-with-gin/internal/utils"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		DBHost     string `yaml:"db_host"`
		DBPort     int    `yaml:"db_port"`
		DBUser     string `yaml:"db_user"`
		DBPassword string
		DBName     string `yaml:"db_name"`
		DBSSLMode  string `yaml:"db_sslmode"`
	} `yaml:"database"`
}

var Cfg Config

func LoadConfig(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := yaml.Unmarshal(f, &Cfg); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func InitDatabaseConn() string {
	if Cfg.Database.DBPassword == "" {
		Cfg.Database.DBPassword = utils.GetEnv("DB_PASSWORD")
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		Cfg.Database.DBHost,
		Cfg.Database.DBPort,
		Cfg.Database.DBUser,
		Cfg.Database.DBPassword,
		Cfg.Database.DBName,
		Cfg.Database.DBSSLMode)

	return connStr
}

func InitServerPort() string {
	port := fmt.Sprintf(":%d", Cfg.Server.Port)

	return port
}
