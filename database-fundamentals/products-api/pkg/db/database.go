package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDatabase() (engine *gin.Engine, db *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Loading .env")
	}

	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "192.168.111.136.",
		DBName: os.Getenv("DBNAME"),
	}

	// db, err := sql.Open("mysql", dataSource) --> dataSource:="user:password@tcp(localhost:3306)/DBName"
	db, err = sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	engine = gin.Default()

	return engine, db
}
