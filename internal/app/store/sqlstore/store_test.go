package sqlstore_test

import (
	"os"
	"testing"
)

var (
	DatabaseURL string
)

func TestMau(m *testing.M) {
	DatabaseURL = os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		DatabaseURL = "host=localhost dname = restapi_test sslmode = disable"
	}

	os.Exit(m.Run())
}
