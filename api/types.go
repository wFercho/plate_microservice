package api

import (
	"net/http"
	"plate_microservice/db"

	auth "plate_microservice/middleware"
)

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIServer struct {
	listenAddr     string
	store          db.Storage
	authMiddleware auth.AuthMiddleware
}
