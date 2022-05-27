package model

import (
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/internal/jsontypes"
)

type User struct {
	Name      string
	PubKey    crypto.PubKey
	Moderator bool
}

func (u *User) TypeTag() string { return "io.interchain.forum" }
func init()                     { jsontypes.MustRegister(&User{}) }
