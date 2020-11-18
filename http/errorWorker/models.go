package errorWorker

type ResponseError struct {
	httpCode      int
	responseError error
	fullError     error
}

func (re ResponseError) Error() string {
	return re.responseError.Error()
}

func (re ResponseError) ResponseError() error {
	return re.responseError
}

func (re ResponseError) FullError() error {
	return re.fullError
}
