package service

import (
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/event_bus"
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/persistence/wkdb"
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/types"
	wkproto "github.com/WuKongIM/WuKongIMGoProto"
)

var Webhook IWebhook

type IWebhook interface {
	// Online 设备上线
	Online(uid string, deviceFlag wkproto.DeviceFlag, connId int64, deviceOnlineCount int, totalOnlineCount int)
	// Offline 设备下线
	Offline(uid string, deviceFlag wkproto.DeviceFlag, connId int64, deviceOnlineCount int, totalOnlineCount int)
	// NotifyOfflineMsg 离线消息通知
	NotifyOfflineMsg(events []*eventbus.Event)
	// TriggerEvent 触发事件
	TriggerEvent(event *types.Event)

	// AppendMessageOfNotifyQueue 追加消息到通知队列
	AppendMessageOfNotifyQueue(messages []wkdb.Message) error
}
