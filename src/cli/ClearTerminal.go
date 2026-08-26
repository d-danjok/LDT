package cli

import (
	"os"
	"os/exec"
	"runtime"
)

// ClearTerminal clears the terminal screen based on the OS
func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows command
	} else {
		cmd = exec.Command("clear") // Unix/Linux/Mac command
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
