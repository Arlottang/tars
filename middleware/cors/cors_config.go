package cors

import (
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/gorilla/handlers"
)

func ServerOption() http.ServerOption {
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	return http.Filter(corsHandler)
}
