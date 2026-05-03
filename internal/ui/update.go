package ui

import (
	"context"
	"log"
	"strings"

	"github.com/StepanBasharov/tui-redis-go/internal/processor"

	tea "github.com/charmbracelet/bubbletea"
)

// Update handles all incoming messages for ViewerModel.
func (m ViewerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// terminal resize
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.connList.SetSize(msg.Width-4, msg.Height-4)

	// key presses
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			for _, conn := range m.connections {
				if err := conn.Client.Close(); err != nil {
					log.Printf("error closing connection: %s", err)
				}
			}
			return m, tea.Quit

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
					// TODO: popup with error
					return m, nil
				}

				addr := connection.Client.GetRedisAdd()

				if _, ok := m.connections[addr]; !ok {
					m.connections[addr] = connection
				}

				m.setAddNewConnection = false
				m.connForm.ResetForm()
				m.syncConnList()

				return m, nil
			}

			if m.setListOfConnection {
				if item, ok := m.connList.SelectedItem().(ConnectionItem); ok {
					addr := item.Addr

					if conn, ok := m.connections[addr]; ok {
						m.currentConnection = conn
						m.setWorkspace = true
						m.setListOfConnection = false
						m.cmdInput.Focus()
					} else {
						panic("connection not found")
					}
				}
			}

			if m.setWorkspace {
				command := strings.TrimSpace(m.cmdInput.Value())
				m.cmdInput.Reset()

				if command == "" {
					return m, nil
				}

				if command == "clear" {
					m.currentConnection.History = m.currentConnection.History[:0]

					return m, nil
				}

				prc := processor.NewProcessor(m.currentConnection.Client)

				out := prc.ProcessCmd(context.Background(), command)

				m.currentConnection.History = append(
					m.currentConnection.History,
					History{
						cmdHist: command,
						resHist: string(out.Data),
					},
				)

				return m, nil
			}

		case "esc":
			m.setAddNewConnection = false
			m.setListOfConnection = false
			m.setWorkspace = false
			m.currentConnection = nil
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

	// forward messages to the active form input
	if m.setAddNewConnection {
		cmd := m.connForm.Update(msg)
		return m, cmd
	}

	// forward messages to the connections list
	if m.setListOfConnection {
		var cmd tea.Cmd
		m.connList, cmd = m.connList.Update(msg)
		return m, cmd
	}

	if m.setWorkspace {
		var cmd tea.Cmd
		m.cmdInput, cmd = m.cmdInput.Update(msg)
		return m, cmd
	}

	return m, nil
}
