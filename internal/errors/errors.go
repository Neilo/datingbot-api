package errors

//Error implement error interface
type Error struct {
	DevMSG string `json:"-"`
	MSG    string `json:"msg"`
	Status int    `json:"-"`
}

//Error return dev msg
func (e Error) Error() string {
	return e.DevMSG
}

//New generate internal error from other errors
func (e Error) New(err error) *Error {
	e.DevMSG = err.Error()
	return &e
}
