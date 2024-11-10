package authhandlerv1

import (
	authv1 "github.com/sazonovItas/mocosso/gen/go/rest/v1/auth"
)

type AuthHandler struct {
	authv1.Unimplemented
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}
