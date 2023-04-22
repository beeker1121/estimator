package context

import (
	"log"

	"estimator/cmd/api/config"
	"estimator/services"
)

// Context defines the API context, which will contain the configuration,
// services, logger, and anything else needed across HTTP handlers.
type Context struct {
	Config   *config.Config
	Logger   *log.Logger
	Services *services.Services
}

// New creates a new API context.
func New(config *config.Config, logger *log.Logger, services *services.Services) *Context {
	return &Context{
		Config:   config,
		Logger:   logger,
		Services: services,
	}
}
