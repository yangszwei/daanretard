package errors_test

import (
	"daanretard/internal/infra/errors"
	"testing"
)

func TestNew(t *testing.T) {
	_ = errors.New("test error")
	_ = errors.New("test error", "test message")
}

func TestFrom(t *testing.T) {
	_ = errors.From(errors.New("test error"))
}

func TestError_SetMessage(t *testing.T) {
	_ = errors.New("test error").SetMessage("test message")
}

func TestError_Error(t *testing.T) {
	_ = errors.New("test error").Error()
	_ = errors.New("test error", "test message").Error()
}

func TestErrors_Error(t *testing.T) {
	_ = errors.Errors{
		errors.New("test error", "#1"),
		errors.New("test error", "#2"),
	}.Error()
}
