package handler

type Error struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func EntityNotFound(m string) *Error {
	return &Error{Status: 404, Code: "NOT_FOUND", Message: m}
}

func BadRequest(m string) *Error {
	return &Error{Status: 400, Code: "BAD_REQUEST", Message: m}
}

func Exception(m string) *Error {
	return &Error{Status: 500, Code: "INTERNAL_SERVER", Message: m}
}
