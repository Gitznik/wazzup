package db

import (
	"context"
	"log"

	"github.com/gitznik/wazzup/ent"

	_ "github.com/mattn/go-sqlite3"
)

func Startup() *ent.Client {
	dsn := "file:wazzup.db?_fk=1"
	client, err := ent.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
