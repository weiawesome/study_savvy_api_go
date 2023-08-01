package utils

type Error struct {
	Error string `json:"error"`
}
type RegistrationError struct {
	Message string
}

func (err RegistrationError) Error() string {
	return err.Message
}
