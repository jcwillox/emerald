package emerald

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"sort"
	"strings"
	"testing"
)

func TestPlain(t *testing.T) {
	SetColorState(false)
	PrintStyles()
}

func TestStyles(t *testing.T) {
	SetColorState(true)
	PrintStyles()
}

func TestDisableColors(t *testing.T) {
	fn := ColorFunc("red")

	buf := colorCode("off")
	if buf.String() != "" {
		t.Fail()
	}
	SetColorState(false)
	if Black != "" {
		t.Fail()
	}
	code := ColorCode("red")
	if code != "" {
		t.Fail()
	}
	s := fn("foo")
	if s != "foo" {
		t.Fail()
	}
	SetColorState(true)
	if Black == "" {
		t.Fail()
	}
	code = ColorCode("red")
	if code == "" {
		t.Fail()
	}
	// will have escape codes around it
	index := strings.Index(fn("foo"), "foo")
	if index <= 0 {
		t.Fail()
	}
}

func TestAttributeReset(t *testing.T) {
	boldRed := ColorCode("red+b")
	greenUnderline := ColorCode("green+u")
	s := fmt.Sprintf("normal %s bold red %s green underline %s", boldRed, greenUnderline, Reset)
	// See the results on the terminal for regression tests.
	fmt.Printf("Colored string: %s\n", s)
	fmt.Printf("Escaped string: %q\n", s)
	if s != "normal \x1b[0;1;31m bold red \x1b[0;4;32m green underline \x1b[0m" {
		t.Error("Attributes are not being reset")
	}
}

// PrintStyles prints all style combinations to the terminal.
func PrintStyles() {
	// for compatibility with Windows, not needed for *nix
	stdout := colorable.NewColorableStdout()

	bgColors := []string{
		"",
		":black",
		":red",
		":green",
		":yellow",
		":blue",
		":magenta",
		":cyan",
		":white",
	}

	keys := make([]string, 0, len(Colors))
	for k := range Colors {
		keys = append(keys, k)
	}

	sort.Sort(sort.StringSlice(keys))

	for _, fg := range keys {
		for _, bg := range bgColors {
			fmt.Fprintln(stdout, padColor(fg, []string{"" + bg, "+b" + bg, "+bh" + bg, "+u" + bg}))
			fmt.Fprintln(stdout, padColor(fg, []string{"+s" + bg, "+i" + bg}))
			fmt.Fprintln(stdout, padColor(fg, []string{"+uh" + bg, "+B" + bg, "+Bb" + bg /* backgrounds */, "" + bg + "+h"}))
			fmt.Fprintln(stdout, padColor(fg, []string{"+b" + bg + "+h", "+bh" + bg + "+h", "+u" + bg + "+h", "+uh" + bg + "+h"}))
		}
	}
}

func pad(s string, length int) string {
	for len(s) < length {
		s += " "
	}
	return s
}

func padColor(color string, styles []string) string {
	buffer := ""
	for _, style := range styles {
		buffer += Color(pad(color+style, 20), color+style)
	}
	return buffer
}
