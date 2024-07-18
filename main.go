package main

import (
	"context"
	"log"
	"os"
	"plate_microservice/api"
	db "plate_microservice/db/mongodb"
	"plate_microservice/middleware"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/joho/godotenv"
)

var authMiddleware *middleware.AuthMiddleware

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	url_realm := os.Getenv("KEYCLOAK_URL_REALM")
	client_id := os.Getenv("KEYCLOAK_CLIENT_ID")
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, url_realm)
	if err != nil {
		panic(err)
	}

	authMiddleware = middleware.NewAuthMiddleware(provider, client_id)
}

func main() {
	store, err := db.NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}

	// if err := store.Init(); err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Printf("%+v\n", store)
	server := api.NewAPIServer(":3000", store, *authMiddleware)
	server.Run()
}
