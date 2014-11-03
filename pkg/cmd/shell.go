package cmd

// shell is a global variable that is implement on a per-OS basis.
var shell CmdShell

// CmdShell is the interface to an OS-agnostic shell to call commands.
type CmdShell interface {
	Exec(command string) (string, error)
}
