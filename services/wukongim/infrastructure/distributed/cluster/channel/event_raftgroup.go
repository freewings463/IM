package channel

import (
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/distributed/raft/raftgroup"
	"github.com/WuKongIM/WuKongIM/services/wukongim/infrastructure/monitoring/tracing/trace"
)

// OnAddRaft 添加raft
func (s *Server) OnAddRaft(r raftgroup.IRaft) {
	if trace.GlobalTrace != nil {
		trace.GlobalTrace.Metrics.Cluster().ChannelActiveCountAdd(1)
	}
}

// OnRemoveRaft 移除raft
func (s *Server) OnRemoveRaft(r raftgroup.IRaft) {
	if trace.GlobalTrace != nil {
		trace.GlobalTrace.Metrics.Cluster().ChannelActiveCountAdd(-1)
	}
}
