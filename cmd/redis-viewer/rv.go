// Command redis-viewer launches the interactive Redis TUI.
package main

import (
	"fmt"

	"redis-viewer/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.InitialViewerModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
