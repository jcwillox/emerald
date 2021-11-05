package emerald

import (
	"fmt"
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
	Printf(CursorUpSeq, lines)
}

func CursorUpVar(lines int) string {
	return fmt.Sprintf(CursorUpSeq, lines)
}

func CursorDown(lines int) {
	Printf(CursorDownSeq, lines)
}

func CursorDownVar(lines int) string {
	return fmt.Sprintf(CursorDownSeq, lines)
}

func CursorRight(columns int) {
	Printf(CursorRightSeq, columns)
}

func CursorRightVar(columns int) string {
	return fmt.Sprintf(CursorRightSeq, columns)
}

func ShowCursor() {
	Print(ShowCursorSeq)
}

func HideCursor() {
	Print(HideCursorSeq)
}

func EnableLineWrap() {
	Print(EnableLineWrapSeq)
}

func DisableLineWrap() {
	Print(DisableLineWrapSeq)
}
