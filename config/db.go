package config

import (
	"airdb/helpers"
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
		panic("db username not found")
	}
	dbpassword, err := helpers.Getenv("dbPassword")
	if err != nil {
		panic("db password not found")
	}
	dbhost, err := helpers.Getenv("dbHost")
	if err != nil {
		panic("db host not found")
	}
	dbname, err := helpers.Getenv("dbName")
	if err != nil {
		panic("db name not found")
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
