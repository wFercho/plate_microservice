package api

import (
	"net/http"
	"plate_microservice/db"

	kclk "plate_microservice/keycloack"
)

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIServer struct {
	listenAddr string
	store      db.Storage
	kclk       kclk.Keycloak
}

type PageData struct {
	Plate   string
	Allowed bool
}
