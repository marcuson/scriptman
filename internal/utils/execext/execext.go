package execext

import "regexp"

func StrToArgs(cmdStr string) []string {
	return regexp.MustCompile(`\s`).Split(cmdStr, -1)
}
