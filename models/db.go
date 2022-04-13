package models

import (
	"fmt"
	_redis "github.com/go-redis/redis/v7"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

type MyDb struct {
	*gorm.DB
}

var (
	once sync.Once
	db   *MyDb
)

//Init ...
func Init() {

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	ConnectDB(dbinfo)
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		fmt.Printf("error while creating MyDb extension 'uuid-ossp': %s\n", err)
	}
	db.AutoMigrate(&User{}, &Token{}, &Food{}, &Category{}, &Basket{}, &BasketItem{})
}

//ConnectDB ...,
func ConnectDB(dataSourceName string) error {
	var err error
	var gdb *gorm.DB
	once.Do(func() {
		gdb, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dataSourceName,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
	})
	db = &MyDb{gdb}
	return nil
}

//GetDB ...
func GetDB() *MyDb {
	return db
}

//RedisClient ...
var RedisClient *_redis.Client

//InitRedis ...
func InitRedis(selectDB ...int) {

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB[0],
	})

}

//GetRedis ...
func GetRedis() *_redis.Client {
	return RedisClient
}

func (g *MyDb) Paging(sort string, offset, limit int) *gorm.DB {
	return db.Offset((offset - 1) * limit).Limit(limit).Order(sort)
}
