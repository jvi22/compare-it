package main

import (
	"compare-it/config"
	"compare-it/internal/api"
	"compare-it/internal/persistence"
	"log"
	

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Load .env before anything else runs
func init() {
	err := godotenv.Load(".env") // Adjust path if your .env is in a different directory
	if err != nil {
		log.Println("No .env file found or failed to load .env")
	}

}

func main() {
	// Connect to Redis
	if err := persistence.ConnectRedis(config.RedisURI); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	// Connect to MongoDB
	if err := persistence.ConnectMongo(config.MongoURI); err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	// Set up Gin router
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Register routes
	app := &api.Config{Router: r}
	app.Routes()

	// Run server on port 8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
