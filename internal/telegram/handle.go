package telegram

import (
	"gopkg.in/telebot.v4"
	"os"
)

func (t *Telegram) handle(c telebot.Context) error {
	return c.Reply(`🖼 به تصدانه (Pixel) خوش اومدی!

توی این بازی تو می‌تونی با کمک دوستات، با بقیه رقابت کنید و نقاشیتون رو بکشید!

برای شروع روی «اجرای بازی» کلیک کن.`, &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{
				{
					Text: "🎮 اجرای بازی",
					WebApp: &telebot.WebApp{
						URL: os.Getenv("WEBAPP_URL"),
					},
				},
			},
		},
	})
}
