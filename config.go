package main

import (
	"fmt"
	"strings"
)

type Config struct {
	sessionName string
	rootDir     string
	lines       []string
}

func (cfg *Config) ParseConfig() {

	for _, line := range cfg.lines {
		fragments := strings.Split(line, " ")
		cmd := fragments[0]

		switch cmd {
		case NEWSESSION.String():
			if HasSession(cfg.sessionName) {
				fmt.Println("Attaching to existing session: ", cfg.sessionName)
				AttachSession(cfg.sessionName)
				return
			} else {
				NewSession(cfg.sessionName, cfg.rootDir)
			}
		case SPLITPANE.String():
			SplitPanes(cfg.sessionName, cfg.rootDir, fragments[1], fragments[2])
		case SELECTPANE.String():
			SelectPane(fragments[1])
		case SENDKEYS.String():
			SendKeys(cfg.sessionName, fragments[1], fragments[2:]...)
		case ATTACH.String():
			AttachSession(cfg.sessionName)
		case SESSIONNAME.String():
		case ROOTDIR.String():
		case COMMENT.String():
		default:
			fmt.Println("Unknown command: ", cmd)
		}
	}
}
