package main

import (
	"fmt"
	"strings"
)

type SplitCmd int

const (
	SPLITH SplitCmd = iota
	SPLITV
)

func (bp SplitCmd) String() string {
	return []string{"h", "v"}[bp]
}

func NewSession(sessionName string, rootDir string) {
	runCmd(false, "tmux", "new-session", "-d", "-s", sessionName, "-c", rootDir)
}

func SplitPanes(sessionName string, rootDir string, hOrV string, percentage string) {
	splitFlag := ""

	switch hOrV {
	case SPLITH.String():
		splitFlag = "-h"
	case SPLITV.String():
		splitFlag = "-v"
	}

	runCmd(true, "tmux", "split-window", "-d", "-t", sessionName, splitFlag, "-p", percentage, "-c", rootDir)
}

func SendKeys(sessionName string, windowIndex string, command ...string) {
	x := fmt.Sprintf("%s:%s", sessionName, windowIndex)
	y := []string{
		"tmux",
		"send-keys",
		"-t",
		x,
		strings.Join(command, " "),
		"C-m",
	}
	runCmd(true, y...)
}

func SelectPane(paneIndex string) {
	runCmd(true, "tmux", "select-pane", "-t", fmt.Sprintf("%s", paneIndex))
}

func AttachSession(sessionName string) {
	runCmd(true, "tmux", "attach-session", "-t", sessionName)
}

func HasSession(sessionName string) bool {
	err := runCmd(false, "tmux", "has-session", "-t", sessionName)
	return err == nil
}
