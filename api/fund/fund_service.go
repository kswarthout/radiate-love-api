package fundapi

import (
	"github.com/kswarthout/radiate-love-api/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FundsService is the repositorty for API access to DB funds
type FundsService struct {
	DAL             *service.DAL
	fundsCollection *mongo.Collection
	funds           []bson.M
}

// Load retrieves all documents from the funds collection
func (u *FundsService) Load() {
	// TODO
}
