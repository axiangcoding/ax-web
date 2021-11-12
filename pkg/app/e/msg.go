package e

// MsgFlags 这里是业务错误码对应的msg
var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "err",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
