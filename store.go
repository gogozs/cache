package cache

import (
	"container/list"
	"sync"
)

const (
	defaultLength = 100000
)

type Store struct {
	l         *list.List
	m         map[string]interface{}
	maxLength int

	sync.RWMutex
}

type (
	storeOption struct {
		length int
	}
	IStoreOption interface {
		apply(*storeOption)
	}
	lengthOption int
)

func (l lengthOption) apply(o *storeOption) {
	o.length = int(l)
}

func WithLength(length int) IStoreOption {
	return lengthOption(length)
}

func NewStore(options ...IStoreOption) *Store {
	initialOption := storeOption{length: defaultLength}
	for _, o := range options {
		o.apply(&initialOption)
	}
	return &Store{
		l:         list.New(),
		m:         make(map[string]interface{}),
		maxLength: initialOption.length,
	}
}

func (s *Store) Dump() (*list.List, map[string]interface{}) {
	return s.copyList(), s.copyMap()
}

func (s *Store) copyList() *list.List {
	var l list.List
	n := s.l.Front()
	for n != nil {
		l.PushBack(n)
	}
	return &l
}

func (s *Store) copyMap() map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range s.m {
		m[k] = v
	}
	return m
}

func BuildStore(q *list.List, m map[string]interface{}, options ...IStoreOption) *Store {
	if m == nil {
		m = make(map[string]interface{})
	}
	initialOption := storeOption{length: defaultLength}
	for _, o := range options {
		o.apply(&initialOption)
	}
	return &Store{
		l:         q,
		m:         m,
		maxLength: initialOption.length,
	}
}
