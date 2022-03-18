package leveldb1

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

var (
	ErrEmptyKey = errors.New("key could not be empty")
)

type LevelDB struct {
	db *leveldb.DB
}

// CreateLevelDB
func CreateLevelDB(path string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(path, nil)
	}
	if err != nil {
		return nil, err
	}
	result := &LevelDB{
		db: db,
	}
	return result, nil
}

// Get
func (db *LevelDB) Get(key string) ([]byte, error) {
	return db.db.Get([]byte(key), nil)
}

// Put
func (db *LevelDB) Put(key string, value interface{}) error {
	if len(key) < 1 {
		return ErrEmptyKey
	}
	res, _ := json.Marshal(value)
	return db.db.Put([]byte(key), []byte(res), nil)
}

// Has
func (db *LevelDB) Has(key string) (bool, error) {
	return db.db.Has([]byte(key), nil)
}

// Delete
func (db *LevelDB) Delete(key string) error {
	return db.db.Delete([]byte(key), nil)
}

func (db *LevelDB) SelectAll() iterator.Iterator {
	return db.db.NewIterator(nil, nil)
}
