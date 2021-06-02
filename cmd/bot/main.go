package main

import (
	"log"
	"os"
	"time"

	telegram "github.com/mbtamuli/telegram-bot/handler"
	"github.com/mbtamuli/telegram-bot/sqlite"
)

const shutdownTimeout = 3600

func main() {
	dsn := "data/telegram.sqlite"

	store, err := sqlite.NewStore(dsn)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("Starting handler")
		time.Sleep(shutdownTimeout * time.Second)
		log.Printf("Shutting down handler after %d seconds\n", shutdownTimeout)
		os.Exit(0)
	}()

	telegram.Handler(store)
}
