package ui_components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// HistoryItem holds a single command and its output for display.
type HistoryItem struct {
	Command       string
	CommandOutput string
}

// RenderWorkspace renders the Redis workspace: scrollable history on top,
// command input pinned to the bottom.
func RenderWorkspace(width, height int, addr string, history []HistoryItem, input textinput.Model) string {
	promptStyle := lipgloss.NewStyle().
		Foreground(redisLight).
		Bold(true)

	outputStyle := lipgloss.NewStyle().
		Foreground(hintColor)

	// build history lines
	title := titleStyle.Render("\uE76D "+addr+" \uE76D") + "\n\n"

	var histLines []string
	for _, h := range history {
		histLines = append(histLines, promptStyle.Render("> "+h.Command))
		histLines = append(histLines, outputStyle.Render("  "+h.CommandOutput))
		histLines = append(histLines, "")
	}

	inputBarHeight := 3
	titleHeight := 3 // title + 2 newlines
	historyHeight := height - inputBarHeight - titleHeight - 2

	if historyHeight < 1 {
		historyHeight = 1
	}

	histContent := strings.Join(histLines, "\n")

	lines := strings.Split(histContent, "\n")
	if len(lines) > historyHeight {
		lines = lines[len(lines)-historyHeight:]
	}
	histContent = strings.Join(lines, "\n")

	if len(lines) < historyHeight {
		padding := strings.Repeat("\n", historyHeight-len(lines))
		histContent = histContent + padding
	}

	histPanel := lipgloss.NewStyle().
		Width(width - 4).
		Height(historyHeight).
		Render(histContent)

	inputBar := fmt.Sprintf("%s %s",
		promptStyle.Render("redis>"),
		input.View(),
	)

	inputPanel := separatorStyle.Render(strings.Repeat("─", width-6)) + "\n" + inputBar

	content := lipgloss.JoinVertical(lipgloss.Left, title, histPanel, inputPanel)

	box := borderStyle.Width(width - 4).Render(content)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}
