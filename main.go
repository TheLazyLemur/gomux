package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func LoadConfig(name string) *Config {
	homeDir, _ := os.UserHomeDir()

	fileBytes, _ := os.ReadFile(homeDir + "/.gomux/" + name + ".conf")
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
		return nil
	}

	rootDir := strings.Split(nonEmptyLines[1], " ")[1]
	if len(rootDir) == 0 {
		fmt.Println("Root dir is empty")
		return nil
	}

	cfg := &Config{
		lines:       nonEmptyLines[2:],
		sessionName: sessionName,
		rootDir:     rootDir,
	}

	return cfg
}

func main() {
	var configName string
	flag.StringVar(&configName, "c", "", "config file")
	flag.Parse()

	if len(configName) == 0 {
		fmt.Println("No config file specified")
		os.Exit(1)
	}

	if !FileExists(configName) {
		fmt.Println("Config file does not exist")
		os.Exit(1)
	}

	cfg := LoadConfig(configName)
	cfg.ParseConfig()
}
