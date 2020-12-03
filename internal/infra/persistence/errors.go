package persistence

import (
	"daanretard/internal/infra/errors"
	"strings"
)

// parseError parse sql error into errors.Error
func parseError(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "Error 1406") {
		e := strings.Split(err.Error(), "'")
		return errors.New("data too lang", e[1])
	} else if strings.Contains(err.Error(), "Error 1062") {
		e := strings.Split(err.Error(), "'")
		return errors.ErrDuplicateEntry.SetMessage(e[3])
	} else if err.Error() == "record not found" {
		return errors.ErrRecordNotFound
	}
	return errors.From(err)
}
