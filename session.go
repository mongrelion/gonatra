package gonatra

import (
    "sync"
)

type session struct {
    sync.RWMutex
    m map[string]string
}

// Retrieve a value from session
func (s *session) Get(k string) (val string) {
    s.RLock()
    val, _ = s.m[k]
    s.RUnlock()
    return
}

// Set a value in session
func (s *session) Set(k, v string) string {
    s.Lock()
    s.m[k] = v
    s.Unlock()
    return v
}
