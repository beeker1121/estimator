package projects

import (
	"encoding/json"
	"net/http"

	apictx "estimator/cmd/api/context"
	"estimator/cmd/api/errors"
	"estimator/cmd/api/middleware/auth"
	"estimator/cmd/api/response"
	serverrors "estimator/services/errors"
	"estimator/services/projects"
	"estimator/types"

	"github.com/beeker1121/httprouter"
)

// Project defines the project request/response.
type Project struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
}

// ResultCreate defines the response data for the HandleCreate handler.
type ResultCreate struct {
	Data Project `json:"data"`
}

// ResultGet defines the response data for the HandleGet handler.
type ResultGet struct {
	Data Project `json:"data"`
}

// ResultUpdate defines the response data for the HandleUpdate handler.
type ResultUpdate struct {
	Data Project `json:"data"`
}

// New creates a new project handler.
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes.
	router.POST("/api/v1/project", HandleCreate(ac))
	router.GET("/api/v1/project/:id", HandleGet(ac))
	router.POST("/api/v1/project/:id", HandleUpdate(ac))
}

// HandleCreate is the HTTP handler function for creating a project.
func HandleCreate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body.
		var p Project
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Create a new services project.
		sp, err := ac.Services.Projects.Create(&types.Project{
			AccountID: p.AccountID,
			Name:      p.Name,
		})
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("projects.Create() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result.
		result := ResultCreate{
			Data: Project{
				ID:        sp.ID,
				AccountID: sp.AccountID,
				Name:      sp.Name,
			},
		}

		// Respond with JSON.
		if err := response.JSON(w, true, result); err != nil {
			ac.Logger.Printf("response.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}

// HandleGet is the HTTP handler function for getting a project.
func HandleGet(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the project ID.
		id := httprouter.GetParam(r, "id")

		// Get this user from the request context.
		user, err := auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Get the project.
		sp, err := ac.Services.Projects.GetByIDAndUserID(id, user.ID)
		if err == projects.ErrProjectNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("projects.GetByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
		}

		// Create a new result.
		result := ResultGet{
			Data: Project{
				ID:        sp.ID,
				AccountID: sp.AccountID,
				Name:      sp.Name,
			},
		}

		// Respond with JSON.
		if err := response.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}

// HandleUpdate is the HTTP handler function for updating a project.
func HandleUpdate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the project ID.
		id := httprouter.GetParam(r, "id")

		// Parse the request body.
		var p types.ProjectUpdateParams
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Get this user from the request context.
		user, err := auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Update the project.
		sa, err := ac.Services.Projects.UpdateByIDAndUserID(id, user.ID, &types.ProjectUpdateParams{
			AccountID: p.AccountID,
			Name:      p.Name,
		})
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("projects.UpdateByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result.
		result := ResultUpdate{
			Data: Project{
				ID:   sa.ID,
				Name: sa.Name,
			},
		}

		// Respond with JSON.
		if err := response.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}