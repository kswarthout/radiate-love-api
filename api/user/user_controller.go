package userapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kswarthout/radiate-love-api/api"
)

// Controller handles user API endpoints
type Controller struct {
	api.BaseController
	Repo *UsersService
}

// AddRoutes registers user endpoints and handlers
func (c *Controller) AddRoutes(router *mux.Router) *mux.Router {
	fmt.Println(&c.Repo.users)
	basePath := "/api"
	router.HandleFunc(basePath+"/login", c.login).Methods("POST")
	router.HandleFunc(basePath+"/register", c.register).Methods("POST")
	router.HandleFunc(basePath+"/users", c.createUser).Methods("POST")
	router.HandleFunc(basePath+"/users", c.listUsers).Methods("GET")
	router.HandleFunc(basePath+"/users/{id}", c.getUser).Methods("GET")
	router.HandleFunc(basePath+"/users/{id}", c.updateUser).Methods("PUT")
	router.HandleFunc(basePath+"/users/{id}", c.deleteUser).Methods("DELETE")
	return router
}

func (c *Controller) createUser(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *Controller) listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listUsers hit")
	c.SendJSON(w, r, c.Repo.users, http.StatusOK)
}

func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *Controller) updateUser(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *Controller) deleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login hit")
	credentials := Credentials{}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		c.SendJSON(
			w,
			r,
			api.Error{Message: `Bad request. Email and/or password missing or invalid.`, Error: err.Error()},
			http.StatusBadRequest,
		)
	} else {
		fmt.Println(credentials)
		c.SendJSON(w, r, credentials, http.StatusOK)
	}
}

func (c *Controller) register(w http.ResponseWriter, r *http.Request) {
	//TODO
}
