package userapi

import (
	"fmt"

	"github.com/kswarthout/radiate-love-api/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UsersService is the repositorty for API access to DB users
type UsersService struct {
	DAL             *service.DAL
	usersCollection *mongo.Collection
	users           []User
}

// Load retrieves all documents from the user collection
func (u *UsersService) Load() {
	var users []User
	u.usersCollection = u.DAL.DB.Collection("users")
	usersCursor, err := u.usersCollection.Find(u.DAL.Ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	if err = usersCursor.All(u.DAL.Ctx, &users); err != nil {
		panic(err)
	}
	u.users = users
	fmt.Println(u.users)
}
