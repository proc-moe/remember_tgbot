package main

import (
	"time"

	"github.com/proc-moe/remember_tgbot/conf"
	"github.com/proc-moe/remember_tgbot/dal"
	"github.com/proc-moe/remember_tgbot/utils/klog"
)

func Init() {
	klog.Init()
	dal.Init()
	conf.Init()
}
func main() {
	Init()
	dal.Cli.Tg.Listen()
	time.Sleep(60 * time.Second)
}
