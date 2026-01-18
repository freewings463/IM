package handler

import (
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/event_bus"
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/monitoring/tracing/trace"
	wkproto "github.com/WuKongIM/WuKongIMGoProto"
)

func (h *Handler) ping(event *eventbus.Event) {
	conn := event.Conn

	trace.GlobalTrace.Metrics.App().PingCountAdd(1)
	trace.GlobalTrace.Metrics.App().PingBytesAdd(1) // ping的大小就是1

	trace.GlobalTrace.Metrics.App().PongCountAdd(1)
	trace.GlobalTrace.Metrics.App().PongBytesAdd(1)
	eventbus.User.ConnWrite(event.ReqId, conn, &wkproto.PongPacket{})
}
