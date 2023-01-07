package tg

import (
	"fmt"
	"log"
	"os"

	"github.com/IliaBelov/RPDB/music"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	//"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot   *tgbotapi.BotAPI
	store *music.Store
}

// конструктор
func NewBot() *Bot {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	db, err := music.NewPosthresDB(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("failed intializade DB: %s", err.Error())
	}

	return &Bot{bot: bot, store: db}
}

func (b *Bot) Start() error {
	b.bot.Debug = true

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)
	/*updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}*/

	b.handleUpdates(updates)
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.Message == nil { // If we got a message
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}
		b.handleMessage(update.Message)
	}
}
