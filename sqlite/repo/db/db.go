package db

import "goExploration/sqlite/task"

// Store is a SQLite storage.
// It implements task.repository interface.
type Store struct {
	dbpath, pragma string
}

func Init(path, params string) Store {
	return Store{
		dbpath: path,
		pragma: params,
	}
}

func (s *Store) Save(t task.Item) (ID int) {
}
