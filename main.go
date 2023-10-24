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
	tx := utils.HandleFunctionFetal(dbHandler.BeginTx(ctx, nil))
	q := db.New(tx)
	print(q)
	defer db.Close()
}
