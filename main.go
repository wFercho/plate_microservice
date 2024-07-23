package main

import (
	"log"
	"plate_microservice/api"
	db "plate_microservice/db/mongodb"
	kclk "plate_microservice/keycloack"
)

func main() {
	kclk := kclk.NewKeycloak()
	if kclk == nil {
		log.Fatal("No está creando el objeto de Keycloak")
	}

	store, err := db.NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}

	// if err := store.Init(); err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Printf("%+v\n", store)
	server := api.NewAPIServer(":3000", store, *kclk)
	server.Run()
}
