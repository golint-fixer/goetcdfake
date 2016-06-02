package main

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

var (
	db         *bolt.DB
	dbFileName = "etcdfake.db"
	dbBucket   = []byte("myBuc")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func openDB() error {
	return db.Update(func(tx *bolt.Tx) error {
		_, createBuckerErr := tx.CreateBucketIfNotExists(dbBucket)
		return createBuckerErr
	})
}

func main() {
	var openDBErr error
	db, openDBErr = bolt.Open(dbFileName, 0644, nil)
	if openDBErr != nil {
		log.Fatal(openDBErr)
	}
	defer func() {
		dbCloseErr := db.Close()
		if dbCloseErr != nil {
			log.Fatal(dbCloseErr)
		}
	}()
	if dbErr := openDB(); dbErr != nil {
		log.Fatalf("Open DB error: %v", dbErr)
	}

	e := echo.New()

	setRoutes(e)

	e.SetDebug(true)
	e.Run(standard.New(":9732"))
}
