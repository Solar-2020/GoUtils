package errorWorker

type ResponseError struct {
	httpCode      int
	responseError error
	fullError     error
}

func (re ResponseError) Error() string {
	return re.responseError.Error()
}
