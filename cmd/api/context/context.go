package context

import "estimator/services"

// Context defines the API context, which will contain the configuration,
// services, logger, and anything else needed across HTTP handlers.
type Context struct {
	Services *services.Services
}

// New creates a new API context.
func New(s *services.Services) *Context {
	return &Context{
		Services: s,
	}
}
