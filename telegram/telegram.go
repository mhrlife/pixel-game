package telegram

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
	"os"
	"time"
)

type Telegram struct {
	bot *telebot.Bot
}

func NewTelegram() (*Telegram, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		logrus.WithError(err).Error("couldn't create telegram bot")
		return nil, err
	}

	return &Telegram{
		bot: bot,
	}, nil
}

func (t *Telegram) Start() {
	t.bot.Handle(telebot.OnText, t.handle)

	t.bot.Start()
}
