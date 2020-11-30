package generator

type ExitError struct {
	ExitCode int
	Message  string
}

func (e ExitError) Error() string {
	return e.Message
}

const (
	Successful int = iota
	ConfigurationError
	FileReadError
	SpecificationError
	FileWriteError
)
