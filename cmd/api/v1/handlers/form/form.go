package form

import (
	"fmt"
	"net/http"

	apictx "estimator/cmd/api/context"
	"estimator/cmd/api/response"

	"github.com/beeker1121/httprouter"
)

// Form defines the form response.
type Form struct {
	ID string `json:"id"`
}

// New creates a new form handler.
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes.
	router.GET("/api/v1/form/:id", HandleGetForm(ac))
}

// HandleGetForm is the HTTP handler function for getting a form.
func HandleGetForm(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the form ID.
		id := httprouter.GetParam(r, "id")

		// Get the form.
		servf, err := ac.Services.Form.GetByID(id)
		if err != nil {
			// TODO: Create response package to handle sending back JSON and
			//       errors.
			w.Write([]byte("error getting form"))
		}

		// Map to API form response.
		f := &Form{
			ID: servf.ID,
		}

		// Respond with JSON.
		if err := response.JSON(w, true, f); err != nil {
			// TODO: Use logger.
			fmt.Printf("error in handler: %v\n", err)
		}
	}
}
