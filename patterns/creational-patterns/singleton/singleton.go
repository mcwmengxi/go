package main

import (
	"fmt"
	"sync"
)

type DB struct {
	db string
}

var db *DB
var once sync.Once
func NewDB() *DB {
	once.Do(func() {
		db = &DB{
			db: "localhost",
		}
	})
	return db
}

func main() {

	d := NewDB()
	fmt.Printf("%p, %v\n", d, d.db)
  d = NewDB()
  fmt.Printf("%p, %v\n", d, d.db)
}
