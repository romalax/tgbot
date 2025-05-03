package api

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (a *Api) SendMessage(msg string) error {
	msgConfig := tgbotapi.NewMessage(a.cfg.ChatID, msg)

	_, err := a.bot.Send(msgConfig)

	return err
}
