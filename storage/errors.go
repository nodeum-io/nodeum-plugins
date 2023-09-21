package storage

// ErrorType defines a type of error on file operation
type ErrorType string

// Enumeration of type of errors
var (
	ErrUnknown        ErrorType = "unknown"
	ErrNotExist       ErrorType = "not-exist"
	ErrNotDir         ErrorType = "not-dir"
	ErrPermission     ErrorType = "permission"
	ErrExist          ErrorType = "exist"
	ErrNotEmpty       ErrorType = "not-empty"
	ErrTimeout        ErrorType = "timeout"
	ErrUnknownCommand ErrorType = "unknown-command"
)

type Error interface {
	error
	Type() ErrorType
}

type OpError struct {
	error
	_type ErrorType
}

func (e OpError) Type() ErrorType {
	return e._type
}

func NewOpError(err error, t ErrorType) Error {
	if err == nil {
		return nil
	}
	return OpError{err, t}
}
