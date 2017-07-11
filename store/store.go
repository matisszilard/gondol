package store

// Store interface for all of the supported model Stores
type Store struct {
	Users UserStore `inject:""`
}

var instance Store

// Load return a store implenetation
func Load() *Store {
	return &instance
}
