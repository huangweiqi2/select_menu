package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()
var RDB = InitRedisDB()

func Init() *gorm.DB {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/select_menu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Gorm init Error" + err.Error())
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Food{})
	if err != nil {
		panic(err)
	}

	return db
}
func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
