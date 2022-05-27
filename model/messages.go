package model

import "github.com/tendermint/tendermint/crypto/ed25519"

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
