package cmd

// osshell is a global variable that is implement on a per-OS basis.
var osshell shell

// shell is the interface to an OS-agnostic shell to call commands.
type shell interface {
	// Exec executes the command with the system's configured osshell.
	Exec(command string) (string, error)
}

// Exec executes the command with the system's configured osshell.
func Exec(command string) (string, error) {
	return osshell.Exec(command)
}
