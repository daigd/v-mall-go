package repository

import (
	"sync"
)

type userRepository struct {
	mu sync.RWMutex
}

func (r *userRepository) Exec(query Query, action Query, limit int, mode int) (ok bool) {
	if ReadOnlyMode == mode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	return false
}
