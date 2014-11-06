package cmd

// osshell is a global variable that is implement on a per-OS basis.
var osshell Shell

// Shell is the interface to an OS-agnostic shell to call commands.
type Shell interface {
	// Exec executes the command with the system's configured osshell.
	Exec(command string) (string, error)
}

// New returns the current instance of the underlying shell.
func New() Shell {
	return osshell
}

// Exec executes the command with the system's configured osshell.
func Exec(command string) (string, error) {
	return osshell.Exec(command)
}
