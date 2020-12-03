package fbgraph

import (
	"daanretard/internal/infra/errors"
	"strings"
)

// handleError handle errors returned by facebook api requests
func handleError(err error) error {
	if err == nil {
		return nil
	} else if strings.Contains(err.Error(), "Invalid OAuth access token") {
		return errors.New("invalid access token")
	} else if strings.Contains(err.Error(), "Unsupported delete request.") {
		return errors.New("unsupported delete request")
	}
	return errors.From(err)
}
