package fundapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kswarthout/radiate-love-api/api"
)

// Controller handles user API endpoints
type Controller struct {
	api.BaseController
	Repo *FundsService
}

// AddRoutes registers user endpoints and handlers
func (c *Controller) AddRoutes(router *mux.Router) *mux.Router {
	fmt.Println(&c.Repo.users)
	basePath := "/api/funds"
	router.HandleFunc(basePath+"/", c.createRecord).Methods("POST")
	router.HandleFunc(basePath+"/", c.listRecords).Methods("GET")
	router.HandleFunc(basePath+"/{id}", c.getRecord).Methods("GET")
	router.HandleFunc(basePath+"/{id}", c.updateRecord).Methods("PUT")
	router.HandleFunc(basePath+"/{id}", c.deleteRecord).Methods("DELETE")
	return router
}

func (c *Controller) createRecord(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *Controller) listRecords(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *Controller) getRecord(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *Controller) updateRecord(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *Controller) deleteRecord(w http.ResponseWriter, r *http.Request) {
	// TODO
}
