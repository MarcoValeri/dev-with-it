package util

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func Connect() *sql.DB {

	// ENV data
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	userEnv := os.Getenv("USER")
	passwordEnv := os.Getenv("PASSWORD")
	newEnv := os.Getenv("NET")
	addrEnv := os.Getenv("ADDR")
	databaseEnv := os.Getenv("DATABASE")

	// Capture connection properties.
	cfg := mysql.Config{
		User:                 userEnv,
		Passwd:               passwordEnv,
		Net:                  newEnv,
		Addr:                 addrEnv,
		DBName:               databaseEnv,
		AllowNativePasswords: true,
	}

	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db

}
