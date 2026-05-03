package main

import (
	"fmt"
	"redis-viewer/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := ui.ViewerModel{
		Items: []string{"Nike", "Adidas", "New Balance"},
	}

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
