package e

// MsgFlags Code message map
var MsgFlags = map[int]string{
	Success:               "ok",
	Error:                 "err",
	RequestParamsNotValid: "request params not valid",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
