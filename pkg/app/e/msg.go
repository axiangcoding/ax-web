package e

// MsgFlags Code message map
var MsgFlags = map[int]string{
	Success:               "ok",
	Error:                 "system error",
	RequestParamsNotValid: "request params not valid",
	TokenNotExist:         "token not exist",
	TokenNotLegal:         "token is required",
	TokenExpired:          "token expired",
	LoginFailed:           "login failed",
	RegisterFailed:        "register failed",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
