package main

import (
	"fmt"

	"github.com/joshprewer/go-rest-api/internal/comment"
	"github.com/joshprewer/go-rest-api/internal/db"
	transportHttp "github.com/joshprewer/go-rest-api/internal/transport/http"
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
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
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
