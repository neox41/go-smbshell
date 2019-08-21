package exec

import (
	"os/exec"
)

func Shell(args string) string {
	cmd := exec.Command("cmd.exe", "/c", args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}
	return string(out)
}
