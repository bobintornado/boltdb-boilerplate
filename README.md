# boltdb-boilerplate
boltdb-boilerplate is a simple & stupid boilerplate project wrapping around [boltdb](https://github.com/boltdb/bolt), and aim to make simple calls as one-liners.

# Methods Summary

* `InitBolt`: Init the database
* `Close`: Close the database
* `Get`: Retrieve a value by key
* `Put`: Put a key/value pair into a bucket
* `Delete`: Delete a key/value pair from a bucket
* `GetAllKeys`: Get all keys from a bucket in `[][]byte` fromat
* `GetAllKeyValues`: Get all key/value pairs from a bucket in `[]BoltPair` fromat

# Quick Demo
```go
// import
import "github.com/bobintornado/boltdbboilerplate"

// Init
buckets := []string{"ownerBucket", "sensors"}

err := boltdbboilerplate.InitBolt("./database.boltdb", buckets)
if err != nil {
  log.Fatal("Can't init boltDB")
}

// Put
err = boltdbboilerplate.Put([]byte("ownerBucket"), []byte("ownerKey"), []byte("username"))

// Get owner 
value := boltdbboilerplate.Get([]byte("ownerBucket"), []byte("ownerKey"))

// Delete
err = boltdbboilerplate.Delete([]byte("ownerBucket"), []byte("ownerKey"))

// Insert two key/value
err = boltdbboilerplate.Put([]byte("sensors"), []byte("key1"), []byte("value1"))
err = boltdbboilerplate.Put([]byte("sensors"), []byte("key2"), []byte("value2"))

// Get all keys
keys := boltdbboilerplate.GetAllKeys([]byte("sensors"))
// keys = [key1, key2]

// Get all key/value pairs
pairs := boltdbboilerplate.GetAllKeyValues([]byte("sensors"))
// pairs = [{Key:key1, Value:value1}, {Key: key2, Value:value2}]

// Close
boltdbboilerplate.Close()
```

# Docs

[GoDoc](https://godoc.org/github.com/bobintornado/boltdb-boilerplate)
