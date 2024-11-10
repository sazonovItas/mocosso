package serverv1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	authv1 "github.com/sazonovItas/mocosso/gen/go/rest/v1/auth"
	httputils "github.com/sazonovItas/mocosso/internal/handler/http/utils"
)

type Server struct {
	http.Handler

	authHandler authv1.ServerInterface
}

func NewServer(router chi.Router, authHandler authv1.ServerInterface) *Server {
	router.Route("/v1", func(r chi.Router) {
		authv1.HandlerWithOptions(authHandler, authv1.ChiServerOptions{
			BaseURL:          "/au",
			BaseRouter:       r,
			ErrorHandlerFunc: httputils.DefaultErrorHandlerFunc,
		})
	})

	return &Server{
		Handler: router,

		authHandler: authHandler,
	}
}
