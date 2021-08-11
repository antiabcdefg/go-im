package args

const (
	UNKNOWN_ERROR      = 10000
	CODE_NOT_LOGIN     = 20001
	CODE_LOGIN_FAIL    = 20002
	CODE_REGISTER_FAIL = 30001
)

type ErrorArg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
