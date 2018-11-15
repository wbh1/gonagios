package nagios

import (
	"fmt"
	"os"
)

type status int

const (
	// OK is exit code 0
	OK status = iota

	// WARNING is exit code 1
	WARNING

	// CRITICAL is exit code 2
	CRITICAL

	// UNKNOWN is exit code 3
	UNKNOWN
)

// Result represents a Nagios Result with an exit code (e.g. nagios.OK)
// and a message to print after it
type Result struct {
	ExitCode status
	Msg      string
}

// NewResult creates a Nagios result with initial defaults
func NewResult() *Result {
	return &Result{
		ExitCode: UNKNOWN,
		Msg:      "<no output>",
	}
}

func (e status) string() string {
	returnCodes := []string{
		"OK",
		"WARNING",
		"CRITICAL",
		"UNKNOWN",
	}

	// default to returning UNKNOWN if status code isn't in 0-3
	if e < OK || e > UNKNOWN {
		return "UNKNOWN"
	}

	return returnCodes[e]
}

func (e status) int() int {
	return int(e)
}

// Exit prints the Nagios status and message, then exits with the corresponding exit code
func (nr *Result) Exit() {
	fmt.Println(nr.ExitCode.string(), "—", nr.Msg)
	os.Exit(nr.ExitCode.int())
}
