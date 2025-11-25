package db

import (
	"context"
	"log"

	"github.com/gitznik/wazzup/ent"

	_ "github.com/mattn/go-sqlite3"
)

func Startup(dbDSN string) *ent.Client {
	client, err := ent.Open("sqlite3", dbDSN)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
