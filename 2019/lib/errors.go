package lib

import (
	"fmt"
	"os"
)

func Abort(format string, args ...interface{}) {
	LogError(format, args...)
	os.Exit(1)
}

func LogError(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Errorf(format, args...).Error())
}
