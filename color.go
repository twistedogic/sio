package sio

import "github.com/charmbracelet/lipgloss"

// https://coolors.co/e4fde1-b7e4b5-8acb88-77a785-648381-5e6d71-575761-ab8b54-ffbf46
var colorPalete = []lipgloss.Color{
	lipgloss.Color("#E4FDE1"),
	lipgloss.Color("#B7E4B5"),
	lipgloss.Color("#8ACB88"),
	lipgloss.Color("#77A785"),
	lipgloss.Color("#648381"),
	lipgloss.Color("#5E6D71"),
	lipgloss.Color("#575761"),
	lipgloss.Color("#AB8B54"),
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
