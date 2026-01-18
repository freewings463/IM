package handler

import (
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/configuration/options"
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/container/service_locator"
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/event_bus"
	wkproto "github.com/WuKongIM/WuKongIMGoProto"
	"go.uber.org/zap"
)

func (h *Handler) connack(ctx *eventbus.UserContext) {
	for _, event := range ctx.Events {
		conn := event.Conn
		frame := event.Frame
		uid := conn.Uid
		if conn.NodeId == 0 {
			h.Error("processConnack: from node is 0", zap.String("uid", uid))
			return
		}
		if frame == nil {
			h.Error("processConnack: frame is nil", zap.String("uid", uid))
			return
		}
		connack := frame.(*wkproto.ConnackPacket)
		if connack.ReasonCode == wkproto.ReasonSuccess {
			// 设置连接最大空闲时间
			if options.G.IsLocalNode(conn.NodeId) {
				realConn := service.ConnManager.GetConn(conn.ConnId)
				if realConn != nil {
					realConn.SetMaxIdle(options.G.ConnIdleTime)
					realConn.SetContext(conn)
				}
			}
			connack.NodeId = options.G.Cluster.NodeId
			// 更新连接
			eventbus.User.UpdateConn(conn)
		}
		eventbus.User.ConnWrite(event.ReqId, conn, connack)
	}

}
