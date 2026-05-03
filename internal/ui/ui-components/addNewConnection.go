package ui_components

import "github.com/charmbracelet/lipgloss"

func RenderAddNewConnection(width, height int) string {
	return lipgloss.JoinVertical(
		lipgloss.Left, "Text")
}
