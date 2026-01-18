package raftgroup

import "github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/distributed/raft/types"

type ITransport interface {
	// Send 发送事件
	Send(key string, event types.Event)
}
