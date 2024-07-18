package entities

type Car struct {
	//Id           string `json:"id" bson:"_id"`
	Id           string `json:"id" bson:"_id"`
	Plate_number string `json:"plateNumber" `
	IsAvailable  bool   `json:"isAvailable"`
	Brand        string `json:"brand"`
	Model        uint   `json:"model"`
	Color        string `json:"color"`
}
