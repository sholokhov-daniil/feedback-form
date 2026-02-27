package docs

import (
	_ "github.com/sholokhov-daniil/feedback-form/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func RunSwagger(mux *http.ServeMux) {
	mux.Handle("GET /docs/swagger/", httpSwagger.WrapHandler)

}
