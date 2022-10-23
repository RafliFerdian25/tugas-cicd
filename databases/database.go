package databases

import (
	"belajar-go-echo/config"
	"belajar-go-echo/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	cfg := config.Cfg

	username := cfg.DB_USERNAME
	password := cfg.DB_PASSWORD
	host := cfg.DB_HOST
	port := cfg.DB_PORT
	dbName := cfg.DB_NAME
	
	// username := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName,
	)
	fmt.Printf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	username,
	password,
	host,
	port,
	dbName,
	)
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}
