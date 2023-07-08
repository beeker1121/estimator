package v1

import (
	apictx "estimator/cmd/api/context"
	"estimator/cmd/api/v1/handlers/forms"
	"estimator/cmd/api/v1/handlers/login"
	"estimator/cmd/api/v1/handlers/signup"

	"github.com/beeker1121/httprouter"
)

// New creates a new v1 API.
func New(ac *apictx.Context, r *httprouter.Router) {
	forms.New(ac, r)
	login.New(ac, r)
	signup.New(ac, r)
}
