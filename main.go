package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	userapi "github.com/kswarthout/radiate-love-api/api/user"
	"github.com/kswarthout/radiate-love-api/domain"
	"github.com/kswarthout/radiate-love-api/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

var dal *service.DAL
var users *userapi.UsersService
var config *domain.Config

// Routes configures all API routes
func Routes() *mux.Router {

	// INIT CONTROLLERS
	uc := userapi.Controller{
		Repo: users,
	}

	// REGISTER ROUTES
	router := mux.NewRouter().StrictSlash(true)
	router = uc.AddRoutes(router)

	return router
}

func initConfig() {
	config = &domain.Config{}
	configService := service.ConfigService{
		Config:   config,
		Location: getConfigPath(),
	}
	configService.Reload()
	go configService.Watch(time.Second * 30)
}

func initDataAccess() {

	// INIT DB CONNECTION
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri, err := getMongoURI()
	if err != nil {
		log.Fatal(err)
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// CHECK CONNECTION
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// INIT DATA ACCESS LAYER
	dal = &service.DAL{
		Client: client,
		Config: config,
	}
	dal.Connect()

	users = &userapi.UsersService{
		DAL: dal,
	}
	users.Load()
}

func getConfigPath() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"config.", env, ".yaml"}
	return strings.Join(filename, "")
}

func getAPIPort() (string, error) {
	var port = os.Getenv("PORT")
	if port == "" {
		apiConfig, err := config.Get("base")
		if err != nil {
			return "8080", err
		}
		if configPort, ok := apiConfig["port"].(string); ok {
			port = configPort
		} else {
			port = "8080"
		}

		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port, nil
}

func getMongoURI() (string, error) {
	var uri = os.Getenv("DATABASE_URL")
	if uri == "" {
		var port = ""
		var host = ""

		dbConfig, err := config.Get("db")
		if err != nil {
			return "", err
		}

		if configPort, ok := dbConfig["port"].(string); ok {
			port = configPort
		} else {
			port = "27017"
		}

		if configHost, ok := dbConfig["host"].(string); ok {
			host = configHost
		} else {
			host = "localhost"
		}

		uri = "mongodb://" + host + ":" + port
		fmt.Println("INFO: No DATABASE_URL environment variable detected, defaulting to " + uri)
	}

	return uri, nil
}

func main() {
	fmt.Println("Starting the application...")

	// LOAD APP CONFIG & INIT SERVICES
	initConfig()
	initDataAccess()

	// INIT ROUTER
	router := Routes()
	apiPort, err := getAPIPort()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Serving application at Port :" + apiPort)
	log.Fatal(http.ListenAndServe(apiPort, router))

}
