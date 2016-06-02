package main

import (
	"log"

	"github.com/boltdb/bolt"
)

func addKey(key, value []byte) error {
	log.Printf("Received new key: %s with value: %s", key, value)

	addKeyErr := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(dbBucket)
		err := b.Put(key, value)
		return err
	})
	if addKeyErr != nil {
		log.Fatalf("Write key %s failed - %v", key, addKeyErr)
	}
	return nil
}

func findKey(key []byte) []byte {
	var value []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(dbBucket)
		value = b.Get(key)
		return nil
	})
	log.Print(err)
	return value
}
