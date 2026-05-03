package ui

import "redis-viewer/internal/ui/ui-components"

func (m ViewerModel) View() string {
	if m.setAddNewConnection {
		return ui_components.RenderAddNewConnection(m.width, m.height)
	}
	return ui_components.RenderBanner(m.width, m.height, len(m.connections))
}
