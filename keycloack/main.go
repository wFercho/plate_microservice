package keycloack

import (
	"log"
	"os"

	"github.com/Nerzal/gocloak/v13"
	"github.com/joho/godotenv"
)

type Keycloak struct {
	Gocloak      *gocloak.GoCloak // keycloak client
	ClientId     string           // clientId specified in Keycloak
	ClientSecret string           // client secret specified in Keycloak
	Realm        string           // realm specified in Keycloak
}

func NewKeycloak() *Keycloak {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	keycloak_url := os.Getenv("KEYCLOAK_URL")
	client_id := os.Getenv("KEYCLOAK_CLIENT_ID")
	realm := os.Getenv("KEYCLOAK_REALM")
	secret := os.Getenv("KEYCLOAK_CLIENT_SECRET")

	// client := gocloak.NewClient(keycloak_url)
	// restyClient := client.RestyClient()
	// restyClient.SetDebug(true)
	// restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	return &Keycloak{
		//Gocloak:      *client,
		Gocloak:      gocloak.NewClient(keycloak_url),
		ClientId:     client_id,
		ClientSecret: secret,
		Realm:        realm,
	}
}
