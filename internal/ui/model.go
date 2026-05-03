// Package ui implements the bubbletea TUI for redis-viewer.
package ui

import (
	"redis-viewer/internal/ui/forms"

	tea "github.com/charmbracelet/bubbletea"
)

// ViewerModel is the top-level bubbletea model for the TUI.
type ViewerModel struct {
	connections []*Connection

	// state
	cursor              int
	setAddNewConnection bool
	setListOfConnection bool
	Items               []string

	// new connection form
	connForm forms.ConnectionForm

	// terminal dimensions
	width  int
	height int

	// status
	loading bool
	err     error
}

// InitialViewerModel returns a zero-state ViewerModel ready for tea.NewProgram.
func InitialViewerModel() ViewerModel {
	model := ViewerModel{}
	model.connForm = forms.NewConnectionForm(model.width)
	model.connections = make([]*Connection, 0)
	return model
}

// Init satisfies tea.Model; no initial command is needed.
func (m ViewerModel) Init() tea.Cmd {
	return nil
}
