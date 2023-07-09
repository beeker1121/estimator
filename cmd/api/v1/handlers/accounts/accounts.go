package form

import (
	"encoding/json"
	"net/http"

	apictx "estimator/cmd/api/context"
	"estimator/cmd/api/errors"
	"estimator/cmd/api/middleware/auth"
	"estimator/cmd/api/response"
	"estimator/services/accounts"
	serverrors "estimator/services/errors"
	"estimator/types"

	"github.com/beeker1121/httprouter"
)

// Account defines the account request/response.
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ResultCreate defines the response data for the HandleCreate handler.
type ResultCreate struct {
	Data Account `json:"data"`
}

// ResultGet defines the response data for the HandleGet handler.
type ResultGet struct {
	Data Account `json:"data"`
}

// ResultUpdate defines the response data for the HandleUpdate handler.
type ResultUpdate struct {
	Data Account `json:"data"`
}

// New creates a new account handler.
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes.
	router.POST("/api/v1/account", HandleCreate(ac))
	router.GET("/api/v1/account/:id", HandleGet(ac))
	router.POST("/api/v1/account/:id", HandleUpdate(ac))
}

// HandleCreate is the HTTP handler function for creating a form.
func HandleCreate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body.
		var a Account
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Create a new services account.
		sa, err := ac.Services.Accounts.Create(&types.Account{
			Name: a.Name,
		})
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("accounts.Create() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Set account response.
		ares := Account{
			ID:   sa.ID,
			Name: sa.Name,
		}

		// Create a new Result.
		result := ResultCreate{
			Data: ares,
		}

		// Respond with JSON.
		if err := response.JSON(w, true, result); err != nil {
			ac.Logger.Printf("response.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}

// HandleGet is the HTTP handler function for getting a account.
func HandleGet(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the account ID.
		id := httprouter.GetParam(r, "id")

		// Get this user from the request context.
		user, err := auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Get the account.
		sa, err := ac.Services.Accounts.GetByIDAndUserID(id, user.ID)
		if err == accounts.ErrAccountNotFound {
			errors.Default(ac.Logger, w, errors.New(http.StatusNotFound, "", err.Error()))
			return
		} else if err != nil {
			ac.Logger.Printf("accounts.GetByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
		}

		// Create a new result.
		result := ResultGet{
			Data: Account{
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

// HandleUpdate is the HTTP handler function for updating an account.
func HandleUpdate(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the account ID.
		id := httprouter.GetParam(r, "id")

		// Parse the request body.
		var a types.AccountUpdateParams
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Get this user from the request context.
		user, err := auth.GetUserFromRequest(r)
		if err != nil {
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Update the account.
		sa, err := ac.Services.Accounts.UpdateByIDAndUserID(id, user.ID, &types.AccountUpdateParams{
			Name: a.Name,
		})
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("accounts.UpdateByID() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new result.
		result := ResultUpdate{
			Data: Account{
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
