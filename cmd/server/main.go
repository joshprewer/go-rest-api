package main

import (
	"context"
	"fmt"

	"github.com/joshprewer/go-rest-api/internal/comment"
	"github.com/joshprewer/go-rest-api/internal/db"
)

// Setup app layers
func Run() error {
	fmt.Println("starting app")

	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("failed to connect to databse")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate db")
		return err
	}

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"b2059fe8-f11d-41e7-9b12-2250642ef27d",
	))

	return nil
}

func main() {
	fmt.Println("Go Rest API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
