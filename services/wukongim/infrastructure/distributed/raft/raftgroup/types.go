package raftgroup

import "github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/distributed/raft/types"

type Event struct {
	RaftKey string
	types.Event
	WaitC chan error
}
