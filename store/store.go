package store

import (
	"sync"

	"github.com/matisszilard/gondol/store/rethinkstore"
)

// Store interface for all of the supported model Stores
type Store interface {
	Users() UserStore
}

type store struct {
	name string
	impl *rethinkstore.Store
}

var instance *store
var once sync.Once

// getInstance return a store implementation
func getImplementation() *store {
	once.Do(func() {
		instance = &store{name: "gondol-rethink-db"}
		instance.impl = rethinkstore.Load()
	})
	return instance
}

func (s *store) Users() UserStore {
	return s.impl.Users
}
