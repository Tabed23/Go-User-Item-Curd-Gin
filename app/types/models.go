package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string `bson:first_name"`
	LastName  string `bson:last_name"`
	Product   Item `bson:"items"`
}

type Item struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string `bson:"title"`
	Price float32 `bson:"price"`
}
