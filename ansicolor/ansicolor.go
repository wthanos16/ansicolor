package ansicolor

import "fmt"

const END = "\033[0m"

func ColorString(color string, input string) string {
	var output string
	colors := map[string]string{
		"black":  "\033[30m",
		"blue":   "\033[34m",
		"grey":   "\033[1;30m",
		"red":    "\033[1;31m",
		"green":  "\033[1;32m",
		"yellow": "\033[1;33m",
		"purple": "\033[1;35m",
		"cyan":   "\033[1;36m",
		"white":  "\033[1;37m",
	}
	for i, v := range colors {
		if i == color {
			output = fmt.Sprintf("%s%s%s", v, input, END)
		}
	}
	return output
}

func StyleString(style string, input string) string {
	var output string
	styles := map[string]string{
		"bold":      "\033[1m",
		"italic":    "\033[3m",
		"underline": "\033[4m",
		"blinking":  "\033[5m",
		"crossout":  "\033[9m",
	}
	for i, v := range styles {
		if i == style {
			output = fmt.Sprintf("%s%s%s", v, input, END)
		}
	}
	return output
}

func HighlightString(color string, input string) string {
	var output string
	colors := map[string]string{
		"red":    "\033[30;41m",
		"green":  "\033[30;42m",
		"yellow": "\033[30;43m",
		"blue":   "\033[30;44m",
		"purple": "\033[30;45m",
		"cyan":   "\033[30;46m",
		"white":  "\033[30;47m",
	}
	for i, v := range colors {
		if i == color {
			output = fmt.Sprintf("%s%s%s", v, input, END)
		}
	}
	return output
}
