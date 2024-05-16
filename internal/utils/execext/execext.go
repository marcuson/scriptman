package execext

import (
	"regexp"
	"slices"
)

func StrToArgs(cmdStr string) []string {
	split := regexp.MustCompile(`\s`).Split(cmdStr, -1)
	return slices.DeleteFunc(split, func(s string) bool { return s == "" })
}
