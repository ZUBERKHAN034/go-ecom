package main

import (
	"log"

	"github.com/ZUBERKHAN034/go-ecom/cmd/api"
	"github.com/ZUBERKHAN034/go-ecom/config"
	"github.com/ZUBERKHAN034/go-ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	addr := ":8080"

	db, err := db.NewMySQL(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(addr, db)
	if err := server.RUN(); err != nil {
		log.Fatal(err)
	}
}
