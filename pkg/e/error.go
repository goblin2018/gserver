package e

type Error interface {
	Error() string
	Code() uint32
	Set(msg string)
	Add(msg string)
	From(err error) Error
}

type mErr struct {
	code uint32
	msg  string
}

func (e mErr) Error() string {
	return e.msg
}

func (e mErr) Code() uint32 {
	return e.code
}

func (e *mErr) Set(msg string) {
	e.msg = msg
}

func (e *mErr) Add(msg string) {
	e.msg = e.msg + msg
}

func (e *mErr) From(err error) Error {
	if err == nil {
		return OK
	}
	e.msg = err.Error()
	return e
}

func newErr(code uint32, msg string) Error {
	return &mErr{
		msg:  msg,
		code: code,
	}
}
