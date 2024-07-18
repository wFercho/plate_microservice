package db

import e "plate_microservice/entities"

type Storage interface {
	PlateStorage
}

type PlateStorage interface {
	//GetPlateByID(int) (*e.Plate, error)
	GetCarByPlate(string) (*e.Car, error)
}
