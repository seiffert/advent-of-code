package lib

import (
	"fmt"
	"os"
)

func Abort(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Errorf(format, args...).Error())
	os.Exit(1)
}
