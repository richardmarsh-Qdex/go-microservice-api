package config

import (
	"os"
)

type Config struct {
	Port       string
	MongoURI   string
	DBName     string
	JWTSecret  string
	Env        string
	CORSOrigin string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "microservice_db"
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key"
	}
	cors := os.Getenv("CORS_ORIGIN")
	if cors == "" {
		cors = "*"
	}
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	return Config{
		Port:       port,
		MongoURI:   mongoURI,
		DBName:     dbName,
		JWTSecret:  jwtSecret,
		Env:        env,
		CORSOrigin: cors,
	}
}
