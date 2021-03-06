package storage

import (
	"container/ring"
	"sync"

	"github.com/artofey/sysmon"
)

type RingStorage struct {
	count int

	mu sync.Mutex
	s  *ring.Ring
}

func NewRingStorage(count int) *RingStorage {
	return &RingStorage{
		count: count,
		s:     ring.New(count),
	}
}

func (s *RingStorage) Add(st sysmon.Stats) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.s.Value = st
	s.s = s.s.Next()
	return nil
}

func (s *RingStorage) Len() int {
	return s.realLen()
}

func (s *RingStorage) realLen() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	var len int
	var r *ring.Ring

	r = s.s
	for i := 0; i < r.Len(); i++ {
		if r.Value != nil {
			len++
		}
		r = r.Next()
	}
	return len
}

func (s *RingStorage) GetLast(l int) []sysmon.Stats {
	rLen := s.Len()
	if rLen < l {
		return s.get(rLen)
	}
	return s.get(l)
}

func (s *RingStorage) get(l int) []sysmon.Stats {
	s.mu.Lock()
	defer s.mu.Unlock()

	res := make([]sysmon.Stats, 0, l)
	ring := s.s.Prev()
	for i := 0; i < l; i++ {
		res = append(res, ring.Value.(sysmon.Stats))
		ring = ring.Prev()
	}
	return res
}
