package config

import "os"

type EnvConfig struct {
	DBURL    string
	Platform string
	Port     string
	Secret   string
}

func Load() EnvConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return EnvConfig{
		DBURL:    os.Getenv("DB_URL"),
		Platform: os.Getenv("PLATFORM"),
		Port:     port,
		Secret:   os.Getenv("SECRET"),
	}
}
