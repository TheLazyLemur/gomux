package main

import (
	"flag"
	"fmt"
	"os"
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

func ParseConfig(config string) {
	homeDir, _ := os.UserHomeDir()

	fileBytes, _ := os.ReadFile(homeDir + "/.gomux/" + config + ".conf")
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	for _, line := range lines {
		fragments := strings.Split(line, " ")

		if fragments[0] == NEWSESSION {
			if HasSession(fragments[1]) {
				fmt.Println("Attaching to existing session: ", fragments[1])
				AttachSession(fragments[1])
				return
			} else {
				NewSession(fragments[1], fragments[2])
			}
		}

		if fragments[0] == SPLITPANE {
			SplitPanes(fragments[1], fragments[2], fragments[3], fragments[4])
		}

		if fragments[0] == SELECTPANE {
			SelectPane(fragments[1])
		}

		if fragments[0] == SENDKEYS {
			SendKeys(fragments[1], fragments[2], fragments[3:]...)
		}

		if fragments[0] == ATTACH {
			AttachSession(fragments[1])
		}
	}
}

func FileExists(filename string) bool {
	homeDir, _ := os.UserHomeDir()

	f := homeDir + "/.gomux/" + filename + ".conf"
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
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
