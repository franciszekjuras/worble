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

	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error parsing database config: %v\n", err)
	}
	config.MinConns = 0
	config.MaxConns = 4
	config.MaxConnIdleTime = 10 * time.Minute
	db, err := pgxpool.NewWithConfig(context.Background(), config)
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
	sessionManager.Store = pgxstore.NewWithCleanupInterval(db, 2*time.Hour) // must be more than db scale to zero time (one hour)
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
