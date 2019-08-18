package main

import (
	"os"

	"github.com/JedBeom/wbot/model"
	"github.com/go-pg/pg"
)

var (
	db *pg.DB
)

func ConnectDB() {
	db = pg.Connect(&pg.Options{
		User:     config.DB.User,
		Password: config.DB.Password,
		Database: config.DB.Database,
	})

	mode := os.Getenv("MODE")
	if mode == "CREATE" {
		err := model.CreateTables(db)
		if err != nil {
			panic(err)
		}
	}
}
