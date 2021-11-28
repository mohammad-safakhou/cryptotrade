package utils

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func PostgresConnection(host, port, user, pass, database, sslmode string) (*sql.DB, error) {
	connString := postgresURI(host, port, user, pass, database, sslmode)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error in openning postgres connection: %w", err)
	}

	dbContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = conn.PingContext(dbContext)
	if err != nil {
		return nil, fmt.Errorf("error in pinging postgres database: %w", err)
	}
	return conn, nil
}

func postgresURI(host, port, user, pass, database, sslmode string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, database, sslmode)
}
