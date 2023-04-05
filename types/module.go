package types

import (
	"errors"

	"estimator/utils"
)

// ModuleTypes defines the available types for a module.
var ModuleTypes []string = []string{
	"short-text",
	"multiple-choice",
	"heading",
	"full-name",
}

// Module defines the module interface.
type Module interface {
	SetID(id string)
	GetType() string
	Validate() error
}

// ValidateType handles validating the module type.
func ValidateType(t string) error {
	if !utils.SliceContains(ModuleTypes, t) {
		return errors.New("invalid module type")
	}

	return nil
}
