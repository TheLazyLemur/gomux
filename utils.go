package main

import (
	"os"
	"os/exec"
)

func runCmd(showOutput bool, c ...string) error {
	cmd := exec.Command(c[0], c[1:]...)

	if showOutput {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
