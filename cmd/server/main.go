package main

import (
	"context"
	"fmt"

	"internal/db/db"
)

// Setup app layers
func Run() error {
	fmt.Println("starting app")

	db, err = db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
