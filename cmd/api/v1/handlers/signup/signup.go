package signup

import (
	"encoding/json"
	"net/http"

	apictx "estimator/cmd/api/context"
	"estimator/cmd/api/errors"
	"estimator/cmd/api/middleware/auth"
	"estimator/cmd/api/response"
	serverrors "estimator/services/errors"
	"estimator/types"

	"github.com/beeker1121/httprouter"
)

// ResultPost defines the response data for the HandlePost handler.
type ResultPost struct {
	Data string `json:"data"`
}

// New creates the routes for the signup endpoints of the API.
func New(ac *apictx.Context, router *httprouter.Router) {
	// Handle the routes.
	router.POST("/api/v1/signup", HandlePost(ac))
}

// HandlePost handles the /api/v1/signup POST route of the API.
func HandlePost(ac *apictx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the parameters from the request body.
		user := &types.User{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			errors.Default(ac.Logger, w, errors.ErrBadRequest)
			return
		}

		// Create the user.
		user, err := ac.Services.Users.Create(user)
		if pes, ok := err.(*serverrors.ParamErrors); ok && err != nil {
			errors.Params(ac.Logger, w, http.StatusBadRequest, pes)
			return
		} else if err != nil {
			ac.Logger.Printf("users.Create() service error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Issue a new JWT for this user.
		token, err := auth.NewJWT(ac, user.Password, user.ID)
		if err != nil {
			ac.Logger.Printf("auth.NewJWT() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}

		// Create a new Result.
		result := ResultPost{
			Data: token,
		}

		// Respond with JSON.
		if err := response.JSON(w, true, result); err != nil {
			ac.Logger.Printf("render.JSON() error: %s\n", err)
			errors.Default(ac.Logger, w, errors.ErrInternalServerError)
			return
		}
	}
}
