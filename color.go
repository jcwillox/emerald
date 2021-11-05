package emerald

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
)

var ColorEnabled bool

type ColorState int

const (
	ColorDisabled ColorState = iota
	ColorForced
	ColorUnset
)

// IsTerminal returns true if stdout is attached to a terminal
func IsTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}

// ColorPreference returns the color preference based on the environment variables
func ColorPreference() ColorState {
	if _, present := os.LookupEnv("NO_COLOR"); present {
		return ColorDisabled
	}
	if val := os.Getenv("CLICOLOR"); val == "0" {
		return ColorDisabled
	}
	if val, present := os.LookupEnv("CLICOLOR_FORCE"); present && val != "0" {
		return ColorForced
	}
	return ColorUnset
}

// IsColorTerm combines IsColor and IsTerminal
func IsColorTerm() bool {
	state := ColorPreference()
	if state != ColorUnset {
		return state != 0
	}
	return IsTerminal()
}

// ColorIndexFg convert 256 color code to ANSI foreground color
func ColorIndexFg(index int) string {
	if index < 8 {
		return fmt.Sprintf("\x1b[%dm", index+30)
	}
	if index < 16 {
		return fmt.Sprintf("\x1b[%dm", index+82)
	}
	return fmt.Sprintf("\x1b[38;5;%dm", index)
}

// ColorIndexBg convert 256 color code to ANSI background color
func ColorIndexBg(index int) string {
	if index < 8 {
		return fmt.Sprintf("\x1b[%dm", index+40)
	}
	if index < 16 {
		return fmt.Sprintf("\x1b[%dm", index+92)
	}
	return fmt.Sprintf("\x1b[48;5;%dm", index)
}

// AutoSetColorState automatically SetColorState based on the whether
// we are outputting to a terminal or the environment variables
func AutoSetColorState() {
	SetColorState(IsColorTerm())
}

// SetColorState enables or disables ANSI support in the library using
// this function is preferred to directly setting `ColorEnabled` as
// it allows us to perform other actions.
func SetColorState(enabled bool) {
	ColorEnabled = enabled
	disableAnsiColors(!enabled)
}
