package emerald

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	black = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white
	defaultt = 9

	normalIntensityFG = 30
	highIntensityFG   = 90
	normalIntensityBG = 40
	highIntensityBG   = 100

	start         = "\033["
	normal        = "0;"
	bold          = "1;"
	dim           = "2;"
	underline     = "4;"
	blink         = "5;"
	inverse       = "7;"
	strikethrough = "9;"

	// Reset is the ANSI reset escape sequence
	Reset = "\033[0m"
	// DefaultBG is the default background
	DefaultBG = "\033[49m"
	// DefaultFG is the default foreground
	DefaultFG = "\033[39m"
)

// Black FG
var Black string

// Red FG
var Red string

// Green FG
var Green string

// Yellow FG
var Yellow string

// Blue FG
var Blue string

// Magenta FG
var Magenta string

// Cyan FG
var Cyan string

// White FG
var White string

// LightBlack FG
var LightBlack string

// LightRed FG
var LightRed string

// LightGreen FG
var LightGreen string

// LightYellow FG
var LightYellow string

// LightBlue FG
var LightBlue string

// LightMagenta FG
var LightMagenta string

// LightCyan FG
var LightCyan string

// LightWhite FG
var LightWhite string

var (
	plain = false
	// Colors maps common color names to their ANSI color code.
	Colors = map[string]int{
		"black":   black,
		"red":     red,
		"green":   green,
		"yellow":  yellow,
		"blue":    blue,
		"magenta": magenta,
		"cyan":    cyan,
		"white":   white,
		"default": defaultt,
	}
)

func init() {
	for i := 0; i < 256; i++ {
		Colors[strconv.Itoa(i)] = i
	}

	Black = ColorCode("black")
	Red = ColorCode("red")
	Green = ColorCode("green")
	Yellow = ColorCode("yellow")
	Blue = ColorCode("blue")
	Magenta = ColorCode("magenta")
	Cyan = ColorCode("cyan")
	White = ColorCode("white")
	LightBlack = ColorCode("black+h")
	LightRed = ColorCode("red+h")
	LightGreen = ColorCode("green+h")
	LightYellow = ColorCode("yellow+h")
	LightBlue = ColorCode("blue+h")
	LightMagenta = ColorCode("magenta+h")
	LightCyan = ColorCode("cyan+h")
	LightWhite = ColorCode("white+h")
}

// ColorCode returns the ANSI color code for style.
func ColorCode(style string) string {
	return colorCode(style).String()
}

// Gets the ANSI color code for a style.
func colorCode(style string) *bytes.Buffer {
	buf := bytes.NewBufferString("")
	if plain || style == "" {
		return buf
	}
	if style == "reset" {
		buf.WriteString(Reset)
		return buf
	} else if style == "off" {
		return buf
	}

	foregroundBackground := strings.Split(style, ":")
	foreground := strings.Split(foregroundBackground[0], "+")
	fgKey := foreground[0]
	fg := Colors[fgKey]
	fgStyle := ""
	if len(foreground) > 1 {
		fgStyle = foreground[1]
	}

	bg, bgStyle := "", ""

	if len(foregroundBackground) > 1 {
		background := strings.Split(foregroundBackground[1], "+")
		bg = background[0]
		if len(background) > 1 {
			bgStyle = background[1]
		}
	}

	buf.WriteString(start)
	base := normalIntensityFG
	buf.WriteString(normal) // reset any previous style
	if len(fgStyle) > 0 {
		if strings.Contains(fgStyle, "b") {
			buf.WriteString(bold)
		}
		if strings.Contains(fgStyle, "d") {
			buf.WriteString(dim)
		}
		if strings.Contains(fgStyle, "B") {
			buf.WriteString(blink)
		}
		if strings.Contains(fgStyle, "u") {
			buf.WriteString(underline)
		}
		if strings.Contains(fgStyle, "i") {
			buf.WriteString(inverse)
		}
		if strings.Contains(fgStyle, "s") {
			buf.WriteString(strikethrough)
		}
		if strings.Contains(fgStyle, "h") {
			base = highIntensityFG
		}
	}

	// parse color
	n, err := strconv.Atoi(fgKey)
	if err == nil {
		if n < 8 {
			fmt.Fprintf(buf, "%d;", n+30)
		} else if n < 16 {
			fmt.Fprintf(buf, "%d;", n+82)
		} else {
			fmt.Fprintf(buf, "38;5;%d;", n)
		}
	} else if strings.HasPrefix(fgKey, "#") && len(fgKey) == 7 {
		r, _ := strconv.ParseInt(fgKey[1:3], 16, 64)
		g, _ := strconv.ParseInt(fgKey[3:5], 16, 64)
		b, _ := strconv.ParseInt(fgKey[5:7], 16, 64)
		fmt.Fprintf(buf, "38;2;%d;%d;%d;", r, g, b)
	} else {
		fmt.Fprintf(buf, "%d;", base+fg)
	}

	base = normalIntensityBG
	if len(bg) > 0 {
		if strings.Contains(bgStyle, "h") {
			base = highIntensityBG
		}
		// parse color
		n, err := strconv.Atoi(bg)
		if err == nil {
			if n < 8 {
				fmt.Fprintf(buf, "%d;", n+40)
			} else if n < 16 {
				fmt.Fprintf(buf, "%d;", n+92)
			} else {
				fmt.Fprintf(buf, "48;5;%d;", n)
			}
		} else if strings.HasPrefix(bg, "#") && len(bg) == 7 {
			r, _ := strconv.ParseInt(bg[1:3], 16, 64)
			g, _ := strconv.ParseInt(bg[3:5], 16, 64)
			b, _ := strconv.ParseInt(bg[5:7], 16, 64)
			fmt.Fprintf(buf, "48;2;%d;%d;%d;", r, g, b)
		} else {
			fmt.Fprintf(buf, "%d;", base+Colors[bg])
		}
	}

	// remove last ";"
	buf.Truncate(buf.Len() - 1)
	buf.WriteRune('m')
	return buf
}

// Color colors a string based on the ANSI color code for style.
func Color(s, style string) string {
	if plain || len(style) < 1 {
		return s
	}
	buf := colorCode(style)
	buf.WriteString(s)
	buf.WriteString(Reset)
	return buf.String()
}

// Colorizer takes an input string and colorizes it based on the style it was created with.
type Colorizer func(s string) string

// ColorFunc creates a closure to avoid computation ANSI color code.
func ColorFunc(style string) Colorizer {
	if style == "" {
		return func(s string) string {
			return s
		}
	}
	color := ColorCode(style)
	return func(s string) string {
		if plain || s == "" {
			return s
		}
		buf := bytes.NewBufferString(color)
		buf.WriteString(s)
		buf.WriteString(Reset)
		result := buf.String()
		return result
	}
}

// ColorFuncVar creates a closure to avoid computation ANSI color code and returns ANSI sequence.
func ColorFuncVar(style string) (Colorizer, string) {
	if style == "" {
		return func(s string) string {
			return s
		}, ""
	}
	color := ColorCode(style)
	return func(s string) string {
		if plain || s == "" {
			return s
		}
		buf := bytes.NewBufferString(color)
		buf.WriteString(s)
		buf.WriteString(Reset)
		result := buf.String()
		return result
	}, color
}

func ColorInterface(style string) ColorPrinter {
	color := ColorCode(style)
	return ColorPrinter{ANSI: color}
}

type ColorPrinter struct {
	ANSI string
}

func (c ColorPrinter) Print(a ...interface{}) {
	Print(c.ANSI)
	Print(a...)
	Print(Reset)
}

func (c ColorPrinter) Printf(format string, a ...interface{}) {
	Print(c.ANSI)
	Printf(format, a...)
	Print(Reset)
}

func (c ColorPrinter) Println(a ...interface{}) {
	Print(c.ANSI)
	Print(a...)
	Println(Reset)
}

// DisableColors disables ANSI color codes. The default is false (colors are on).
func disableAnsiColors(disable bool) {
	if plain == disable {
		return
	}
	plain = disable
	if plain {
		Black = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Magenta = ""
		Cyan = ""
		White = ""
		LightBlack = ""
		LightRed = ""
		LightGreen = ""
		LightYellow = ""
		LightBlue = ""
		LightMagenta = ""
		LightCyan = ""
		LightWhite = ""
	} else {
		Black = ColorCode("black")
		Red = ColorCode("red")
		Green = ColorCode("green")
		Yellow = ColorCode("yellow")
		Blue = ColorCode("blue")
		Magenta = ColorCode("magenta")
		Cyan = ColorCode("cyan")
		White = ColorCode("white")
		LightBlack = ColorCode("black+h")
		LightRed = ColorCode("red+h")
		LightGreen = ColorCode("green+h")
		LightYellow = ColorCode("yellow+h")
		LightBlue = ColorCode("blue+h")
		LightMagenta = ColorCode("magenta+h")
		LightCyan = ColorCode("cyan+h")
		LightWhite = ColorCode("white+h")
	}
}
