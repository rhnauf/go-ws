package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading environment variables")
		return
	}

	mux := routes()

	port := os.Getenv("PORT")
	log.Println("starting web server on port:", port)

	err = mux.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
