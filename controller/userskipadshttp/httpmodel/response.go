package httpmodel

type Response struct {
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`  // for client show to user
	Reason   string      `json:"reason"`   // error code type string
	Metadata interface{} `json:"metadata"` // some information for debug like trace id, raw error
}
