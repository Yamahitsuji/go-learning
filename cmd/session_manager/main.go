package main

import (
	"fmt"
	"sync"
)

type Session struct {
	tx bool
	ops []string
}

type SessionManager struct {
	mu sync.Mutex
	sessions map[interface{}]*Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		mu: sync.Mutex{},
		sessions: make(map[interface{}]*Session),
	}
}

func (m *SessionManager) Begin(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.sessions[key]
	if ok && s.tx {
		return
	}
	m.sessions[key] = &Session{
		tx: true,
		ops: make([]string, 0),
	}
}

func (m *SessionManager) Commit(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.sessions[key]
	if !ok {
		return
	}

	for _, v := range s.ops {
		fmt.Println(v)
	}
	delete(m.sessions, key)
}

func (m *SessionManager) Rollback(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.sessions, key)
}

func (m *SessionManager) Do(key, query string) {
	s, ok := m.sessions[key]
	if !ok || s == nil || !s.tx {
		fmt.Println(query)
		return
	}

	s.ops = append(s.ops, query)
}

func main() {
	manager := NewSessionManager()
	key1 := "key1"
	key2 := "key2"
	key3 := "key3"

	manager.Begin(key1)
	manager.Do(key1, "query1-1")
	manager.Do(key1, "query1-2")
	manager.Do(key1, "query1-3")
	manager.Commit(key1)

	manager.Begin(key2)
	manager.Do(key2, "query2-1")
	manager.Do(key2, "query2-2")
	manager.Rollback(key2)

	manager.Do(key3, "query3-1")
}
