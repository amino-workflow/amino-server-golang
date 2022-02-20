package config

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/mysql"    // dummy import mysql driver
	"gorm.io/driver/postgres" // dummy postgres driver import
	"gorm.io/gorm"
)

const POSTGRES_DB_TYPE string = "psql"
const MYSQL_DB_TYPE string = "mysql"

var (
	dbHost          = Get("DB_HOST")
	dbPort          = Get("DB_PORT")
	dbName          = Get("DB_NAME")
	dbUser          = Get("DB_USER")
	dbPassword      = Get("DB_PASSWORD")
	dbMaxConnection = GetInt("DB_MAX_CONNECTION")
	dbMinConnection = GetInt("DB_MIN_CONNECTION")
	dbType          = Get("DB_TYPE")
)

var once sync.Once
var db *gorm.DB

// getMySQLURL -> builds mysql url for mysql
func getMySQLConnection() (*gorm.DB, error) {
	mysqlUrl := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	return gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{})
}

// getPostgresSQLURL -> builds postgres url for mysql
func getPostgresSQLConnection() (*gorm.DB, error) {
	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	return gorm.Open(postgres.Open(psqlUrl), &gorm.Config{})
}

// instantiateDB -> instantiates DB
func instantiateDB() *gorm.DB {
	dbConnection, err := func() (*gorm.DB, error) {
		if dbType == POSTGRES_DB_TYPE {
			return getPostgresSQLConnection()
		} else {
			return getMySQLConnection()
		}
	}()

	if err != nil {
		panic("Unable to Connect to DB" + err.Error())
	}
	dbInstance, _ := dbConnection.DB()
	dbInstance.SetMaxIdleConns(dbMinConnection)
	dbInstance.SetMaxOpenConns(dbMaxConnection)
	dbInstance.SetConnMaxLifetime(time.Hour)
	db = dbConnection
	return dbConnection
}

// StartDB -> Configures db and opens db connection
func StartDB() {
	once.Do(func() {
		instantiateDB()
	})
}

// GetDB -> Configures db and opens db connection
func GetDB() *gorm.DB {
	return db
}

func CloseDB() *gorm.DB {
	dbInstance, _ := db.DB()
	dbInstance.Close()
	return db
}
