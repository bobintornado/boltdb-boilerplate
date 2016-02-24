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

# Docs

[GoDoc](https://godoc.org/github.com/bobintornado/boltdb-boilerplate)
