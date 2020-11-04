package actions

import (
	"os"
	"strings"
)

func GetInput(s string) string {
	i := os.Getenv(strings.ToUpper("INPUT_" + s))
	return i
}
