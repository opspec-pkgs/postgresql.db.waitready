package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	username := os.Getenv("dbUsername")
	password := os.Getenv("dbPassword")
	host := os.Getenv("dbHostname")
	port := os.Getenv("dbPort")
	schema := os.Getenv("dbSchema")
	seconds := os.Getenv("timeout")

	if seconds == "" {
		log.Fatalf("Unable to get timeout from environment")
	}

	timeoutSecs, err := strconv.Atoi(seconds)

	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, schema)
	if err != nil {
		log.Fatalf("Unable to get timeout %v", err)
	}

	timeout := time.Duration(timeoutSecs) * time.Second

	config, err := pgx.ParseConfig(connectionStr)
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	startTime := time.Now()

	for {
		fmt.Println("trying to connect...")
		conn, err := pgx.ConnectConfig(context.Background(), config)
		if err == nil {
			defer conn.Close(context.Background())

			if err = conn.Ping(context.Background()); err == nil {
				break
			}

		}

		if time.Now().Sub(startTime) > timeout {
			log.Fatal("Unable to connect to the database!")
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Println("connection established...")
}
