package sio

import (
	"context"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"golang.org/x/term"
)

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80 // Fallback width
	}
	return width - 20
}

type Source interface {
	Name() string
	Response(context.Context, string) (string, error)
}

func prompt() (string, error) {
	width := getTerminalWidth()
	var prompt string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().Title("Enter prompt").ShowLineNumbers(true).Editor("vim").Value(&prompt),
		),
	)
	err := form.WithWidth(width).Run()
	return prompt, err
}

func response(ctx context.Context, prompt string, src Source) error {
	var res string
	var err error
	PrintMessage("You", prompt)
	spinner.New().Title("processing...").
		Action(func() {
			res, err = src.Response(ctx, prompt)
		}).Run()
	if err == nil {
		PrintMessage(src.Name(), res)
	}
	return err
}

func Start(ctx context.Context, name string, src Source) error {
	for {
		p, err := prompt()
		if err != nil {
			if err.Error() == "user aborted" {
				break
			}
			return err
		}
		if err := response(ctx, p, src); err != nil {
			return err
		}
	}
	return nil
}
