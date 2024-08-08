package config

import (
	"os"
)

type Config struct {
	TGToken string
	DBURL   string
}

func LoadConfig() Config {
	return Config{
		TGToken: os.Getenv("TG_TOKEN"),
		DBURL:   os.Getenv("DB_URL"),
	}
}
