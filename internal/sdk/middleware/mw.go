package middleware

import "net/http"

// Middleware is a function that wraps an http.Handler.
type Middleware func(http.Handler) http.Handler

// ChainMiddleware chains multiple middlewares together.
// The first middleware in the slice is the first to be executed.
func WrapMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
	// Apply middlewares in reverse order
	for i := len(middlewares) - 1; i >= 0; i-- {
		if middlewares[i] != nil {
			h = middlewares[i](h)
		}
	}
	return h
}

// package middleware

// import "net/http"

// // Middleware is a function that wraps an http.Handler.
// type Middleware func(http.Handler) http.Handler

// // ChainMiddleware chains multiple middlewares together.
// func mwFunc(h http.Handler, middlewares ...Middleware) http.Handler {
// 	for _, m := range middlewares {
// 		h = m(h)
// 	}
// 	return h
// }
