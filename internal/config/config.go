package config

import "os"

type Config struct {
	DBURL    string
	Platform string
	Port     string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		DBURL:    os.Getenv("DB_URL"),
		Platform: os.Getenv("PLATFORM"),
		Port:     port,
	}
}
