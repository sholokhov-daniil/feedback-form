package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// Chain создает цепочку middleware
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
    // Применяем middleware в обратном порядке (первый в списке = внешний слой)
    for i := len(middlewares) - 1; i >= 0; i-- {
        h = middlewares[i](h)
    }
    return h
}

// Комбинируем несколько middleware в один
func Compose(middlewares ...Middleware) Middleware {
    return func(next http.Handler) http.Handler {
        return Chain(next, middlewares...)
    }
}