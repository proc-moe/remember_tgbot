package dal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/proc-moe/remember_tgbot/conf"
	klog "github.com/proc-moe/remember_tgbot/utils/klog"
)

type TGProxy struct {
	bot *tgbotapi.BotAPI
}

func InitTGProxy() *TGProxy {
	klog.I("[InitTGProxy] start...")
	p := &TGProxy{}
	bot, err := tgbotapi.NewBotAPI(conf.Conf.Telegram.Api_token)
	if err != nil {
		panic(err)
	}
	klog.I("Authorized on account:%s", bot.Self.UserName)
	bot.Debug = true
	p.bot = bot
	klog.I("telegram bot inited")
	return p
}

func (p *TGProxy) Listen() {
	go func() {
		klog.I("[tgbot] start listening")
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := p.bot.GetUpdatesChan(u)

		for update := range updates {
			klog.I("[%s]:%s", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			p.bot.Send(msg)
		}
	}()
}

func (p *TGProxy) Send() {

}
