package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGORMDatabase() *gorm.DB {
	env := os.Getenv("APP_ENV")
	var DSN string
	switch env {
	case "pro":
		DSN = os.Getenv("DSN")
	case "dev":
		DSN = "kazuki:secret@tcp(db:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	default:
		log.Print("databse environment isn't properly selected")
	}
	db, err := gorm.Open(
		mysql.Open(DSN),
	)
	if err != nil {
		log.Fatalf("failed to connect with %s database: %s", env, err.Error())
	} else {
		log.Printf("%s database connected 📦", env)
	}
	return db
}

func NewRedisCilent(ctx context.Context) *redis.Client {
	env := os.Getenv("APP_ENV")
	client := &redis.Client{}
	switch env {
	case "pro":
		REDIS_URL := os.Getenv("REDIS_URL")
		option, _ := redis.ParseURL(REDIS_URL)
		client = redis.NewClient(option)
	case "dev":
		client = redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		})
	}
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to %s redis: %v", env, err)
	}
	log.Printf("%s redis client connected", env)
	return client
}
