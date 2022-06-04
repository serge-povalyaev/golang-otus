package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	successCode = 0
	errorCode   = 1
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	for envName, envValue := range env {
		command.Env = append(command.Env, fmt.Sprintf("%s=%s", envName, envValue.Value))
	}

	err := command.Run()
	if err != nil {
		return errorCode
	}

	return successCode
}
