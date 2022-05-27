package model

import (
	"encoding/json"
	"errors"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/internal/jsontypes"
)

type User struct {
	Name      string
	PubKey    crypto.PubKey
	Moderator bool
	Banned    bool
	// schema versioning
	Version int
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name      string `json:"name"`
		PubKey    []byte `json:"pub_key"`
		Moderator bool   `json:"moderator"`
		Banned    bool   `json:"banned"`
		Version   int    `json:"version"`
	}{
		Name:      u.Name,
		PubKey:    jsontypes.Marshal(u.PubKey),
		Moderator: u.Moderator,
		Version:   u.Version,
	})
}

func (u *User) UnmarshalJSON(in []byte) error {
	intermediate := struct {
		Name      string `json:"name"`
		PubKey    []byte `json:"pub_key"`
		Moderator bool   `json:"moderator"`
		Banned    bool   `json:"banned"`
		Version   int    `json:"version"`
	}{}

	if err := json.Unmarshal(in, &intermediate); err != nil {
		return err
	}

	if u.Version != intermediate.Version {
		return errors.New("schema has changed")
	}

	if err := jsontypes.Unmarshal(intermediate.PubKey, u.PubKey); err != nil {
		return err
	}

	u.Name = intermediate.Name
	u.Moderator = intermediate.Moderator
	u.Banned = intermediate.Banned

	return nil
}
