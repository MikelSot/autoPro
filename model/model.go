package model

import "errors"

var (
	ErrClientCanNotBeNill = errors.New("El cliente no puede ser nula")
	ErrIDClientDoesNotExists = errors.New("El cliente no existe")
)



