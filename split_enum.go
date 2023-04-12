package main

type SplitCmd int

const (
	SPLITH SplitCmd = iota
	SPLITV
)

func (enum SplitCmd) String() string {
	enumToValue := map[SplitCmd]string{
		0: "h",
		1: "v",
	}

	e, ok := enumToValue[enum]
	if !ok {
		return ""
	}
	return e
}
