# Emerald

A basic color library for use in my Go projects, built on top of [mgutz/ansi](https://github.com/mgutz/ansi).

Package ansi is a small, fast library to create ANSI colored strings and codes.

## Install

Get it

```sh
go get -u github.com/jcwillox/emerald
```

## Example

```go
import "github.com/jcwillox/emerald"

// colorize a string, SLOW
msg := emerald.Color("foo", "red+b:white")

// create a FAST closure function to avoid computation of ANSI code
phosphorize := emerald.ColorFunc("green+h:black")
msg = phosphorize("Bring back the 80s!")
msg2 := phospohorize("Look, I'm a CRT!")

// cache escape codes and build strings manually
lime := emerald.ColorCode("green+h:black")
reset := emerald.ColorCode("reset")

fmt.Println(lime, "Bring back the 80s!", reset)
```

Other examples

```go
Color(s, "red")            // red
Color(s, "red+d")          // red dim
Color(s, "red+b")          // red bold
Color(s, "red+B")          // red blinking
Color(s, "red+u")          // red underline
Color(s, "red+bh")         // red bold bright
Color(s, "red:white")      // red on white
Color(s, "red+b:white+h")  // red bold on white bright
Color(s, "red+B:white+h")  // red blink on white bright
Color(s, "off")            // turn off ansi codes
```

To view color combinations, from project directory in terminal.

```sh
go test
```

## Style format

```go
"foregroundColor+attributes:backgroundColor+attributes"
```

Colors

* black
* red
* green
* yellow
* blue
* magenta
* cyan
* white
* 0...255 (256 colors)

Foreground Attributes

* B = Blink
* b = bold
* h = high intensity (bright)
* d = dim
* i = inverse
* s = strikethrough
* u = underline

Background Attributes

* h = high intensity (bright)

## Constants

* emerald.Reset
* emerald.DefaultBG
* emerald.DefaultFG
* emerald.Black
* emerald.Red
* emerald.Green
* emerald.Yellow
* emerald.Blue
* emerald.Magenta
* emerald.Cyan
* emerald.White
* emerald.LightBlack
* emerald.LightRed
* emerald.LightGreen
* emerald.LightYellow
* emerald.LightBlue
* emerald.LightMagenta
* emerald.LightCyan
* emerald.LightWhite

## References

Wikipedia ANSI escape codes [Colors](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors)

General [tips and formatting](http://misc.flogisoft.com/bash/tip_colors_and_formatting)
