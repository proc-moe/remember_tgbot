package conf

import (
	"errors"

	"github.com/pelletier/go-toml/v2"
	"github.com/proc-moe/remember_tgbot/utils/files"
	"github.com/proc-moe/remember_tgbot/utils/klog"
)

type ConfT struct {
	Telegram ConfTelegramT
}

type ConfTelegramT struct {
	Api_token string
}

var Conf *ConfT

func Init() {
	Conf = &ConfT{}
	ok, err := files.CheckExist(CONFIG_RELATIVE_PATH)
	if err != nil {
		klog.E("[Init Conf] check config file failed, err=%v", err)
		panic(err)
	}
	if !ok {
		klog.E("[Init Conf] config file not found")
		panic(errors.New("config file not found"))
	}

	config, err := files.Read(CONFIG_RELATIVE_PATH)
	if err != nil {
		klog.E("[Init Conf] read config file failed,err=%v", err)
		panic(err)
	}

	err = toml.Unmarshal(config, Conf)
	if err != nil {
		klog.E("[Init Conf] toml unmarshal failed,err=%v", err)
		panic(err)
	}
}
