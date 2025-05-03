package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

type Api struct {
	cfg Config
	bot *tgbotapi.BotAPI
}

func New(c Config) (*Api, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		return nil, err
	}

	return &Api{
		cfg: c,
		bot: bot,
	}, nil
}
