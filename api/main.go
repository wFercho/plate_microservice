package api

import (
	"log"
	"net/http"
	"plate_microservice/db"
	auth "plate_microservice/middleware"
)

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddr string, store db.Storage, authMiddleware auth.AuthMiddleware) *APIServer {
	return &APIServer{
		listenAddr:     listenAddr,
		store:          store,
		authMiddleware: authMiddleware,
	}
}

func (s *APIServer) Run() {

	sMux := http.NewServeMux()

	sMux.HandleFunc("GET /plate/{plate}", s.authMiddleware.Authenticate(makeHTTPHandleFunc(s.handleIsPlateAvailable)))

	log.Println("Plates API running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, sMux)
}
