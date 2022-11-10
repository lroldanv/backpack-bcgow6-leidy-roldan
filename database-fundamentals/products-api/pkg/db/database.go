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
		log.Fatal("Error: Loading .env ", err.Error())
	}

	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		DBName: os.Getenv("DBNAME"),
	}

	// db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DATABASE")))
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
