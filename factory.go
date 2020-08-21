package main

import (
	"net/url"

	"github.com/google/uuid"
)

type Generator func() string

var DefaultGenerator = func() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
type Factory struct {
	store     Store
	generator Generator
}
func NewFactory(generator Generator, store Store) *Factory {
	return &Factory{
		store:     store,
		generator: generator,
	}
}
func (f *Factory) Gen(uri string) (key string, err error) {
	_, err = url.ParseRequestURI(uri)
	if err != nil {
		return "", err
	}
	key = f.generator()
	for {
		if v := f.store.Get(key); v == "" {
			break
		}
		key = f.generator()
	}
	return key, nil
}
