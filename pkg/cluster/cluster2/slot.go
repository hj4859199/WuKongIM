package cluster

import (
	"fmt"
	"time"

	pb "github.com/WuKongIM/WuKongIM/pkg/cluster/clusterconfig/cpb"
	"github.com/WuKongIM/WuKongIM/pkg/cluster/reactor"
	"github.com/WuKongIM/WuKongIM/pkg/cluster/replica"
	"github.com/WuKongIM/WuKongIM/pkg/wklog"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

var _ reactor.IHandler = &slot{}

type slot struct {
	rc         *replica.Replica
	key        string
	st         *pb.Slot
	isPrepared bool
	wklog.Log
	opts *Options

	lastIndex atomic.Uint64 // 当前频道最后一条日志索引
}

func newSlot(st *pb.Slot, opts *Options) *slot {
	s := &slot{
		key:        SlotIdToKey(st.Id),
		st:         st,
		opts:       opts,
		Log:        wklog.NewWKLog(fmt.Sprintf("slot[%d]", st.Id)),
		isPrepared: true,
	}
	s.rc = replica.New(opts.NodeId, replica.WithLogPrefix(fmt.Sprintf("slot-%d", st.Id)), replica.WithElectionOn(false), replica.WithReplicas(st.Replicas), replica.WithStorage(newProxyReplicaStorage(s.key, s.opts.SlotLogStorage)))
	return s
}

func (s *slot) becomeFollower(term uint32, leader uint64) {
	s.rc.BecomeFollower(term, leader)
}

func (s *slot) becomeLeader(term uint32) {
	s.rc.BecomeLeader(term)
}

func (s *slot) LastLogIndexAndTerm() (uint64, uint32) {

	return s.rc.LastLogIndex(), s.rc.Term()
}

func (s *slot) HasReady() bool {
	return s.rc.HasReady()
}

func (s *slot) Ready() replica.Ready {

	return s.rc.Ready()
}

func (s *slot) GetAndMergeLogs(msg replica.Message) ([]replica.Log, error) {

	unstableLogs := msg.Logs
	startIndex := msg.Index
	if len(unstableLogs) > 0 {
		startIndex = unstableLogs[len(unstableLogs)-1].Index + 1
	}
	shardNo := s.key
	lastIndex := s.lastIndex.Load()
	var err error
	if lastIndex == 0 {
		lastIndex, err = s.opts.SlotLogStorage.LastIndex(shardNo)
		if err != nil {
			s.Error("handleSyncGet: get last index error", zap.Error(err))
			return nil, err
		}
	}

	var resultLogs []replica.Log
	if startIndex <= lastIndex {
		logs, err := s.getLogs(startIndex, lastIndex+1, uint64(s.opts.LogSyncLimitSizeOfEach))
		if err != nil {
			s.Error("get logs error", zap.Error(err), zap.Uint64("startIndex", startIndex), zap.Uint64("lastIndex", lastIndex))
			return nil, err
		}
		resultLogs = extend(logs, unstableLogs)
	} else {
		// c.Warn("handleSyncGet: startIndex > lastIndex", zap.Uint64("startIndex", startIndex), zap.Uint64("lastIndex", lastIndex))
	}
	return resultLogs, nil

}

func (s *slot) AppendLog(logs []replica.Log) error {

	lastLog := logs[len(logs)-1]
	start := time.Now()
	s.Debug("append log", zap.Uint64("lastLogIndex", lastLog.Index))
	err := s.opts.SlotLogStorage.AppendLog(s.key, logs)
	if err != nil {
		s.Panic("append log error", zap.Error(err))
	}
	s.Debug("append log done", zap.Uint64("lastLogIndex", lastLog.Index), zap.Duration("cost", time.Since(start)))
	return nil
}

func (s *slot) ApplyLog(startLogIndex, endLogIndex uint64) error {

	if s.opts.OnSlotApply != nil {
		logs, err := s.getLogs(startLogIndex, endLogIndex, 0)
		if err != nil {
			s.Panic("get logs error", zap.Error(err))
		}
		if len(logs) == 0 {
			s.Panic("logs is empty", zap.Uint64("startLogIndex", startLogIndex), zap.Uint64("endLogIndex", endLogIndex))
		}
		err = s.opts.OnSlotApply(s.st.Id, logs)
		if err != nil {
			s.Panic("on slot apply error", zap.Error(err))
		}
	}
	return nil
}

func (s *slot) SlowDown() {

}

func (s *slot) SetHardState(hd replica.HardState) {

}

func (s *slot) Tick() {
	s.rc.Tick()
}

func (s *slot) Step(m replica.Message) error {

	return s.rc.Step(m)
}

func (s *slot) SetLastIndex(index uint64) error {
	lastIndex := s.lastIndex.Load()
	shardNo := s.key
	err := s.opts.SlotLogStorage.SetLastIndex(shardNo, lastIndex)
	if err != nil {
		s.Error("set last index error", zap.Error(err))
	}
	return nil
}

func (s *slot) SetAppliedIndex(index uint64) error {
	shardNo := s.key
	err := s.opts.SlotLogStorage.SetAppliedIndex(shardNo, index)
	if err != nil {
		s.Error("set applied index error", zap.Error(err))
	}
	return nil
}

func (s *slot) IsPrepared() bool {
	return s.isPrepared
}

func (s *slot) getLogs(startLogIndex uint64, endLogIndex uint64, limitSize uint64) ([]replica.Log, error) {
	shardNo := s.key
	logs, err := s.opts.SlotLogStorage.Logs(shardNo, startLogIndex, endLogIndex, limitSize)
	if err != nil {
		s.Error("get logs error", zap.Error(err))
		return nil, err
	}
	return logs, nil
}

func extend(dst, vals []replica.Log) []replica.Log {
	need := len(dst) + len(vals)
	if need <= cap(dst) {
		return append(dst, vals...) // does not allocate
	}
	buf := make([]replica.Log, need) // allocates precisely what's needed
	copy(buf, dst)
	copy(buf[len(dst):], vals)
	return buf
}
