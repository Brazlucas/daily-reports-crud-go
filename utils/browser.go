package utils

import (
	"os"
	"os/exec"
)

func OpenBrowser(url string) {
	var cmd string
	var args []string

	switch os := os.Getenv("OS"); os {
	case "Windows_NT":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}

	exec.Command(cmd, args...).Start()
}
