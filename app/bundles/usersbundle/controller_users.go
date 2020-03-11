package usersbundle

import (
	"net/http"

	"github.com/kswarthout/radiate-love-api/app/bundles/common"
)

// UsersController handle users endpoints
type UsersController struct {
	common.Controller
}

func (c *UsersController) index(w http.ResponseWriter, r *http.Request) {

}
