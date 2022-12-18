package conf

import (
	"testing"

	"github.com/proc-moe/remember_tgbot/utils/klog"
)

func TestConfInit(t *testing.T) {
	klog.Init()
	Init()
}
