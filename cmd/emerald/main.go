package main

import (
	"fmt"
	"github.com/jcwillox/emerald"
	"github.com/mattn/go-colorable"
	"sort"
	"strconv"
)

func main() {
	printColors()
	print256Colors()
	printConstants()
}

func pad(s string, length int) string {
	for len(s) < length {
		s += " "
	}
	return s
}

func padColor(s string, styles []string) string {
	buffer := ""
	for _, style := range styles {
		buffer += emerald.Color(pad(s+style, 20), s+style)
	}
	return buffer
}

func printPlain() {
	emerald.SetColorState(false)
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
	for fg := range emerald.Colors {
		for _, bg := range bgColors {
			println(padColor(fg, []string{"" + bg, "+b" + bg, "+bh" + bg, "+u" + bg}))
			println(padColor(fg, []string{"+uh" + bg, "+B" + bg, "+Bb" + bg /* backgrounds */, "" + bg + "+h"}))
			println(padColor(fg, []string{"+b" + bg + "+h", "+bh" + bg + "+h", "+u" + bg + "+h", "+uh" + bg + "+h"}))
		}
	}
}

func printColors() {
	emerald.SetColorState(true)
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

	keys := []string{}
	for fg := range emerald.Colors {
		_, err := strconv.Atoi(fg)
		if err != nil {
			keys = append(keys, fg)
		}
	}
	sort.Strings(keys)

	for _, fg := range keys {
		for _, bg := range bgColors {
			fmt.Fprintln(stdout, padColor(fg, []string{"" + bg, "+b" + bg, "+bh" + bg, "+u" + bg}))
			fmt.Fprintln(stdout, padColor(fg, []string{"+uh" + bg, "+B" + bg, "+Bb" + bg /* backgrounds */, "" + bg + "+h", "+s" + bg}))
			fmt.Fprintln(stdout, padColor(fg, []string{"+b" + bg + "+h", "+bh" + bg + "+h", "+u" + bg + "+h", "+uh" + bg + "+h"}))
		}
	}
}

func print256Colors() {
	emerald.SetColorState(true)
	stdout := colorable.NewColorableStdout()

	bgColors := []string{""}
	for i := 0; i < 256; i++ {
		key := fmt.Sprintf(":%d", i)
		bgColors = append(bgColors, key)
	}

	keys := []string{}
	for fg := range emerald.Colors {
		n, err := strconv.Atoi(fg)
		if err == nil {
			keys = append(keys, fmt.Sprintf("%3d", n))
		}
	}
	sort.Strings(keys)

	for _, fg := range keys {
		for _, bg := range bgColors {
			fmt.Fprintln(stdout, padColor(fg, []string{"" + bg, "+b" + bg, "+u" + bg}))
			fmt.Fprintln(stdout, padColor(fg, []string{"+B" + bg, "+Bb" + bg, "+s" + bg}))
		}
	}
}

func printConstants() {
	stdout := colorable.NewColorableStdout()
	fmt.Fprintln(stdout, emerald.DefaultFG, "emerald.DefaultFG", emerald.Reset)
	fmt.Fprintln(stdout, emerald.DefaultBG, "emerald.DefaultBG", emerald.Reset)
	fmt.Fprintln(stdout, emerald.Black, "emerald.Black", emerald.Reset)
	fmt.Fprintln(stdout, emerald.Red, "emerald.Red", emerald.Reset)
	fmt.Fprintln(stdout, emerald.Green, "emerald.Green", emerald.Reset)
	fmt.Fprintln(stdout, emerald.Yellow, "emerald.Yellow", emerald.Reset)
	fmt.Fprintln(stdout, emerald.Blue, "emerald.Blue", emerald.Reset)
	fmt.Fprintln(stdout, emerald.Magenta, "emerald.Magenta", emerald.Reset)
	fmt.Fprintln(stdout, emerald.Cyan, "emerald.Cyan", emerald.Reset)
	fmt.Fprintln(stdout, emerald.White, "emerald.White", emerald.Reset)
	fmt.Fprintln(stdout, emerald.LightBlack, "emerald.LightBlack", emerald.Reset)
	fmt.Fprintln(stdout, emerald.LightRed, "emerald.LightRed", emerald.Reset)
	fmt.Fprintln(stdout, emerald.LightGreen, "emerald.LightGreen", emerald.Reset)
	fmt.Fprintln(stdout, emerald.LightYellow, "emerald.LightYellow", emerald.Reset)
	fmt.Fprintln(stdout, emerald.LightBlue, "emerald.LightBlue", emerald.Reset)
	fmt.Fprintln(stdout, emerald.LightMagenta, "emerald.LightMagenta", emerald.Reset)
	fmt.Fprintln(stdout, emerald.LightCyan, "emerald.LightCyan", emerald.Reset)
	fmt.Fprintln(stdout, emerald.LightWhite, "emerald.LightWhite", emerald.Reset)
}
