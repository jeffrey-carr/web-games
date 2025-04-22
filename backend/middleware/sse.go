package middleware

import "net/http"

// SSE represents SSE middleware
type SSE struct{}

// NewSSEMiddleware creates a new SSE Middleware
func NewSSEMiddleware() SSE {
	return SSE{}
}

// GetName prints the name of the middleware
func (sse SSE) GetName() string {
	return "SSE"
}

// Apply applies the SSE middleware
func (sse SSE) Apply(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	return &sseResponseWriter{ResponseWriter: w}, r
}

// sseResponseWriter wraps an http.ResponseWriter to provide SSE-specific functionality
type sseResponseWriter struct {
	http.ResponseWriter
}

// Write overrides the underlying ResponseWriter's Write method to ensure flushing
func (w *sseResponseWriter) Write(data []byte) (int, error) {
	n, err := w.ResponseWriter.Write(data)
	if flusher, ok := w.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}

	return n, err
}
