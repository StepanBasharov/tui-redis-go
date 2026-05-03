package ui

import "github.com/StepanBasharov/tui-redis-go/internal/ui/ui-components"

// View renders the current screen based on model state.
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
		return ui_components.RenderListOfConnections(
			m.width,
			m.height,
			m.connList,
		)
	}

	if m.setWorkspace {
		currentHistory := m.currentConnection.History
		historyForRender := make([]ui_components.HistoryItem, 0, len(currentHistory))

		for _, historyItem := range currentHistory {
			historyForRender = append(historyForRender, ui_components.HistoryItem{
				Command:       historyItem.cmdHist,
				CommandOutput: historyItem.resHist,
			})
		}

		return ui_components.RenderWorkspace(
			m.width,
			m.height,
			m.currentConnection.Client.GetRedisAdd(),
			historyForRender,
			m.cmdInput,
		)
	}
	return ui_components.RenderBanner(m.width, m.height, len(m.connections))
}
