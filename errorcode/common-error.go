package errorcode

// Define common error code here
// FORMAT OF ERROR CODE: CodeErr{ErrorName}
const (
	CodeErrUnknown            = "ERR_UNKNOWN"
	CodeErrSystem             = "ERR_SYSTEM"
	CodeErrAuth               = "ERR_AUTH"
	CodeErrDataRequestInvalid = "ERR_DATA_REQUEST_INVALID"
)

// Common error
// FORMAT OF ERROR: Err{ErrorName}
var (
	ErrUnknown = &errorCode{
		errCode: CodeErrUnknown,
		message: "Call api failed. Could you please try again ?",
	}
	ErrSystem = &errorCode{
		errCode: CodeErrSystem,
		message: "Oops! Looks like we stumbled upon an error. Could you please try again ?",
	}
	ErrInvalidRequest = &errorCode{
		errCode: CodeErrDataRequestInvalid,
		message: "Data request invalid",
	}
	ErrAuth = &errorCode{
		errCode: CodeErrAuth,
		message: "UnAuthorization, please login again",
	}
)
