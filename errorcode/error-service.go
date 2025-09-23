package errorcode

type ErrorService struct {
	InternalError error
	ErrorCode     *errorCode
}

type errorCode struct {
	errCode string
	message string
}

func (err *errorCode) GetErrCode() string {
	return err.errCode
}

func (err *errorCode) GetMessage() string {
	return err.message
}

func (errSer *ErrorService) Error() string {
	if errSer.InternalError != nil {
		return errSer.InternalError.Error()
	}
	return ""
}
