// Command tui-redis-go launches the interactive Redis TUI.
package main

import (
	"fmt"

	"github.com/StepanBasharov/tui-redis-go/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.InitialViewerModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
