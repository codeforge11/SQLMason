package app

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/core"
)

func (w *MainWindow) execute(_ bool) {
	w.messageText.Clear()
	w.statusLabel.SetText("")

	if w.db == nil {
		w.statusLabel.SetText("Not connected to the database")
		return
	}

	sqlCode := strings.TrimSpace(w.sqlEntry.ToPlainText())
	queries := strings.Split(sqlCode, ";")
	for _, query := range queries {
		query = strings.TrimSpace(query)

		rows, err := w.db.Query(query)
		if err != nil {
			w.displayMessage(fmt.Sprintf("SQL execution error: %s", err))
			logError(err)
			continue
		}

		if query != "" {
			w.statusLabel.SetText("SQL executed successfully")
			defer rows.Close()
			w.displayResults(rows)
		}
	}

	timer := core.NewQTimer(nil)
	timer.SetSingleShot(true)
	timer.ConnectTimeout(w.clearStatusLabel)
	timer.Start(5000)
}

func (w *MainWindow) clearStatusLabel() {
	w.statusLabel.SetText("")
}
