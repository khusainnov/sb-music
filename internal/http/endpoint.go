package http

import "net/http"

func (s *Server) initEndpoints() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HELLO SWAGGER"))
		return
	})

	return mux
}
