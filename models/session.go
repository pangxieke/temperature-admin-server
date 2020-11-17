package models

import (
	"time"

	"github.com/pkg/errors"
)

type Session struct {
	StoreID int
	Token   JWT
	Exp     time.Time
	Store   *Store
}

var Subject = "temp-admin"

func NewSession(store *Store) (s *Session, err error) {
	exp := time.Now().Add(24 * time.Hour)
	jwt, err := NewJWT(secret, Subject, store.Id, &exp, nil)
	if err != nil {
		return
	}
	s = &Session{StoreID: store.Id, Store: store, Token: *jwt, Exp: exp}
	return
}

func LoadSession(token string) (s *Session, err error) {
	jwt, err := ParseJWT(secret, token)
	if err != nil {
		return
	}
	Uid := int(jwt.Claims["uid"].(float64))
	storeDao := StoreDao{}
	store, err := storeDao.FindById(Uid)
	if err != nil {
		return
	} else if store.Id == 0 {
		return nil, errors.New("store not found")
	}
	return &Session{StoreID: store.Id, Token: *jwt, Store: &store}, nil
}

func DelSession(token string) (s *Session, err error) {
	exp := time.Now().Add(-1 * time.Hour)
	jwt, err := ParseJWT(secret, token)
	if err != nil {
		return
	}
	s = &Session{Token: *jwt, Exp: exp}
	return
}
