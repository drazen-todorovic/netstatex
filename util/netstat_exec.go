package util

import (
	"errors"
	"os/exec"
)

// TryExecuteNetstat executes netstat commnad with -an args
func TryExecuteNetstat() (string, error) {
	binary, err := exec.LookPath("netstat")

	if err != nil {
		return "", errors.New("Cannot find netstat executable")
	}

	result, err := exec.Command(binary, "-an").Output()

	if err != nil {
		return "", err
	}

	return string(result), nil
}
