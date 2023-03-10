package services

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var once sync.Once

var instance *gorm.DB

func NewDatabase() *gorm.DB {
	once.Do(func() {
		instance = initializeDB()
	})
	return instance
}

func CloseDatabase() {
	database := NewDatabase()
	db, err := database.DB()
	CheckError("Error Getting Database", err)
	err = db.Close()
	CheckErrorOrSuccess("Error Closing Database", "Database Closed SuccessFully", err)
}

type database struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

func loadVariables() (db database) {
	err := godotenv.Load()
	CheckError("Error loading .env file", err)
	db.Host = os.Getenv("HOST")
	db.Port = os.Getenv("DBPORT")
	db.User = os.Getenv("USER")
	db.Name = os.Getenv("NAME")
	db.Password = os.Getenv("PASSWORD")
	return
}

func initializeDB() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	dataBase := loadVariables()
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dataBase.Host, dataBase.User, dataBase.Name, dataBase.Password, dataBase.Port)
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	CheckErrorOrSuccess("Cannot Connect to Database", "Successfully Connected to Database...", err)
	return db
}
