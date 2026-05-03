package ui_components

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

// RenderListOfConnections renders the active connections list screen.
func RenderListOfConnections(
	width, height int,
	connList list.Model,
) string {
	box := borderStyle.Render(connList.View())
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}
