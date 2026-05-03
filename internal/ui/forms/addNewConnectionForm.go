// Package forms provides bubbletea form structs for user input screens.
package forms

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// ConnectionForm groups the text inputs for the "add new connection" screen
// and tracks which input is currently focused.
type ConnectionForm struct {
	Host     textinput.Model
	Port     textinput.Model
	Password textinput.Model
	DB       textinput.Model
	focused  int
}

// inputs returns pointers to all form fields in tab-order.
func (f *ConnectionForm) inputs() []*textinput.Model {
	return []*textinput.Model{&f.Host, &f.Port, &f.Password, &f.DB}
}

// CycleFocus moves focus to the next input (wraps around).
func (f *ConnectionForm) CycleFocus() tea.Cmd {
	all := f.inputs()
	all[f.focused].Blur()
	f.focused = (f.focused + 1) % len(all)
	return all[f.focused].Focus()
}

// ResetFocus blurs all inputs and focuses the first one.
func (f *ConnectionForm) ResetFocus() tea.Cmd {
	for _, inp := range f.inputs() {
		inp.Blur()
	}
	f.focused = 0
	return f.Host.Focus()
}

// ResetForm clears all input field values.
func (f *ConnectionForm) ResetForm() {
	f.Host.Reset()
	f.Port.Reset()
	f.DB.Reset()
	f.Password.Reset()
}

// Update forwards a tea.Msg to the currently focused input.
func (f *ConnectionForm) Update(msg tea.Msg) tea.Cmd {
	all := f.inputs()
	var cmd tea.Cmd
	*all[f.focused], cmd = all[f.focused].Update(msg)
	return cmd
}

// NewConnectionForm creates a ConnectionForm with default placeholders.
func NewConnectionForm(width int) ConnectionForm {
	host := textinput.New()
	host.Placeholder = "Redis host"
	host.Focus()
	host.Width = 60

	port := textinput.New()
	port.Placeholder = "Redis port"
	port.CharLimit = 5
	port.Width = 60

	password := textinput.New()
	password.Placeholder = "Redis password"
	password.EchoMode = textinput.EchoPassword
	password.EchoCharacter = '•'
	password.Width = 60

	db := textinput.New()
	db.Placeholder = "Redis database"
	db.Width = 60

	return ConnectionForm{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       db,
	}
}
