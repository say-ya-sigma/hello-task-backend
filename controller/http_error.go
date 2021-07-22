package controller

import (
	"golang.org/x/xerrors"
)

var (
	ErrHTTPNotFound = xerrors.New("not found")
	ErrHTTPInternalServerError = xerrors.New("internal server error")
	ErrHTTPUnauthorized = xerrors.New("unauthorized")
)

