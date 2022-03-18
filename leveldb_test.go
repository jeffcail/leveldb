package leveldb1

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	db  *LevelDB
	err error
)

func init() {
	db, err = CreateLevelDB("./leveldb_data")
	if err != nil {
		panic(err)
	}
}

func TestLevelDB_Put(t *testing.T) {
	err = db.Put("test", "ceshiih")
	if err != nil {
		panic(err)
	}

	type Demo struct {
		Phone     string
		ChannelId int
		Reason    string
	}

	d := new(Demo)
	d.Phone = "18745682512"
	d.ChannelId = 1
	d.Reason = "黑名单用户"

	err = db.Put("46181632293601280", &d)
	if err != nil {
		panic(err)
	}
}

func TestLevelDB_Get(t *testing.T) {
	res, err := db.Get("test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("test: %s\n", res)

	res2, err := db.Get("46181632293601280")
	if err != nil {
		panic(err)
	}

	type Demo struct {
		Phone     string
		ChannelId int
		Reason    string
	}
	d := new(Demo)

	json.Unmarshal([]byte(res2), &d)

	fmt.Printf("Phone: %s\n", d.Phone)
	fmt.Printf("ChannelId: %d\n", d.ChannelId)
	fmt.Printf("Reason: %s\n", d.Reason)

}

func TestLevelDB_Has(t *testing.T) {
	if ok, _ := db.Has("test"); !ok {
		fmt.Printf("key 为%s\n不存在", "test")
	}
	fmt.Printf("key 为%s\n", "test")
}

func TestLevelDB_SelectAll(t *testing.T) {
	iter := db.SelectAll()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Printf("k: %s\n", k)
		fmt.Printf("v: %s\n", v)
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		panic(err)
	}
}

func TestLevelDB_Delete(t *testing.T) {
	err = db.Delete("test")
	if err != nil {
		panic(err)
	}
}
