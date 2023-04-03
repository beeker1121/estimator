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
type Form struct {
	ID      string        `json:"id"`
	Modules []interface{} `json:"modules"`
}

// New creates a new form handler.
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes.
	router.POST("/api/v1/form", HandleCreate(ac))
	router.GET("/api/v1/form/:id", HandleGet(ac))
	router.POST("/api/v1/form/:id", HandleUpdate(ac))
}

// HandleCreate is the HTTP handler function for creating a form.
func HandleCreate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body.
		var f Form
		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
			w.Write([]byte("error decoding request body"))
			return
		}

		// Map modules interface to module types.
		modules, err := ac.Services.Form.InterfaceToModules(f.Modules)
		if err != nil {
			w.Write([]byte("error converting interface to modules"))
			return
		}

		// Create a new services form.
		sf, err := ac.Services.Form.Create(&types.Form{
			Modules: modules,
		})
		if err != nil {
			w.Write([]byte("error creating form"))
			return
		}

		// Create a new response form.
		res := Form{
			ID: sf.ID,
		}

		// Get JSON for modules.
		modulesJSON, err := json.Marshal(sf.Modules)
		if err != nil {
			w.Write([]byte("error marshaling modules to JSON"))
			return
		}
		if err := json.Unmarshal(modulesJSON, &res.Modules); err != nil {
			w.Write([]byte("error unmarshaling modules to interface"))
			return
		}

		// Respond with JSON.
		if err := response.JSON(w, true, res); err != nil {
			// TODO: Use logger.
			fmt.Printf("error in handler: %v\n", err)
		}
	}
}

// HandleGet is the HTTP handler function for getting a form.
func HandleGet(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the form ID.
		id := httprouter.GetParam(r, "id")

		// Get the form.
		sf, err := ac.Services.Form.GetByID(id)
		// TODO: Implement else if for ErrFormNotFound.
		if err != nil {
			// TODO: Create response package to handle sending back JSON and
			//       errors.
			w.Write([]byte("error getting form"))
			return
		}

		// Map to API form response.
		f := &Form{
			ID:      sf.ID,
			Modules: []interface{}{},
		}
		for _, v := range sf.Modules {
			f.Modules = append(f.Modules, v)
		}

		// Respond with JSON.
		if err := response.JSON(w, true, f); err != nil {
			// TODO: Use logger.
			fmt.Printf("error in handler: %v\n", err)
		}
	}
}

// HandleUpdate is the HTTP handler function for creating a form.
func HandleUpdate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body.
		var f Form
		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
			w.Write([]byte("error decoding request body"))
			return
		}

		// Get the form ID.
		id := httprouter.GetParam(r, "id")

		// TODO: Get this member from the request context.

		// Map modules interface to module types.
		modules, err := ac.Services.Form.InterfaceToModules(f.Modules)
		if err != nil {
			w.Write([]byte("error converting interface to modules"))
			return
		}

		// Update the form.
		sf, err := ac.Services.Form.UpdateByIDAndMemberID(id, "", &types.Form{
			Modules: modules,
		})
		// TODO: Implement param error type check first.
		// TODO: Implement else if for ErrFormNotFound.
		if err != nil {
			// TODO: Create response package to handle sending back JSON and
			//       errors.
			w.Write([]byte("error getting form"))
			return
		}

		// Create a new response form.
		res := Form{
			ID: sf.ID,
		}

		// Get JSON for modules.
		modulesJSON, err := json.Marshal(sf.Modules)
		if err != nil {
			w.Write([]byte("error marshaling modules to JSON"))
			return
		}
		if err := json.Unmarshal(modulesJSON, &res.Modules); err != nil {
			w.Write([]byte("error unmarshaling modules to interface"))
			return
		}

		// Respond with JSON.
		if err := response.JSON(w, true, res); err != nil {
			// TODO: Use logger.
			fmt.Printf("error in handler: %v\n", err)
		}
	}
}
