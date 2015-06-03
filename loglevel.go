package logger

import "strings"

type LogLevel int

const (
	DEBUG = 1 + iota
	INFO
	WARNING
	ERROR
	FATAL
)

var levels = [...]string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
}

func (level LogLevel) String() string {
	return levels[level-1]
}

func levelFromString(s string) LogLevel {
	// remove = if present
	if string(s[0]) == "=" {
		s = string(s[1:])
	}

	for i, l := range levels {
		if strings.ToUpper(s) == l {
			return LogLevel(i + 1)
		}
	}
	panic("Unable to find Level from string:" + s)
}
