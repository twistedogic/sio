package sio

import (
	"fmt"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type message struct {
	At            time.Time
	Name, Content string
}

func newMessage(name, content string) message {
	return message{At: time.Now(), Name: name, Content: content}
}

func PrintMessage(name, content string) { fmt.Println(newMessage(name, content)) }

func (m message) String() string {
	color := pickColor(m.Name)
	name := lipgloss.NewStyle().Foreground(color).Bold(true).Render(m.At.Format(time.DateTime) + " " + m.Name + ":")
	content := m.Content
	if md, err := glamour.Render(m.Content, "auto"); err == nil {
		content = md
	}
	return name + "\n" + content + "\n---"

}
