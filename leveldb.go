package leveldb1

import (
	"encoding/json"

	"github.com/syndtr/goleveldb/leveldb/util"

	"github.com/syndtr/goleveldb/leveldb/iterator"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

var (
	ErrEmptyKey = errors.New("key could not be empty")
)

type Database interface {
	Put(key string, value interface{}) error
	Get(key string) ([]byte, error)
	Has(key string) (bool, error)
	Delete(key string) error
	SelectAll() iterator.Iterator
	SelectPrefixSubsetKeyAll(key string) ([]map[string]interface{}, error)
	CountPrefixSubsetKey(key string) (int64, error)
	CountAll() (int64, error)
	DeletePrefixSubsetKey(key string) (bool, error)
}

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

// DeleteAll
func (db *LevelDB) DeleteAll() (bool, error) {
	iter := db.db.NewIterator(nil, nil)
	for iter.Next() {
		if err := db.db.Delete(iter.Key(), nil); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (db *LevelDB) DeletePrefixSubsetKey(key string) (bool, error) {
	iter := db.db.NewIterator(util.BytesPrefix([]byte(key)), nil)
	for iter.Next() {
		if err := db.db.Delete(iter.Key(), nil); err != nil {
			return false, err
		}
	}
	return true, nil
}

// SelectAll
func (db *LevelDB) SelectAll() iterator.Iterator {
	return db.db.NewIterator(nil, nil)
}

// SelectPrefixSubsetKeyAll
// 取出所有指定前缀 key 的数据
func (db *LevelDB) SelectPrefixSubsetKeyAll(key string) ([]map[string]interface{}, error) {
	m := make(map[string]interface{})
	ms := make([]map[string]interface{}, 0)

	iter := db.db.NewIterator(util.BytesPrefix([]byte(key)), nil)
	for iter.Next() {
		k := string(iter.Key())
		v := string(iter.Value())
		m[k] = v
		ms = append(ms, m)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, err
	}
	return ms, nil
}

// CountPrefixSubsetKey
// 计算指定前缀 key 的数量
func (db *LevelDB) CountPrefixSubsetKey(key string) (int64, error) {
	var sum int64 = 0
	iter := db.db.NewIterator(util.BytesPrefix([]byte(key)), nil)
	for iter.Next() {
		sum++
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return 0, err
	}
	return sum, nil
}

// CountAll
func (db *LevelDB) CountAll() (int64, error) {
	var sum int64 = 0
	iter := db.db.NewIterator(nil, nil)
	for iter.Next() {
		sum++
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return 0, err
	}
	return sum, nil
}
