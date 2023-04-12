package main

import (
	"os"
	"os/exec"
)

func FileExists(filename string) bool {
	homeDir, _ := os.UserHomeDir()

	f := homeDir + "/.gomux/" + filename + ".conf"
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

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
