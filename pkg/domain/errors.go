package domain

import "errors"

var (
	ErrIndexNotFound     = errors.New("index not found")
	ErrStackUnderflow    = errors.New("stack underflow")
	ErrUndefinedVarIndex = errors.New("undefined variable index")
)

var (
	NotParsed = errors.New("not parsed")
)
