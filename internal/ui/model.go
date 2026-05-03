package ui

import (
	"redis-viewer/internal/ui/forms"

	tea "github.com/charmbracelet/bubbletea"
)

type ViewerModel struct {
	connections []*Connection

	// состояние
	cursor              int
	setAddNewConnection bool
	setListOfConnection bool
	Items               []string

	// форма нового подключения
	connForm forms.ConnectionForm

	// системное
	width  int
	height int

	// статус
	loading bool
	err     error
}

func InitialViewerModel() ViewerModel {
	model := ViewerModel{}
	model.connForm = forms.NewConnectionForm(model.width)
	model.connections = make([]*Connection, 0)
	return model
}

func (m ViewerModel) Init() tea.Cmd {
	return nil
}
