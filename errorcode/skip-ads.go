package errorcode

const (
	CodeErrUserSkipAdsInsufficient = "ERR_USER_SKIP_ADS_INSUFFICIENT"
)

var (
	ErrUserSkipAdsInsufficient = &errorCode{
		errCode: CodeErrUserSkipAdsInsufficient,
		message: "Error user skip ads insufficient",
	}
)
