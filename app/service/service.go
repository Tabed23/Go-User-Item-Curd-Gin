package service

import (
	"context"
	"fmt"
	"gin-crud/app/database"
	"gin-crud/app/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatUser(usr *types.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	u := types.User{
		ID:        primitive.NewObjectID(),
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Product:   usr.Product,
	}
	fmt.Println(u.Product)
	_, err := database.DB.Database("Users").Collection("users").InsertOne(ctx, &u)
	if err != nil {
		return "", err
	}
	return "ok", nil
}

// func FindById(string) (*types.User, error) {

// }

// func DeleteUser(string) (string, error) {

// }

// func UpdateUser(string) (*types.User, error) {

// }

func GetUsers() (*[]types.User, error) {
	usr := []types.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := database.DB.Database("Users").Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for data.Next(ctx) {
		var u types.User
		if err := data.Decode(&u); err != nil {
			return nil, err
		}
		usr = append(usr, u)
	}
	return &usr, nil

}
