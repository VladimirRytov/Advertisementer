package systeminfo

import (
	"os/exec"
)

func ComputerName() (string, error) {
	cat := exec.Command("cat", "/etc/hostname")
	str, err := cat.Output()
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func CPU() (string, error) {
	command := exec.Command(`bash`, `-c`, `cat /proc/cpuinfo | grep -E "model name" | head -n 1 | cut -d':' -f 2`)
	out, err := command.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
