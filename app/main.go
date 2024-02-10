package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"worble.ow6.foo/app/web"
	"worble.ow6.foo/appui/uitempl"
	"worble.ow6.foo/biz/worble"
)

func main() {
	ts, err := uitempl.InitTemplates()
	if err != nil {
		log.Fatalln("Failed to initialize templates: ", err)
	}

	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	} else {
		log.Println("Connected to database")
	}
	defer db.Close()

	app := web.App{Ts: ts, Game: worble.NewGame()}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = app.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
