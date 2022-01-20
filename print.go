package emerald

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"io"
)

// Wraps standard print methods to use go-colorable which fixes ANSI sequences
// on older Windows machines and has no effect on non-windows platforms.

var (
	Stdout io.Writer
	Stderr io.Writer
)

func init() {
	Stdout = colorable.NewColorableStdout()
	Stderr = colorable.NewColorableStderr()
}

func Print(a ...interface{}) (n int) {
	n, _ = fmt.Fprint(Stdout, a...)
	return n
}

func Printf(format string, a ...interface{}) (n int) {
	n, _ = fmt.Fprintf(Stdout, format, a...)
	return n
}

func Println(a ...interface{}) (n int) {
	n, _ = fmt.Fprintln(Stdout, a...)
	return n
}
