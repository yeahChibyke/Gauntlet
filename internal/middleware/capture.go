package middleware

import "net/http"

// ResponseRecorder captures information about an HTTP response
// while transparently forwarding writes to the real ResponseWriter.
type ResponseRecorder struct {
	http.ResponseWriter

	StatusCode int
	Bytes      int
}

func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

func (r *ResponseRecorder) WriteHeader(code int) {
	r.StatusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *ResponseRecorder) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.Bytes += n
	return n, err
}
