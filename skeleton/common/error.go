package common

// BaseErr basic error class
type BaseErr struct {
	Msg   string
	Code  int
	Trace string
}

func (e BaseErr) Error() string {
	return e.Msg
}

func NewBaseErr(code int, msg interface{}) error {

	return nil
}
