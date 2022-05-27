package model

import (
	"github.com/tendermint/tendermint/crypto/ed25519"
)

type User struct {
	Name      string
	PubKey    ed25519.PubKey // this is just a wrapper around bytes
	Moderator bool
	Banned    bool
	// schema versioning
	Version int
}
