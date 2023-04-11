package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	NEWSESSION = "new-session"
	SPLITPANE  = "split-pane"
	SELECTPANE = "select-pane"
	SENDKEYS   = "send-keys"
	ATTACH     = "attach-session"

	SPLITV = "v"
	SPLITH = "h"
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

func SplitPanes(sessionName string, rootDir string, hOrV string, percentage string) {
	splitFlag := ""

	if hOrV == SPLITH {
		splitFlag = "-h"
	}

	if hOrV == SPLITV {
		splitFlag = "-v"
	}

	runCmd(true, "tmux", "split-window", "-d", "-t", sessionName, splitFlag, "-p", percentage, "-c", rootDir)
}

func SendKeys(sessionName string, windowIndex string, command string) {
	runCmd(true, "tmux", "send-keys", "-t", fmt.Sprintf("%s:%s", sessionName, windowIndex), command, "C-m")
}

func SelectPane(paneIndex string) {
	runCmd(true, "tmux", "select-pane", "-t", fmt.Sprintf("%s", paneIndex))
}

func AttachSession(sessionName string) {
	runCmd(true, "tmux", "attach-session", "-t", sessionName)
}

func ParseConfig(config string) {
	fileBytes, _ := os.ReadFile("/home/dan/.gomux/" + config + ".conf")
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	for _, line := range lines {
		fragments := strings.Split(line, " ")

		if fragments[0] == NEWSESSION {
			NewSession(fragments[1], fragments[2])
		}

		if fragments[0] == SPLITPANE {
			SplitPanes(fragments[1], fragments[2], fragments[3], fragments[4])
		}

		if fragments[0] == SELECTPANE {
			SelectPane(fragments[1])
		}

		if fragments[0] == SENDKEYS {
			SendKeys(fragments[1], fragments[2], fragments[3])
		}

		if fragments[0] == ATTACH {
			AttachSession(fragments[1])
		}
	}
}

func main() {
	var config string
	flag.StringVar(&config, "c", "", "config file")
	flag.Parse()
	ParseConfig(config)
}
