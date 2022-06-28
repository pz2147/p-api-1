package middleware

import (
	"fmt"

	gzTrace "github.com/tal-tech/go-zero/core/trace"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type GatewayMiddleware struct {
}

func NewGatewayMiddleware() *GatewayMiddleware {
	return &GatewayMiddleware{}
}

func (m *GatewayMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		gCtx := r.Context()
		spanCtx := trace.SpanContextFromContext(gCtx)
		if spanCtx.HasSpanID() {
			fmt.Printf("[Handle] traceId %s\n", spanCtx.TraceID())
			//r.Header.Set(gzTrace.TraceIdKey, spanCtx.TraceID().String())
		}

		traceId := r.Header.Get(gzTrace.TraceIdKey)
		fmt.Printf("请求头 traceId %s\n", traceId)
		fmt.Printf("请求头 traceId %s\n", r.Header.Get("traceparent"))

		if len(traceId) > 0 {
			w.Header().Set(gzTrace.TraceIdKey, traceId)
		}

		spanCtx = trace.SpanContextFromContext(gCtx)
		if spanCtx.HasSpanID() {

			fmt.Printf("[Handle] traceId 2 %s\n", spanCtx.TraceID())
			//r.Header.Set(gzTrace.TraceIdKey, spanCtx.TraceID().String())
		}

		//r.Header.Get("device-type"),
		// Passthrough to next handler if need
		next(w, r)
	}
}
