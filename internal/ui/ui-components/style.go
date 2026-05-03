package ui_components

import "github.com/charmbracelet/lipgloss"

var (
	// Redis brand red + lighter accent for the gradient effect.
	redisRed   = lipgloss.Color("#D82C20")
	redisLight = lipgloss.Color("#00ADD8")
	hintColor  = lipgloss.Color("#888888")
	accentDim  = lipgloss.Color("#A02020")

	logoStyle = lipgloss.NewStyle().
			Foreground(redisRed).
			Bold(true)

	goStyle = lipgloss.NewStyle().
		Foreground(redisLight).
		Bold(true)

	hintKeyStyle = lipgloss.NewStyle().
			Foreground(redisLight).
			Bold(true)

	hintTextStyle = lipgloss.NewStyle().
			Foreground(hintColor)

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(accentDim).
			Padding(0, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(redisRed).
			Bold(true).
			MarginBottom(1)

	separatorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#444444"))
)
