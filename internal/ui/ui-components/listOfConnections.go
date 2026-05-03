package ui_components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// RenderListOfConnections renders the active connections list screen.
func RenderListOfConnections(
	width, height int,
	conns []string,
) string {
	content := titleStyle.Render(fmt.Sprintf("\uE76D Active Connections %d \uE76D\n\n", len(conns)))

	for _, conn := range conns {
		content += fmt.Sprintf("\uE76D> %s\n\n", conn)
	}

	box := borderStyle.Render(content)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}
