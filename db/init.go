package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"puny/utils"
)

var DB *sql.DB

func InitDB() *sql.DB {
	if DB != nil {
		return DB
	}
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))
	fmt.Println(connStr)
	var err error
	DB, err = sql.Open("postgres", connStr)
	utils.PanicIfError(err)
	return DB
}

func Close() {
	err := DB.Close()
	print("closeing")
	if err != nil {
		panic(err)
	}
}
