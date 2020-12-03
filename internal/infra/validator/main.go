/* Package validator provide validation service for objects (object package) */
package validator

import (
	"daanretard/internal/infra/errors"
	govalidator "github.com/go-playground/validator/v10"
)

// New create a Validator
func New() *Validator {
	v := new(Validator)
	v.v = govalidator.New()
	return v
}

// Validator wrapper of govalidator.Validate so other packages don't need to
// know anything about it
type Validator struct {
	v *govalidator.Validate
}

// Validate validate struct and return errors.Errors, i should be the pointer
// of the instance
func (v *Validator) Validate(i interface{}) error {
	if e := v.v.Struct(i); e != nil {
		var errs errors.Errors
		for _, err := range e.(govalidator.ValidationErrors) {
			errs = append(errs, errors.New(err.Tag(), err.Field()))
		}
		return errs
	}
	return nil
}
