package boltdbboilerplate

import (
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var database string
var fileMode os.FileMode = 0600 // owner can read and write
var db *bolt.DB

// InitBolt inits bolt database. Create the file if not exist.
// By default it opens the file in 0600 mode, with a 10 seconds timeout period
func InitBolt(path string, buckets []string) error {
	log.Println("Trying to open databse")
	database = path
	var err error
	// open the target file, file mode fileMode, and a 10 seconds timeout period
	db, err = bolt.Open(database, fileMode, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Trying to create buckets")
	err = db.Update(func(tx *bolt.Tx) error {
		for _, value := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(value))
			if err != nil {
				return err
			}
		}

		return nil
	})

	log.Println("Trying to return error")
	return err
}

// Close bolt db
func Close() {
	db.Close()
}

// Get value from bucket by key
func Get(bucket []byte, key []byte) []byte {
	var value []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get(key)
		if v != nil {
			value = append(value, b.Get(key)...)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return value
}

// Put a key/value pair into target bucket
func Put(bucket []byte, key []byte, value []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put(key, value)
		return err
	})

	return err
}

// Delete a key from target bucket
func Delete(bucket []byte, key []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Delete(key)
		return err
	})

	return err
}

// GetAllKeys get all keys from the target bucket
func GetAllKeys(bucket []byte) [][]byte {
	var keys [][]byte

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		b.ForEach(func(k, v []byte) error {
			// Due to
			// Byte slices returned from Bolt are only valid during a transaction. Once the transaction has been committed or rolled back then the memory they point to can be reused by a new page or can be unmapped from virtual memory and you'll see an unexpected fault address panic when accessing it.
			// We copy the slice to retain it
			dst := make([]byte, len(k))
			copy(dst, k)

			keys = append(keys, dst)
			return nil
		})
		return nil
	})

	return keys
}

// BoltPair is a struct to store key/value pair data
type BoltPair struct {
	Key   []byte
	Value []byte
}

// GetAllKeyValues get all key/value pairs from a bucket in BoltPair struct format
func GetAllKeyValues(bucket []byte) []BoltPair {
	var pairs []BoltPair

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		b.ForEach(func(k, v []byte) error {
			// Due to
			// Byte slices returned from Bolt are only valid during a transaction. Once the transaction has been committed or rolled back then the memory they point to can be reused by a new page or can be unmapped from virtual memory and you'll see an unexpected fault address panic when accessing it.
			// We copy the slice to retain it
			dstk := make([]byte, len(k))
			dstv := make([]byte, len(v))
			copy(dstk, k)
			copy(dstv, v)

			pair := BoltPair{dstk, dstv}
			pairs = append(pairs, pair)
			return nil
		})

		return nil
	})

	return pairs
}
