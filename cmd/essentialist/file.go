package main

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"

	"github.com/essentialist-app/essentialist/internal"
)

type uriDeckAccessor struct {
	deck fyne.URI
	db   fyne.URI
}

func (u *uriDeckAccessor) CardsReader() (io.ReadCloser, error) {
	return storage.Reader(u.deck)
}

func (u *uriDeckAccessor) MetaReader() (io.ReadCloser, error) {
	r, err := storage.Reader(u.db)
	if err != nil {
		return nil, err
	}
	return r, err
}

func (u *uriDeckAccessor) MetaWriter() (io.WriteCloser, error) {
	w, err := storage.Writer(u.db)
	if err != nil {
		return nil, err
	}
	return w, err
}

func (u *uriDeckAccessor) DeckName() string {
	return u.deck.Name()
}

func (u *uriDeckAccessor) Path() string {
	return u.deck.Path()
}

func NewDeckAccessor(deck, db fyne.URI) internal.DeckAccessor {
	return &uriDeckAccessor{
		deck: deck,
		db:   db,
	}
}
