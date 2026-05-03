package ui

import (
	"log"

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

		case "ctrl+c":
			for _, conn := range m.connections {
				if err := conn.Client.Close(); err != nil {
					log.Printf("error closing connection: %s", err)
				}
			}
			return m, tea.Quit

		case "q":
			if !m.setAddNewConnection {
				return m, tea.Quit
			}

		case "ctrl+n":
			m.setAddNewConnection = true
			cmd := m.connForm.ResetFocus()
			return m, cmd

		case "ctrl+l":
			m.setListOfConnection = true
			return m, nil

		case "enter":
			if m.setAddNewConnection {
				hostForm := m.connForm.Host.Value()
				portForm := m.connForm.Port.Value()
				passwordForm := m.connForm.Password.Value()
				databaseFrom := m.connForm.DB.Value()

				connection, err := NewConnection(hostForm, portForm, passwordForm, databaseFrom)
				if err != nil {
					//
					return m, nil
				}

				m.connections = append(m.connections, connection)
				m.setAddNewConnection = false
				m.connForm.ResetForm()

				return m, nil
			}

		case "esc":
			m.setAddNewConnection = false
			m.setListOfConnection = false
			return m, nil

		case "tab":
			if m.setAddNewConnection {
				cmd := m.connForm.CycleFocus()
				return m, cmd
			}

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

	// пробрасываем сообщения в активный инпут формы
	if m.setAddNewConnection {
		cmd := m.connForm.Update(msg)
		return m, cmd
	}

	return m, nil
}
