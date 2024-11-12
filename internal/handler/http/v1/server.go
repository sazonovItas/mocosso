package httpserverv1

import (
	"github.com/go-chi/chi/v5"
	authv1 "github.com/sazonovItas/mocosso/gen/go/rest/v1/auth"
	httputils "github.com/sazonovItas/mocosso/internal/handler/http/utils"
)

func RegisterServer(router chi.Router, authHandler authv1.ServerInterface) {
	authv1.HandlerWithOptions(authHandler, authv1.ChiServerOptions{
		BaseRouter:       router,
		BaseURL:          "/au",
		ErrorHandlerFunc: httputils.DefaultErrorHandlerFunc,
	})
}
