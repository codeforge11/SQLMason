package app

import (
	"fmt"
	"strings"
)

func (w *MainWindow) execute(_ bool) {
	w.messageText.Clear()
	w.statusLabel.SetText("")

	if w.db != nil {
		sqlCode := strings.TrimSpace(w.sqlEntry.ToPlainText())
		queries := strings.Split(sqlCode, ";")
		for _, query := range queries {
			query = strings.TrimSpace(query)
			if query == "" {
				continue
			}
			rows, err := w.db.Query(query)
			if err != nil {
				w.displayMessage(fmt.Sprintf("SQL execution error: %s", err))
				logError(err)
				continue
			}
			defer rows.Close()

			w.displayResults(rows)
		}
		w.statusLabel.SetText("SQL executed successfully")
	} else {
		w.statusLabel.SetText("Not connected to the database")
	}
}
