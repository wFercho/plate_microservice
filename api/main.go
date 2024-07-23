package api

import (
	"log"
	"net/http"
	"plate_microservice/db"
	kclk "plate_microservice/keycloack"
)

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddr string, store db.Storage, kclk kclk.Keycloak) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
		kclk:       kclk,
	}
}

func (s *APIServer) Run() {

	sMux := http.NewServeMux()

	//sMux.HandleFunc("GET /", s.kclk.KeycloakMiddleware(makeHTTPHandleFunc(s.homePage)))
	sMux.HandleFunc("GET /", s.kclk.SessionMiddleware(makeHTTPHandleFunc(s.homePage)))
	sMux.HandleFunc("GET /login", s.kclk.SessionMiddleware(makeHTTPHandleFunc(s.loginHandler)))
	sMux.HandleFunc("POST /login", makeHTTPHandleFunc(s.loginHandler))
	sMux.HandleFunc("GET /resultado", s.kclk.SessionMiddleware(makeHTTPHandleFunc(s.resultadoPage)))
	sMux.HandleFunc("GET /vehicle", s.kclk.SessionMiddleware(s.kclk.KeycloakMiddleware(makeHTTPHandleFunc(s.handleIsPlateAvailable))))
	sMux.HandleFunc("POST /vehicle", s.kclk.SessionMiddleware(s.kclk.KeycloakMiddleware(makeHTTPHandleFunc(s.handleIsPlateAvailable))))

	sMux.HandleFunc("POST /logout", s.kclk.SessionMiddleware(makeHTTPHandleFunc(s.logoutHandler)))

	log.Println("Plates API running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, sMux)
}
