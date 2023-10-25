package models

type SrvBookModel struct {
	Id    string `json:"id" bson:"id"`
	Title string `json:"title" bson:"title"`
	Price int    `json:"price" bson:"price"`
}

type UpdateSrvBookModel struct {
	Title string `json:"title" bson:"title,omitempty"`
	Price int    `json:"price" bson:"price,omitempty"`
}
