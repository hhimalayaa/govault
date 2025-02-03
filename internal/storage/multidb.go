package storage

import (
	"sync"
)

// MultiDB manages multiple logical databases
type MultiDB struct {
	mu     sync.RWMutex
	dbs    map[int]*Store
	active int
}

func NewMultiDB() *MultiDB {
	return &MultiDB{
		dbs:    map[int]*Store{0: NewStore()}, // Default DB 0
		active: 0,
	}
}

func (mdb *MultiDB) SelectDB(dbNum int) {
	mdb.mu.Lock()
	defer mdb.mu.Unlock()
	if _, exists := mdb.dbs[dbNum]; !exists {
		mdb.dbs[dbNum] = NewStore()
	}
	mdb.active = dbNum
}

func (mdb *MultiDB) ActiveDB() *Store {
	mdb.mu.RLock()
	defer mdb.mu.RUnlock()
	return mdb.dbs[mdb.active]
}
