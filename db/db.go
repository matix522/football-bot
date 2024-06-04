package db

import (
	"fmt"
	"os"
)

const CONNECTION_STRING string = "postgresql://%s:%s@db:5432/%s?sslmode=disable"

func ConnectionString() string {
	return fmt.Sprintf(CONNECTION_STRING, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
}
