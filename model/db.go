package model

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"sync"

	"github.com/dgraph-io/badger"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/timshannon/badgerhold"
)

var bufPool = sync.Pool{New: func() any { return &bytes.Buffer{} }}

type DB struct {
	store *badgerhold.Store
}

func New(dbPath string) (*DB, error) {
	store, err := badgerhold.Open(
		badgerhold.Options{
			Encoder: func(v interface{}) ([]byte, error) {
				jby, err := json.Marshal(v)
				if err != nil {
					return nil, err
				}

				buf := bufPool.Get().(*bytes.Buffer)
				defer bufPool.Put(buf)
				buf.Reset()
				gz, err := flate.NewWriter(buf, 5)
				if err != nil {
					return nil, err
				}
				defer gz.Close()

				if _, err := gz.Write(jby); err != nil {
					return nil, err
				}

				return buf.Bytes(), nil
			},
			Decoder: func(in []byte, val interface{}) error {
				buf := bufPool.Get().(*bytes.Buffer)
				defer bufPool.Put(buf)
				buf.Reset()

				gz := flate.NewReader(buf)
				defer gz.Close()

				if _, err := gz.Read(in); err != nil {
					return err
				}
				return json.Unmarshal(buf.Bytes(), val)
			},
			Options: badger.DefaultOptions(dbPath),
		},
	)
	if err != nil {
		return nil, err
	}

	return &DB{
		store: store,
	}, nil

}

func (db *DB) Close() error                               { return db.store.Close() }
func (db *DB) CreateUser(u *User) error                   { return db.store.Insert(u.PubKey, u) }
func (db *DB) SaveUser(u *User) error                     { return db.store.Update(u.PubKey, u) }
func (db *DB) FindUser(key ed25519.PubKey) (*User, error) { return nil, nil }
