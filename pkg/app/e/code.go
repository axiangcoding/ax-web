package e

const (
	Success               = 00000
	Error                 = 10000
	RequestParamsNotValid = 10001

	TokenNotExist = 11000
	TokenNotLegal = 11001
	TokenExpired  = 11002

	LoginFailed    = 12000
	RegisterFailed = 12001
)

var errCodeText = map[int]string{
	Success:               "OK",
	Error:                 "System error",
	RequestParamsNotValid: "Request params not valid",

	TokenNotExist: "Token not exist",
	TokenNotLegal: "Token is required",
	TokenExpired:  "Token expired",

	LoginFailed:    "Login failed",
	RegisterFailed: "Register failed",
}

func CodeText(code int) string {
	return errCodeText[code]
}
