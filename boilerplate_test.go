package boltdbboilerplate

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	initTest() //setup
	retCode := m.Run()
	cleanBoltDB() //Teardown
	os.Exit(retCode)
}

func initTest() {
	log.Println("Init BoltDB")

	buckets := []string{"owner", "sensors", "beacons"}
	err := InitBolt("./test.boltdb", buckets)
	log.Println("After init db")

	if err != nil {
		log.Fatal("Can't init boltDB")
	}
}

func cleanBoltDB() {
	os.Remove("./test.boltdb")
}

func TestInitBolt(t *testing.T) {
	_, err := os.Stat("./test.boltdb")
	assert.Nil(t, err, "test.boltdb should exist")
}

func TestPut(t *testing.T) {
	err := Put([]byte("owner"), []byte("owner"), []byte("testUserName"))
	assert.Nil(t, err, "There should be no errors putting value")

	value := Get([]byte("owner"), []byte("owner"))
	assert.Equal(t, "testUserName", string(value), "value should be testUserName")

	value = Get([]byte("owner"), []byte("nil"))
	assert.Nil(t, value, "value should be nil")
}

func TestDelete(t *testing.T) {
	err := Put([]byte("owner"), []byte("owner"), []byte("testUserName"))

	err = Delete([]byte("owner"), []byte("owner"))
	assert.Nil(t, err, "There should be no errors deleting")

	value := Get([]byte("owner"), []byte("owner"))
	assert.Nil(t, value, "value should be nil")
}

func TestGetAllKeys(t *testing.T) {
	err := Put([]byte("sensors"), []byte("key1"), []byte("value1"))
	err = Put([]byte("sensors"), []byte("key2"), []byte("value2"))

	keys := GetAllKeys([]byte("sensors"))
	assert.Nil(t, err, "There should be no errors getting all keys")

	assert.Equal(t, 2, len(keys), "there should be 2 keys")

	assert.Equal(t, "key1", string(keys[0]), "first key should be key1")
	assert.Equal(t, "key2", string(keys[1]), "second key should be key2")
}

func TestGetAllKeyValuePairs(t *testing.T) {
	err := Put([]byte("sensors"), []byte("key1"), []byte("value1"))
	err = Put([]byte("sensors"), []byte("key2"), []byte("value2"))

	pairs := GetAllKeyValues([]byte("sensors"))
	assert.Nil(t, err, "There should be no errors getting all boltpairs")

	assert.Equal(t, 2, len(pairs), "there should be 2 paris")

	assert.Equal(t, "key1", string(pairs[0].Key), "key of first pair should be key1")
	assert.Equal(t, "value1", string(pairs[0].Value), "value of first pair should be value1")

	assert.Equal(t, "key2", string(pairs[1].Key), "key of first pair should be key2")
	assert.Equal(t, "value2", string(pairs[1].Value), "value of first pair should be value2")
}
