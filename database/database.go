package database

import (
	"encoding/json"
	"log"
	"sort"

	badger "github.com/dgraph-io/badger/v2"
)

// Database holds a badget db, duh
type Database struct {
	db *badger.DB
}

// NewDatabase returns a new or existing db depending on if a db exists at dbPath or not
func NewDatabase(dbPath string) *Database {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.Fatal(err)
	}

	return &Database{db: db}
}

type User struct {
	ID     string
	Points int32
}

func (d *Database) SetUser(u User) error {
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}
	if err := d.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(u.ID), b)
		return err
	}); err != nil {
		return err
	}
	return nil
}

func (d *Database) GetUser(id string) (User, error) {
	var u User
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			return json.Unmarshal(val, &u)
		})
		return err
	})
	return u, err
}

func (d *Database) GetTopUsers(limit int) ([]User, error) {
	var users []User
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			var u User
			err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &u)
			})
			if err != nil {
				return err
			}
			users = append(users, u)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Points > users[j].Points
	})

	if len(users) > limit {
		users = users[:limit]
	}

	return users, nil
}
