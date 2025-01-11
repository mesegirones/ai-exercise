package config

import (
	"os"
)

func parsePort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func NewConfig() *Config {
	service := os.Getenv("CONTAINER_NAME")
	prod := os.Getenv("ENVIRONMENT") == "production"
	local := os.Getenv("ENVIRONMENT") == "local"
	return &Config{
		restConfig: ConfigRest{
			prod:    prod,
			local:   local,
			service: service,
			port:    parsePort(),
		},
		urlConfig: ConfigURL{
			service: os.Getenv("SERVICE_URL"),
		},
		healthConfig: ConfigHealth{
			service: service,
		},
	}
}
