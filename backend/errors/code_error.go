package errors

type HTTPCodeErr struct {
	Code int
	Err  error
}

func (e *HTTPCodeErr) Error() string {
	return e.Err.Error()
}

func (e *HTTPCodeErr) HTTPCode() int {
	return e.Code
}
