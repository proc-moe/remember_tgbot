package dal

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/proc-moe/remember_tgbot/utils/logs"
)

type TGProxy struct {
	bot *tgbotapi.BotAPI
}

func InitTGProxy() *TGProxy {
	p := &TGProxy{}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	bot.Debug = true
	logs.I("Authorized on account %s", bot.Self.UserName)
	if err != nil {
		panic(err)
	}
	p.bot = bot
	logs.I("telegram bot inited")
	return p
}

func (p *TGProxy) Message() {
	go func() {
		logs.I("[tgbot] start listening")
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := p.bot.GetUpdatesChan(u)

		for update := range updates {
			logs.I("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			p.bot.Send(msg)
		}
	}()
}
