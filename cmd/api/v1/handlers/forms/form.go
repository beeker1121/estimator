package forms

import (
	"encoding/json"
	"net/http"

	apictx "estimator/cmd/api/context"
	"estimator/cmd/api/errors"
	"estimator/cmd/api/response"
	serverrors "estimator/services/errors"
	"estimator/services/forms"
	"estimator/types"

	"github.com/beeker1121/httprouter"
)

// Form defines the form request/response.
type Form struct {
	ID         string        `json:"id"`
	ProjectID  string        `json:"project_id"`
	Name       string        `json:"name"`
	Properties Properties    `json:"properties"`
	Button     Button        `json:"button"`
	Modules    []interface{} `json:"modules"`
}

// Properties defines the form properties.
type Properties struct {
	BackgroundColor string `json:"background_color"`
	FontColor       string `json:"font_color"`
}

// Button defines the form button.
type Button struct {
	BackgroundColor string `json:"background_color"`
	Color           string `json:"color"`
	FontSize        string `json:"font_size"`
	FontFamily      string `json:"font_family"`
}

// ResultCreate defines the response data for the HandleCreate handler.
type ResultCreate struct {
	Data Form `json:"data"`
}

// ResultGet defines the response data for the HandleGet handler.
type ResultGet struct {
	Data Form `json:"data"`
}

// ResultUpdate defines the response data for the HandleUpdate handler.
type ResultUpdate struct {
	Data Form `json:"data"`
}

// New creates a new forms handler.
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes.
	router.POST("/api/v1/forms", HandleCreate(ac))
	router.GET("/api/v1/forms/:id", HandleGet(ac))
	router.POST("/api/v1/forms/:id", HandleUpdate(ac))
}

// HandleCreate is the HTTP handler function for creating a form.
func HandleCreate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body.
		var f Form
		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Map modules interface to module types.
		modules, err := ac.Services.Forms.InterfaceToModules(f.Modules)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("form.InterfaceToModules() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new services form.
		sf, err := ac.Services.Forms.Create(&types.Form{
			ProjectID: f.ProjectID,
			Name:      f.Name,
			Properties: types.FormProperties{
				BackgroundColor: f.Properties.BackgroundColor,
				FontColor:       f.Properties.FontColor,
			},
			Button: types.FormButton{
				BackgroundColor: f.Button.BackgroundColor,
				Color:           f.Button.Color,
				FontSize:        f.Button.FontSize,
				FontFamily:      f.Button.FontFamily,
			},
			Modules: modules,
		})
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("forms.Create() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Set form response.
		fres := Form{
			ID:        sf.ID,
			ProjectID: sf.ProjectID,
			Name:      sf.Name,
			Properties: Properties{
				BackgroundColor: sf.Properties.BackgroundColor,
				FontColor:       sf.Properties.FontColor,
			},
			Button: Button{
				BackgroundColor: sf.Button.BackgroundColor,
				Color:           sf.Button.Color,
				FontSize:        sf.Button.FontSize,
				FontFamily:      sf.Button.FontFamily,
			},
		}

		// Get JSON for modules.
		modulesJSON, err := json.Marshal(sf.Modules)
		if err != nil {
			ac.Logger.Printf("json.Marshal() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
		if err := json.Unmarshal(modulesJSON, &fres.Modules); err != nil {
			ac.Logger.Printf("json.Unmarshal() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result.
		result := ResultCreate{
			Data: fres,
		}

		// Respond with JSON.
		if err := response.JSON(w, true, result); err != nil {
			ac.Logger.Printf("response.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}

// HandleGet is the HTTP handler function for getting a form.
func HandleGet(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the form ID.
		id := httprouter.GetParam(r, "id")

		// Get the form.
		sf, err := ac.Services.Forms.GetByID(id)
		if err == forms.ErrFormNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("forms.GetByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
		}

		// Map to API form.
		//
		// TODO: See if we should match the HandleCreate and HandleUpdate
		//       way of setting the Modules field.
		f := &Form{
			ID:        sf.ID,
			ProjectID: sf.ProjectID,
			Name:      sf.Name,
			Properties: Properties{
				BackgroundColor: sf.Properties.BackgroundColor,
				FontColor:       sf.Properties.FontColor,
			},
			Button: Button{
				BackgroundColor: sf.Button.BackgroundColor,
				Color:           sf.Button.Color,
				FontSize:        sf.Button.FontSize,
				FontFamily:      sf.Button.FontFamily,
			},
			Modules: []interface{}{},
		}
		for _, v := range sf.Modules {
			f.Modules = append(f.Modules, v)
		}

		// Create a new result.
		result := ResultGet{
			Data: *f,
		}

		// Respond with JSON.
		if err := response.JSON(w, true, result); err != nil {
			ac.Logger.Printf("response.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}

// HandleUpdate is the HTTP handler function for updating a form.
func HandleUpdate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body.
		var f Form
		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Get the form ID.
		id := httprouter.GetParam(r, "id")

		// TODO: Get this user from the request context.
		// TODO: Get the account and role.

		// Map modules interface to module types.
		modules, err := ac.Services.Forms.InterfaceToModules(f.Modules)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("form.InterfaceToModules() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Update the form.
		sf, err := ac.Services.Forms.UpdateByIDAndUserID(id, "", &types.Form{
			ProjectID: f.ProjectID,
			Name:      f.Name,
			Properties: types.FormProperties{
				BackgroundColor: f.Properties.BackgroundColor,
				FontColor:       f.Properties.FontColor,
			},
			Button: types.FormButton{
				BackgroundColor: f.Button.BackgroundColor,
				Color:           f.Button.Color,
				FontSize:        f.Button.FontSize,
				FontFamily:      f.Button.FontFamily,
			},
			Modules: modules,
		})
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err == forms.ErrFormNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("forms.UpdateByIDAndUserID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new response form.
		res := Form{
			ID:        sf.ID,
			ProjectID: sf.ProjectID,
			Name:      sf.Name,
			Properties: Properties{
				BackgroundColor: sf.Properties.BackgroundColor,
				FontColor:       sf.Properties.FontColor,
			},
			Button: Button{
				BackgroundColor: sf.Button.BackgroundColor,
				Color:           sf.Button.Color,
				FontSize:        sf.Button.FontSize,
				FontFamily:      sf.Button.FontFamily,
			},
		}

		// Get JSON for modules.
		modulesJSON, err := json.Marshal(sf.Modules)
		if err != nil {
			ac.Logger.Printf("json.Marshal() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
		if err := json.Unmarshal(modulesJSON, &res.Modules); err != nil {
			ac.Logger.Printf("json.Unmarshal() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Respond with JSON.
		if err := response.JSON(w, true, res); err != nil {
			ac.Logger.Printf("response.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}
