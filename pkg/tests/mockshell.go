package tests

// MockShell represents a mocked cmd.Shell interface for testing
type MockShell struct {
	// params
	ParamCommand string

	// callers
	CalledExec bool

	// returns
	ReturnString string
	ReturnError  error
}

// NewMockShell instantiates a new MockShell
func NewMockShell(options ...func(*MockShell)) *MockShell {

	s := MockShell{}

	// setup any options
	for _, option := range options {
		option(&s)
	}

	return &s
}

// Exec processes a mock command to a shell.
func (s *MockShell) Exec(command string) (string, error) {
	s.ParamCommand = command
	s.CalledExec = true
	return s.ReturnString, s.ReturnError
}
