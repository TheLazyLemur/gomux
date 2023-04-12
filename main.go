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
	SESSIONNAME
	ROOTDIR
)

func (enum Cmd) String() string {
	enumToValue := map[Cmd]string{
		0: "new-session",
		1: "split-pane",
		2: "select-pane",
		3: "send-keys",
		4: "attach-session",
		5: "session-name",
		6: "root-dir",
	}

	e, ok := enumToValue[enum]
	if !ok {
		return ""
	}
	return e
}

func ParseConfig(config string) {
	homeDir, _ := os.UserHomeDir()

	fileBytes, _ := os.ReadFile(homeDir + "/.gomux/" + config + ".conf")
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")
	nonEmptyLines := make([]string, 0)
	for _, line := range lines {
		if len(line) > 0 {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	sessionName := strings.Split(nonEmptyLines[0], " ")[1]
	if len(sessionName) == 0 {
		fmt.Println("Session name is empty")
		return
	}

	rootDir := strings.Split(nonEmptyLines[1], " ")[1]
	if len(rootDir) == 0 {
		fmt.Println("Root dir is empty")
		return
	}

	for _, line := range nonEmptyLines {
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
		case SESSIONNAME.String():
		case ROOTDIR.String():
		default:
			fmt.Println("Unknown command: ", cmd)
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
