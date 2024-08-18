package exception

type BadRequest struct {
	Message string
}

func NewBadRequest(message string) BadRequest {
	return BadRequest{Message: message}
}

func (e BadRequest) Error() string {
	return e.Message
}

type Unauthorized struct {
	Message string
}

func NewUnauthorized(message string) Unauthorized {
	return Unauthorized{Message: message}
}

func (e Unauthorized) Error() string {
	return e.Message
}

type NotFound struct {
	Message string
}

func NewNotFound(message string) NotFound {
	return NotFound{Message: message}
}

func (e NotFound) Error() string {
	return e.Message
}
