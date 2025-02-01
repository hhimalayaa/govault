package storage

import "sync"

type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

var kvStore = Store{
	data: make(map[string]string),
}

func Set(key, value string) {
	kvStore.mu.Lock()
	defer kvStore.mu.Unlock()
	kvStore.data[key] = value
}

func Get(key string) (string, bool) {
	kvStore.mu.RLock()
	defer kvStore.mu.RUnlock()
	value, ok := kvStore.data[key]
	return value, ok
}

func Delete(value string) bool {
	kvStore.mu.Lock()
	defer kvStore.mu.Unlock()
	_, ok := kvStore.data[value]
	delete(kvStore.data, value)
	return ok
}
