package sio

import "github.com/charmbracelet/lipgloss"

var colorPalete = []lipgloss.Color{
	lipgloss.Color("#E4FDE1"),
	lipgloss.Color("#8ACB88"),
	lipgloss.Color("#648381"),
	lipgloss.Color("#575761"),
	lipgloss.Color("#FFBF46"),
}

func hashmod(s string, d int) int {
	sum := 0
	for _, r := range s {
		sum += int(r)
	}
	return (sum * sum) % d
}

func pickColor(s string) lipgloss.Color {
	return colorPalete[hashmod(s, len(colorPalete))]
}
