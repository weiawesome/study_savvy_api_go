package utils

type ExistSource struct {
	Message string
}

func (err ExistSource) Error() string {
	return err.Message
}

type NotExistSource struct {
	Message string
}

func (err NotExistSource) Error() string {
	return err.Message
}

type DbError struct {
	Message string
}

func (err DbError) Error() string {
	return err.Message
}
