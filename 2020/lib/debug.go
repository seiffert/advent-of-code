package lib

import (
	"fmt"
	"os"
)

func Debug(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
