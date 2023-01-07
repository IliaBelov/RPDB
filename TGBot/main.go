package main

import (
	"log"

	"github.com/IliaBelov/RPDB/tg"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	telegramBot := tg.NewBot()
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
