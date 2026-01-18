package handler

import (
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/configuration/options"
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/container/service_locator"
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/event_bus"
)

func (h *Handler) pushOffline(ctx *eventbus.PushContext) {
	for _, e := range ctx.Events {

		for _, toUid := range e.OfflineUsers {
			fromUid := e.Conn.Uid
			// 是否是AI
			if fromUid != toUid && h.isAI(toUid) && !e.Frame.GetsyncOnce() && !options.G.IsSystemUid(fromUid) {
				// 处理AI推送
				h.processAIPush(toUid, e)
			}
		}
	}
	service.Webhook.NotifyOfflineMsg(ctx.Events)
}
