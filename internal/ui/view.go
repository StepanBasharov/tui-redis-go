package ui

import "redis-viewer/internal/ui/ui-components"

func (m ViewerModel) View() string {
	if m.setAddNewConnection {
		return ui_components.RenderAddNewConnection(
			m.width,
			m.height,
			m.connForm.Host,
			m.connForm.Port,
			m.connForm.Password,
			m.connForm.DB,
		)
	}
	if m.setListOfConnection {
		connections := make([]string, 0, len(m.connections))

		for _, conn := range m.connections {
			connections = append(connections, conn.Client.GetRedisAdd())
		}

		return ui_components.RenderListOfConnections(
			m.width,
			m.height,
			connections,
		)

	}
	return ui_components.RenderBanner(m.width, m.height, len(m.connections))
}
