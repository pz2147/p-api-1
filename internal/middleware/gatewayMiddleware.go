package middleware

import (
	"net/http"
)

type GatewayMiddleware struct {
}

func NewGatewayMiddleware() *GatewayMiddleware {
	return &GatewayMiddleware{}
}

func (m *GatewayMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}
