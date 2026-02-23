package docs

import (
	"net/http"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/sholokhov-daniil/feedback-form/docs"
)

func RunSwagger(mux *http.ServeMux) {
	mux.Handle("GET /docs/swagger/", httpSwagger.WrapHandler)
	
}