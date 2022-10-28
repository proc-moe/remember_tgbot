package main

import (
	"github.com/proc-moe/remember_tgbot/dal"
	"github.com/proc-moe/remember_tgbot/utils/logs"
)

func Init() {
	logs.Init()
	dal.InitDal()
}
func main() {

}
