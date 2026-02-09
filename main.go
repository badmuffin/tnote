package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	vault         = "~/.tnote/notes"
	inputBoxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("120"))
)

type model struct {
	newFileInput       textinput.Model
	isFileInputVisible bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+q":
			return m, tea.Quit
		case "ctrl+n":
			m.isFileInputVisible = true
			return m, nil
		case "enter":
			// save the file
		case "esc":
			m.isFileInputVisible = false
			return m, nil
		}
	}

	if m.isFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}

	return m, cmd

}

func (m model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).
		PaddingLeft(2).
		PaddingRight(2)

	welcome := style.Render("Thinking something... jote down with tnote")

	view := ""
	if m.isFileInputVisible {
		view = m.newFileInput.View()
	}

	keys := "Ctrl+N: new file  Ctrl+L: List  Ctrl+S: save  Esc: back  Ctrl+Q: quit"

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, keys)
}

func initializeMode() model {
	// creating a new directory to save the notes
	err := os.MkdirAll(vault, 0750)
	if err != nil {
		log.Fatal(err)
	}

	ti := textinput.New()
	ti.Placeholder = "What would you like to call it?"
	ti.Focus()
	ti.CharLimit = 128
	ti.Width = 30
	ti.Cursor.Style = inputBoxStyle
	ti.PromptStyle = inputBoxStyle
	ti.TextStyle = inputBoxStyle

	return model{
		newFileInput:       ti,
		isFileInputVisible: false,
	}
}

func main() {
	p := tea.NewProgram(initializeMode())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
