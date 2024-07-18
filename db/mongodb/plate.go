package mongodb

import (
	"context"
	"log"
	e "plate_microservice/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterRequest struct {
	ID primitive.ObjectID `bson:"_id"`
	//userId primitive.ObjectID
	Plate string `bson:"plate"`
	Brand string `bson:"brand"`
	Model uint   `bson:"model"`
	Color string `bson:"color"`
	//technomechanicsId string
	//soatId string
	//ownershipCarId string
	//idCardId string
	//reviewed string
	//comments string
	Approved bool `bson:"approved"`
	//requestDate time.Time

}

func (store *MongoStore) GetCarByPlate(plate string) (*e.Car, error) {
	coll := store.cli.Database(DATA_BASE).Collection(COLLECTION)

	filter := bson.M{"plate": plate}
	//var result RegisterRequest
	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	var carBson RegisterRequest
	bsonBytes, err := bson.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	err = bson.Unmarshal(bsonBytes, &carBson)
	if err != nil {
		log.Fatal(err)
	}
	car := e.Car{
		Id:           carBson.ID.Hex(),
		Plate_number: carBson.Plate,
		IsAvailable:  carBson.Approved,
		Brand:        carBson.Brand,
		Model:        carBson.Model,
		Color:        carBson.Color,
	}

	return &car, nil
}
