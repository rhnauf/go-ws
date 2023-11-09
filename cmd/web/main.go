package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-ws/internal/db"
	"log"
	"os"
)

func Run() error {
	// initialize env variables
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading environment variables")
		return nil
	}

	// initialize db conn
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Println("failed to connect to the database")
		return err
	}

	// initialize routing
	app := routes(dbConn)

	port := os.Getenv("PORT")
	log.Println("starting web server on port:", port)

	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		panic(err)
	}
}
