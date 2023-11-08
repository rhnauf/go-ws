package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-ws/internal/handler"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading environment variables")
		return
	}

	app := routes()

	log.Println("starting chat hub")
	go handler.Hub()

	port := os.Getenv("PORT")
	log.Println("starting web server on port:", port)

	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
