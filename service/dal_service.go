package service

import (
	"context"
	"log"
	"time"

	"github.com/kswarthout/radiate-love-api/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// DAL stores shared data access resources
type DAL struct {
	Client *mongo.Client
	Config *domain.Config
	DB     *mongo.Database
	Ctx    context.Context
}

// Connect inits MongoDB connection
func (d *DAL) Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	d.Ctx = ctx
	dbConfig, err := d.Config.Get("db")
	if err != nil {
		log.Fatal(err)
	}
	d.DB = d.Client.Database(dbConfig["name"].(string))
}
