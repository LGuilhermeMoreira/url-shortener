package config_test

import (
	"testing"

	"github.com/LGuilhermeMoreira/url-shortener/config"
)

func TestCreateConfig(t *testing.T) {
	c, err := config.NewConfig()

	if err != nil {
		t.Fatalf("failed to create a struct config: %v", err)
	}

	if c.DBUser == "" {
		t.Errorf("DBUser is empty")
	}
	if c.DBPassword == "" {
		t.Errorf("DBPassword is empty")
	}
	if c.DBName == "" {
		t.Errorf("DBName is empty")
	}
	if c.DBPort == "" {
		t.Errorf("DBPort is empty")
	}
	if c.DBDriver == "" {
		t.Errorf("DBDriver is empty")
	}
	if c.Port == "" {
		t.Errorf("Port is empty")
	}
	if c.DBUri == "" {
		t.Errorf("DBUri is empty")
	}
}
