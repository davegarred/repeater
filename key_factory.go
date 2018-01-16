package main

import "github.com/google/uuid"

type KeyFactory interface {
	Get() string
}

func (f *DefaultKeyFactory) Get() string {
	return uuid.New().String()
}

func NewKeyFactory() *DefaultKeyFactory {
	return &DefaultKeyFactory{}
}

type DefaultKeyFactory struct {
}
