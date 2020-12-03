package errors

// Default errors
var (
	ErrRecordNotFound = New("record not found")
	ErrDuplicateEntry = New("duplicate entry")
	ErrUnknownError   = New("unknown error")
)
