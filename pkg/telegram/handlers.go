package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	b.Bot.Send(tgbotapi.NewMessage(int64(message.Chat.ID), message.Text))
	return nil
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	b.Bot.Send(tgbotapi.NewMessage(int64(message.Chat.ID), message.Text))
	return nil
}

// func (b *Bot) handleError(chatID int64, err error) error {
// 	_, e = b.Bot.Send(tgbotapi.NewMessage(chatID, err.Error()))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
