package authhandlerv1

import (
	authv1 "github.com/sazonovItas/mocosso/gen/go/rest/v1/auth"
)

type authService interface{}

type AuthHandler struct {
	authSvc authService

	authv1.Unimplemented
}

func NewAuthHandler(authSvc authService) *AuthHandler {
	return &AuthHandler{
		authSvc: authSvc,
	}
}
