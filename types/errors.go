package types

import (
	"fmt"
)

const (
	codeKeyNotFound      = 1
	codeWrongPassword    = 2
	codeKeyAlreadyExists = 3
)

type keybaseError interface {
	error
	Code() int
}

type errWrongPassword struct {
	code int
}

func (e errWrongPassword) Code() int {
	return e.code
}

func (e errWrongPassword) Error() string {
	return "invalid account password"
}

func NewErrWrongPassword() error {
	return errWrongPassword{
		code: codeWrongPassword,
	}
}

func IsErrWrongPassword(err error) bool {
	if err == nil {
		return false
	}
	if keyErr, ok := err.(keybaseError); ok {
		if keyErr.Code() == codeWrongPassword {
			return true
		}
	}
	return false
}

type errKeyNotFound struct {
	code int
	name string
}

func (e errKeyNotFound) Code() int {
	return e.code
}

func (e errKeyNotFound) Error() string {
	return fmt.Sprintf("Key %s not found", e.name)
}

func NewErrKeyNotFound(name string) error {
	return errKeyNotFound{
		code: codeKeyNotFound,
		name: name,
	}
}

func IsErrKeyNotFound(err error) bool {
	if err == nil {
		return false
	}
	if keyErr, ok := err.(keybaseError); ok {
		if keyErr.Code() == codeKeyNotFound {
			return true
		}
	}
	return false
}

type errKeyAlreadyExists struct {
	code int
	name string
}

func (e errKeyAlreadyExists) Code() int {
	return e.code
}

func (e errKeyAlreadyExists) Error() string {
	return fmt.Sprintf("Key %s already exists", e.name)
}

func NewErrKeyAlreadyExists(name string) error {
	return errKeyAlreadyExists{
		code: codeKeyAlreadyExists,
		name: name,
	}
}

func IsErrKeyAlreadyExists(err error) bool {
	if err == nil {
		return false
	}
	if keyErr, ok := err.(keybaseError); ok {
		if keyErr.Code() == codeKeyAlreadyExists {
			return true
		}
	}
	return false
}
