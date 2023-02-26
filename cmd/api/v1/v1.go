package v1

import (
	apictx "estimator/cmd/api/context"
	"estimator/cmd/api/v1/handlers/form"

	"github.com/beeker1121/httprouter"
)

// New creates a new v1 API.
func New(ac *apictx.Context, r *httprouter.Router) {
	form.New(ac, r)
}
