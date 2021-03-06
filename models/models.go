package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Book struct {
	//left in Go right in mongo
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn   string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}

type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User_id  int32              `json:"user_id" bson:"user_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
	Role     string             `json:"role" bson:"role,omitempty"`
	Roledesc string             `json:"roledesc" bson:"roledesc,omitempty"`
}
