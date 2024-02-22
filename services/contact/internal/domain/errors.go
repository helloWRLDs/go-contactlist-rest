package domain

import "errors"

var (
	ErrEmptyNameFields  = errors.New("domain: Name fields cannot be empty")
	ErrWrongPhoneFormat = errors.New("domain: Wrong phone number format")
)
