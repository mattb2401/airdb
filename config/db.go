package config

import (
	"airdb/helpers"
	"fmt"
	"os"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Host     string
	Charset  string
}

func GetConfig() *Config {
	dbusername, err := helpers.Getenv("dbUsername")
	if err != nil {
		fmt.Println("DB configuration not found. Run ./airdb -i to configure the application first")
		os.Exit(103)
	}
	dbpassword, err := helpers.Getenv("dbPassword")
	if err != nil {
		fmt.Println("DB configuration not found. Run ./airdb -i to configure the application first")
		os.Exit(103)
	}
	dbhost, err := helpers.Getenv("dbHost")
	if err != nil {
		fmt.Println("DB configuration not found. Run ./airdb -i to configure the application first")
		os.Exit(103)
	}
	dbname, err := helpers.Getenv("dbName")
	if err != nil {
		fmt.Println("DB configuration not found. Run ./airdb -i to configure the application first")
		os.Exit(103)
	}
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: dbusername,
			Password: dbpassword,
			Host:     dbhost,
			Name:     dbname,
			Charset:  "utf8",
		},
	}
}
