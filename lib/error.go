package lib

import (
	"fmt"
	"os"
)

func Abort(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
