/*
	simple lru
*/

package cache

import (
	"time"
)

type Cache struct {
	key     string
	value   interface{}
	exTime  time.Time
	expired bool
}

func (s *Store) GetCache(key string) (interface{}, bool) {
	if v, ok := s.m[key]; ok {
		c := v.(Cache)
		// if key has expired
		if c.expired && c.exTime.Before(time.Now()) {
			s.RemoveCache(key)
			return nil, false
		}
		s.Lock()
		defer s.Unlock()
		s.MoveFront(c)
		return c.value, true
	}

	return nil, false
}

func (s *Store) SetCache(key string, value interface{}) bool {
	s.Lock()
	defer s.Unlock()
	c := Cache{key: key, value: value}
	if _, ok := s.m[key]; !ok {
		s.m[c.key] = c
		s.l.PushFront(c.key)
		if s.l.Len() > s.maxLength {
			delete(s.m, s.l.Back().Value.(string))
			s.l.Remove(s.l.Back())
		}
	} else {
		s.m[c.key] = c
		s.MoveFront(c)
	}
	return true
}

// set a key expired time
func (s *Store) SetExpired(key string, duration time.Duration) bool {
	s.Lock()
	defer s.Unlock()
	if v, ok := s.m[key]; ok {
		c := v.(Cache)
		c.exTime = time.Now().Add(duration)
		c.expired = true
		s.m[c.key] = c
		return true
	}
	return false
}

// set a key which has expired time
// duration: ns
func (s *Store) SetExpiredCache(key string, value interface{}, duration time.Duration) bool {
	s.Lock()
	defer s.Unlock()
	c := Cache{key: key, value: value, exTime: time.Now().Add(duration), expired: true}
	if _, ok := s.m[key]; !ok {
		s.m[c.key] = c
		s.l.PushFront(c.key)
		if s.l.Len() > s.maxLength {
			delete(s.m, s.l.Back().Value.(string))
			s.l.Remove(s.l.Back())
		}
	} else {
		s.m[c.key] = c
		s.MoveFront(c)
	}
	return true
}

func (s *Store) MoveFront(c Cache) {
	for e := s.l.Front(); e != nil; e = e.Next() {
		if e.Value == c.key {
			s.l.MoveToFront(e)
		}
	}
}

func (s *Store) RemoveCache(key string) bool {
	delete(s.m, key)
	for e := s.l.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			s.l.Remove(e)
			return true
		}
	}
	return false
}

func (s *Store) Clear() bool {
	s.Lock()
	defer s.Unlock()
	for e := s.l.Front(); e != nil; {
		temp := e
		e = e.Next()
		delete(s.m, temp.Value.(string))
		s.l.Remove(temp)
	}
	return true
}
