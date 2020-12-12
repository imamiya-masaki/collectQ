package database

import (
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/appengine"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Init initializes database
func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DATABASE := os.Getenv("DATABASE")
	// USERNAME := os.Getenv("USERNAME")
	// USERPASS := os.Getenv("USERPASS")
	CONNECTIONNAME := "collecttq:asia-northeast1:collectq-mysql"
	USER := "collectqUser"
	PASS := "collectq"
	DBNAME := "collectq_db"
	// socketDir := "cloudsql"
	// set := USER + ":" + PASS + "@cloudsql(" + cfg + ")/" + DBNAME
	// dbURI := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", USER, PASS, cfg, DBNAME)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PASS, "127.0.0.1:3306", DBNAME)
	// dsn := "docker:docker@tcp(mysql:3306)/collectq_db?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := USERNAME + ":" + USERPASS + "@tcp(mysql:3306)/" + DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	//println("dsn", dsn)
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	localConnection := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	cloudSQLConnection := USER + ":" + PASS + "@unix(/cloudsql/" + CONNECTIONNAME + ")/" + DBNAME + "?parseTime=true"
	if appengine.IsAppEngine() {
		db, err = gorm.Open(mysql.Open(cloudSQLConnection), &gorm.Config{})
	} else {
		db, err = gorm.Open(mysql.Open(localConnection), &gorm.Config{})
	}
	if err != nil {
		println(err)
	} else {
		println(("db connection success !!"))
	}
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}
