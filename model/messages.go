package model

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto/ed25519"
)

type MsgSendMessage struct {
	Text string
	From ed25519.PubKey
}

type MsgSetBan struct {
	User  ed25519.PubKey
	State bool
}

type MsgSetModerator struct {
	User  ed25519.PubKey
	State bool
}

type MsgCreateUser struct {
	User User
}

func (db *DB) Process(message interface{}) error {
	switch msg := message.(type) {
	case MsgSendMessage:
		u, err := db.FindUser(msg.From)
		if err != nil {
			return err
		}
		// TODO: implement business logic

		u.Version++
		u.NumMessages++
		return db.SaveUser(u)
	case MsgSetBan:
		u, err := db.FindUser(msg.User)
		if err != nil {
			return err
		}

		u.Version++
		u.Banned = msg.State

		return db.SaveUser(u)
	case MsgSetModerator:
		u, err := db.FindUser(msg.User)
		if err != nil {
			return err
		}
		u.Version++
		u.Moderator = msg.State

		return db.SaveUser(u)
	case MsgCreateUser:
		msg.User.Version = 0
		return db.CreateUser(&msg.User)
	default:
		return fmt.Errorf("message type %T not supported", message)
	}
}
