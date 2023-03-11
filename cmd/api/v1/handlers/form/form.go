package form

import (
	"encoding/json"
	"fmt"
	"net/http"

	apictx "estimator/cmd/api/context"
	"estimator/cmd/api/response"
	"estimator/types"

	"github.com/beeker1121/httprouter"
)

// Form defines the form request/response.
type Form types.Form

// New creates a new form handler.
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes.
	router.POST("/api/v1/form", HandleCreateForm(ac))
	router.GET("/api/v1/form/:id", HandleGetForm(ac))
}

// HandleCreateForm is the HTTP handler function for creating a form.
func HandleCreateForm(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body.
		var f Form
		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
			w.Write([]byte("error decoding request body"))
			return
		}

		// Create a new form.
		sf, err := ac.Services.Form.Create(&types.Form{
			Modules: f.Modules,
		})
		if err != nil {
			w.Write([]byte("error creating form"))
			return
		}

		// Create a new form.
		res := Form{
			ID:      sf.ID,
			Modules: sf.Modules,
		}

		// Respond with JSON.
		if err := response.JSON(w, true, res); err != nil {
			// TODO: Use logger.
			fmt.Printf("error in handler: %v\n", err)
		}
	}
}

// HandleGetForm is the HTTP handler function for getting a form.
func HandleGetForm(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the form ID.
		id := httprouter.GetParam(r, "id")

		// Get the form.
		sf, err := ac.Services.Form.GetByID(id)
		if err != nil {
			// TODO: Create response package to handle sending back JSON and
			//       errors.
			w.Write([]byte("error getting form"))
			return
		}

		// Map to API form response.
		f := &Form{
			ID: sf.ID,
		}

		// Respond with JSON.
		if err := response.JSON(w, true, f); err != nil {
			// TODO: Use logger.
			fmt.Printf("error in handler: %v\n", err)
		}
	}
}
