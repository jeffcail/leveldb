# leveldb

LevelDB is a key-value store embedded database management system programming library developed by Google

# Introduce
Go operates leveldb, including Put, Get, Has, Delete, and SelectAll.

# Install



### Example

# InitClient

```go
db, err := leveldb.CreateLevelDB("./leveldb_data")
```

# Put
### k - v
```go
db.Put("test", "test")
```

### struct
```go
type Demo struct {
    Phone     string
    ChannelId int
    Reason    string
}
d := new(Demo)

d.Phone = "18745682512"
d.ChannelId = 1
d.Reason = "黑名单用户"

db.Put("key", &d)
```
### map
```go
a := make(map[string]interface{})
a["id"] = 1
a["name"] = "map"

db.Put("m", a)
```
### array
```go
array := [...]int{1, 2, 3}
db.Put("array", array)
```

### slice
```go
numbers := []int{0,1,2,3,4,5,6,7,8}
db.Put("n", numbers)
```

# Get
```go
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
```

# Has
```go
if ok, _ := db.Has("test"); !ok {
	fmt.Printf("key 为%s\n不存在", "test")
}
```

# Delete
```go
err = db.Delete("test")
```
# SelectAll
```go
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
```


