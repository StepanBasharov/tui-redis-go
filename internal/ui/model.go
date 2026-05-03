// Package ui implements the bubbletea TUI for redis-viewer.
package ui

import (
	"redis-viewer/internal/ui/forms"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// ViewerModel is the top-level bubbletea model for the TUI.
type ViewerModel struct {
	connections map[string]*Connection

	// state
	cursor              int
	setAddNewConnection bool
	setListOfConnection bool
	setWorkspace        bool
	currentConnection   *Connection
	Items               []string

	// new connection form
	connForm  forms.ConnectionForm
	connList  list.Model
	cmdInput  textinput.Model

	// terminal dimensions
	width  int
	height int
}

// InitialViewerModel returns a zero-state ViewerModel ready for tea.NewProgram.
func InitialViewerModel() ViewerModel {
	model := ViewerModel{}
	model.connForm = forms.NewConnectionForm(model.width)
	model.connections = make(map[string]*Connection)
	model.connList = list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	model.connList.Title = "\uE76D Active Connections \uE76D"
	model.connList.SetShowHelp(false)

	cmdInput := textinput.New()
	cmdInput.Placeholder = "type a Redis command..."
	cmdInput.Width = 60
	model.cmdInput = cmdInput

	return model
}

// Init satisfies tea.Model; no initial command is needed.
func (m ViewerModel) Init() tea.Cmd {
	return nil
}

// syncConnList rebuilds the list items from the current connections slice.
func (m *ViewerModel) syncConnList() {
	items := make([]list.Item, 0, len(m.connections))
	for addr := range m.connections {
		items = append(items, ConnectionItem{Addr: addr})
	}
	m.connList.SetItems(items)
}
