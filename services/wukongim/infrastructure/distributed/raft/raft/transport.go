package raft

import "github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/distributed/raft/types"

type Transport interface {
	// Send 发送事件
	Send(event types.Event)
}
