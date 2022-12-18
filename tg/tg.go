package tg

import (
	"context"

	"github.com/proc-moe/remember_tgbot/utils/klog"
)

type RemindRequest struct {
	Msg string
}

func Remind(ctx context.Context, req RemindRequest) error {
	klog.CtxI(ctx, "[Remind] %v", req.Msg)
	return nil
}
