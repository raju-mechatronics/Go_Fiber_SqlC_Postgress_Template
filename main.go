package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"puny/db"
	"puny/utils"
)

func main() {
	utils.FetalError(godotenv.Load())
	ctx := context.Background()
	dbHandler := db.InitDB()
	tx, err := dbHandler.BeginTx(ctx, nil)
	utils.FetalError(err)
	q := db.New(tx)
	defer db.Close()
}
