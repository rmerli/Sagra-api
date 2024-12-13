package main

import (
	"context"
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"sagre/src/database/model"
	"sagre/src/server"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func run(ctx context.Context, getenv func(string) string) error {

	// conn, err := pgx.Connect(ctx, getenv("DB_URL"))

	conn, err := sql.Open("pgx", getenv("DB_URL"))
	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})

	store, err := pgstore.NewPGStore(getenv("DB_URL"), []byte(getenv("STORE_KEY")))
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer store.Close()

	gob.Register(model.User{})

	defer store.StopCleanup(store.Cleanup(time.Minute * 60 * 24))

	serverReady := make(chan bool)
	server := server.New(db, store, serverReady)
	return server.Start(fmt.Sprintf("%s:%s", getenv("ADDRESS"), getenv("PORT")))
}

func main() {
	ctx := context.Background()

	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load(".env-local")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

	if err := run(ctx, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
