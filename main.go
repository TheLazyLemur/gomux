package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	SplitH = iota
	SplitV
)

func runCmd(showOutput bool, c ...string) {
	cmd := exec.Command(c[0], c[1:]...)

	if showOutput {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

func NewSession(sessionName string, rootDir string) {
	runCmd(false, "tmux", "new-session", "-d", "-s", sessionName, "-c", rootDir)
}

func SplitPanes(sessionName string, rootDir string, hOrV int, percentage string) {
	splitFlag := ""

	if hOrV == SplitH {
		splitFlag = "-h"
	}

	if hOrV == SplitV {
		splitFlag = "-v"
	}

	runCmd(true, "tmux", "split-window", "-d", "-t", sessionName, splitFlag, "-p", percentage, "-c", rootDir)
}

func SendKeys(sessionName string, windowIndex int, command string) {
	runCmd(true, "tmux", "send-keys", "-t", fmt.Sprintf("%s:%d", sessionName, windowIndex), command, "C-m")
}

func SelectPane(paneIndex int) {
	runCmd(true, "tmux", "select-pane", "-t", fmt.Sprintf("%d", paneIndex))
}

func AttachSession(sessionName string) {
	runCmd(true, "tmux", "attach-session", "-t", sessionName)
}

func main() {
	NewSession("hello", "/home/dan/Workspace/Goly")
	SplitPanes("hello", "/home/dan/Workspace/Goly", SplitH, "20")
	SplitPanes("hello", "/home/dan/Workspace/Goly", SplitV, "20")
	SendKeys("hello", 0, "nvim")
	SelectPane(1)
	SendKeys("hello", 0, "top")
	SelectPane(2)
	SendKeys("hello", 0, "lazygit")
	SelectPane(0)
	SplitPanes("hello", "/home/dan/Workspace/Goly", SplitV, "10")
	SelectPane(0)
	AttachSession("hello")
}
