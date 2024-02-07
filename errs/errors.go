package errs

// AppError is a struct that contains the error code and message
type AppError struct {
	Code    int
	Message string
	Status  int
}

func (a *AppError) Error() string {
	return a.Message
}

// MercError is a struct that contains the `Msg` field that stores the error message. MercError implements the
// AppError` struct by defining an `Error()` method that returns the error message.
type MercError struct {
	Msg string
}

func (e *MercError) Error() string {
	return e.Msg
}
