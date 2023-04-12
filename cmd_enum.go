package main

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
