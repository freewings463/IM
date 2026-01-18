package raft

import "github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/distributed/raft/types"

type stepReq struct {
	event types.Event
	resp  chan error
}
