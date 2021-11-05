package emerald

import (
	"fmt"
	"github.com/jcwillox/emerald/emeraldp"
)

const (
	ShowCursorSeq      = "\x1b[?25h"
	HideCursorSeq      = "\x1b[?25l"
	EnableLineWrapSeq  = "\x1b[?7h"
	DisableLineWrapSeq = "\x1b[?7l"

	CursorUpSeq    = "\x1b[%dA"
	CursorRightSeq = "\x1b[%dG"
	CursorDownSeq  = "\x1b[%dB"
)

func CursorUp(lines int) {
	emeraldp.Printf(CursorUpSeq, lines)
}

func CursorUpVar(lines int) string {
	return fmt.Sprintf(CursorUpSeq, lines)
}

func CursorDown(lines int) {
	emeraldp.Printf(CursorDownSeq, lines)
}

func CursorDownVar(lines int) string {
	return fmt.Sprintf(CursorDownSeq, lines)
}

func CursorRight(columns int) {
	emeraldp.Printf(CursorRightSeq, columns)
}

func CursorRightVar(columns int) string {
	return fmt.Sprintf(CursorRightSeq, columns)
}

func ShowCursor() {
	emeraldp.Print(ShowCursorSeq)
}

func HideCursor() {
	emeraldp.Print(HideCursorSeq)
}

func EnableLineWrap() {
	emeraldp.Print(EnableLineWrapSeq)
}

func DisableLineWrap() {
	emeraldp.Print(DisableLineWrapSeq)
}
