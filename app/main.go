package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"worble.ow6.foo/app/web"
	"worble.ow6.foo/appui/uitempl"
)

func main() {
	ts, err := uitempl.InitTemplates()
	if err != nil {
		log.Fatalln("Failed to initialize templates: ", err)
	}

	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to initialize database connection pool: %v\n", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = db.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	} else {
		log.Println("Connected to the database")
	}

	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := web.App{Ts: ts, SessionManager: sessionManager}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = app.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
