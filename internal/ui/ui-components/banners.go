package ui_components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// RenderBanner renders the startup banner centered within the given terminal dimensions.
func RenderBanner(width, height, activeConnections int) string {
	logo := logoStyle.Render(`
████████╗██╗   ██╗██╗██████╗ ███████╗██████╗ ██╗███████╗
╚══██╔══╝██║   ██║██║██╔══██╗██╔════╝██╔══██╗██║██╔════╝
   ██║   ██║   ██║██║██████╔╝█████╗  ██║  ██║██║███████╗
   ██║   ██║   ██║██║██╔══██╗██╔══╝  ██║  ██║██║╚════██║
   ██║   ╚██████╔╝██║██║  ██║███████╗██████╔╝██║███████║
   ╚═╝    ╚═════╝ ╚═╝╚═╝  ╚═╝╚══════╝╚═════╝ ╚═╝╚══════╝`)

	goText := goStyle.Render(`
 ██████╗  ██████╗ 
██╔════╝ ██╔═══██╗
██║  ███╗██║   ██║
██║   ██║██║   ██║
╚██████╔╝╚██████╔╝
 ╚═════╝  ╚═════╝
`)

	banner := lipgloss.JoinHorizontal(lipgloss.Top, logo, goText)

	hints := "\n" +
		hintKeyStyle.
			Render("  Ctrl-n") + hintTextStyle.
		Render("  add new connection") + "\n" +
		hintKeyStyle.
			Render("  Ctrl-l") + hintTextStyle.
		Render(fmt.Sprintf("  list of connections (active connections: %d)", activeConnections)) + "\n" +
		hintKeyStyle.
			Render("  Ctrl-c") + hintTextStyle.
		Render("  quit") + "\n"

	content := banner + "\n" + hints

	box := borderStyle.Render(content)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}
