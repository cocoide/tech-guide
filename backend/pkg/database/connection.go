package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
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
