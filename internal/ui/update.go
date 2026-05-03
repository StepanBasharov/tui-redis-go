package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m ViewerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// resize терминала
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	// клавиши
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "ctrl+a":
			m.setAddNewConnection = true
			return m, nil

		case "ctrl+l":
			fmt.Println("Список подключений")
			return m, nil

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.Items)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}
