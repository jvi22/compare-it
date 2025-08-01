package main

import (
	"compare-it/config"
	"compare-it/internal/api"
	"compare-it/internal/persistence"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load .env")
	}

	if err := persistence.ConnectRedis(config.RedisURI); err != nil {
		log.Fatalf("Redis: %v", err)
	}
	if err := persistence.ConnectMongo(config.MongoURI); err != nil {
		log.Fatalf("Mongo: %v", err)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Initialize Config and register routes
	app := &api.Config{Router: r}
	app.Routes()

	r.Run(":8080")
}
