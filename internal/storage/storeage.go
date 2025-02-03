package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (kvStore *Store) Set(key, value string) {
	kvStore.mu.Lock()
	defer kvStore.mu.Unlock()
	kvStore.data[key] = value
}

func (kvStore *Store) Get(key string) (string, bool) {
	kvStore.mu.RLock()
	defer kvStore.mu.RUnlock()
	value, ok := kvStore.data[key]
	return value, ok
}

func (kvStore *Store) Delete(value string) bool {
	kvStore.mu.Lock()
	defer kvStore.mu.Unlock()
	_, ok := kvStore.data[value]
	delete(kvStore.data, value)
	return ok
}

func (kvStore *Store) SetWithTTL(key, value string, ttl time.Duration) {
	kvStore.Set(key, value)
	go func() {
		time.Sleep(ttl)
		kvStore.Delete(key)
	}()
}

func (kvStore *Store) SaveFile(fileName string) error {
	kvStore.mu.RLock()
	defer kvStore.mu.RUnlock()
	data, err := json.Marshal(kvStore.data)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func (kvStore *Store) LoadFile(fileName string) error {
	kvStore.mu.RLock()
	defer kvStore.mu.RUnlock()
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &kvStore.data)
}
