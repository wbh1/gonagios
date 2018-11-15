package nagios

import "testing"

func TestResult(*testing.T) {
	result := Result{}

	result.ExitCode = OK

	result.Msg = "Disk Usage is at 72%"

	result.Exit()
}
