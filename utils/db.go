package utils

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
	"mini-dsp-backend/config"
)

var DB *gorm.DB
var DorisDB *sql.DB

func InitDB() {
	db, err := gorm.Open(mysql.Open(config.GetConfig().MySQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db
	fmt.Println("MySQL connected successfully.")

	dorisDb, err := sql.Open("mysql", config.GetConfig().DorisDSN)
	if err != nil {
		log.Fatal("Fail to connect Doris: ", err)
	}
	DorisDB = dorisDb
}
