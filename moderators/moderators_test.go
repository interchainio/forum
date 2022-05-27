package moderators_test

import (
	"bytes"
	"testing"

	"github.com/interchainio/forum/model"
	"github.com/interchainio/forum/moderators"
)

func TestAddModerator(t *testing.T) {
	pk1, pk2 := []byte("one"), []byte("two")
	s := moderators.NewSet()
	u := model.User{
		PubKey: pk1,
	}
	s.Add(u)
	u = model.User{
		PubKey: pk2,
	}
	s.Add(u)
	l := s.List()
	if !bytes.Equal(l[0].PubKey, pk1) {
		t.Fatalf("expected the list to contain the added public key")
	}
	if !bytes.Equal(l[1].PubKey, pk2) {
		t.Fatalf("expected the list to contain the added public key")
	}
}

func TestRemoveModerator(t *testing.T) {
	pk1, pk2 := []byte("one"), []byte("two")
	s := moderators.NewSet()
	u1 := model.User{
		PubKey: pk1,
	}
	if !s.Add(u1) {
		t.Fatal("expected add to succeed for user")
	}
	u2 := model.User{
		PubKey: pk2,
	}
	if !s.Add(u2) {
		t.Fatal("expected add to succeed for user")
	}
	l := s.List()
	if !bytes.Equal(l[0].PubKey, pk1) {
		t.Fatalf("expected the list to contain the added public key")
	}
	if !bytes.Equal(l[1].PubKey, pk2) {
		t.Fatalf("expected the list to contain the added public key")
	}

	if !s.Remove(u2) {
		t.Fatal("expected remove to succeed for user")
	}
	if len(s.List()) != 1 {
		t.Fatal("expected List to return a list with one element")
	}
	if !bytes.Equal(s.List()[0].PubKey, u1.PubKey) {
		t.Fatal("expected List to contain the expected user")
	}
}
