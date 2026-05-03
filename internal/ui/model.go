package ui

import (
	"redis-viewer/internal/redis"

	tea "github.com/charmbracelet/bubbletea"
)

type ViewerModel struct {
	connections []redis.AdapterRedisProtocol

	// состояние
	cursor              int
	setAddNewConnection bool
	Items               []string

	// системное
	width  int
	height int

	// статус
	loading bool
	err     error
}

func (m ViewerModel) Init() tea.Cmd {
	return nil
}
