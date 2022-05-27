package model

import (
	"github.com/tendermint/tendermint/crypto/ed25519"
)

type User struct {
	Name          string
	PubKey        ed25519.PubKey `badgerhold:"index"` // this is just a wrapper around bytes
	Moderator     bool
	Banned        bool
	NumMessages   int64
	Version       uint64
	SchemaVersion int
}
