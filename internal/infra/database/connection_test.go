package database_test

import (
	"testing"

	"github.com/LGuilhermeMoreira/url-shortener/config"
	"github.com/LGuilhermeMoreira/url-shortener/internal/infra/database"
)

func TestDBConnectionAndMigration(t *testing.T) {
	c, _ := config.NewConfig()
	if c == nil {
		t.Fatalf("failed to create config")
	}

	db, err := database.NewConnection(c)
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Errorf("failed to ping database: %v", err)
	}

	err = database.Migration(db)

	if err != nil {
		t.Errorf("failed to migrate database: %v", err)
	}
}
