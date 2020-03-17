package authbundle

import (
	"net/http"

	"github.com/kswarthout/radiate-love-api/app/bundles/auth"
)

// UsersController handle users endpoints
type AuthController struct {
	auth.Controller
}

func (c *AuthController) index(w http.ResponseWriter, r *http.Request) {

}
