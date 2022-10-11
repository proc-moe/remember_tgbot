package main

import (
	"context"

	"github.com/proc-moe/remember_tgbot/utils/logs"
)

func Init() {
	logs.Init()
}
func main() {
	Init()
	a := 2
	ctx := context.WithValue(context.Background(), "logId", 3432)
	logs.Log.E("你好 %v", a)
	logs.Log.W("hello world")
	logs.Log.I("hello world")
	logs.Log.CtxE(ctx, "你好呀 %v", a)
	logs.Log.CtxW(ctx, "你好呀 %v", a)
	logs.Log.CtxI(ctx, "你好呀 %v", a)
}
