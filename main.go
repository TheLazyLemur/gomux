package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Cmd int

const (
	NEWSESSION Cmd = iota
	SPLITPANE
	SELECTPANE
	SENDKEYS
	ATTACH
)

func (bp Cmd) String() string {
	return []string{"new-session", "split-pane", "select-pane", "send-keys", "attach-session"}[bp]
}

func ParseConfig(config string) {
	homeDir, _ := os.UserHomeDir()

	fileBytes, _ := os.ReadFile(homeDir + "/.gomux/" + config + ".conf")
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	sessionName := strings.Split(lines[0], " ")[1]
	if len(sessionName) == 0 {
		fmt.Println("Session name is empty")
		return
	}

	rootDir := strings.Split(lines[1], " ")[1]
	if len(rootDir) == 0 {
		fmt.Println("Root dir is empty")
		return
	}

	for _, line := range lines {
		fragments := strings.Split(line, " ")
		cmd := fragments[0]

		switch cmd {
		case NEWSESSION.String():
			if HasSession(sessionName) {
				fmt.Println("Attaching to existing session: ", fragments[1])
				AttachSession(sessionName)
				return
			} else {
				NewSession(sessionName, rootDir)
			}
		case SPLITPANE.String():
			SplitPanes(sessionName, rootDir, fragments[1], fragments[2])
		case SELECTPANE.String():
			SelectPane(fragments[1])
		case SENDKEYS.String():
			SendKeys(sessionName, fragments[1], fragments[2:]...)
		case ATTACH.String():
			AttachSession(sessionName)
		}
	}
}

func main() {
	var config string
	flag.StringVar(&config, "c", "", "config file")
	flag.Parse()

	if len(config) == 0 {
		fmt.Println("No config file specified")
		os.Exit(1)
	}

	if !FileExists(config) {
		fmt.Println("Config file does not exist")
		os.Exit(1)
	}

	ParseConfig(config)
}
