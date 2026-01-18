package handler

import (
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/container/service_locator"
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/event_bus"
)

func (h *Handler) conversation(channelId string, channelType uint8, tagKey string, events []*eventbus.Event) {
	service.ConversationManager.Push(channelId, channelType, tagKey, events)
}
