package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	*widgets.QMainWindow

	db *sql.DB

	hostLabel          *widgets.QLabel
	portLabel          *widgets.QLabel
	userLabel          *widgets.QLabel
	passwordLabel      *widgets.QLabel
	hostInputField     *widgets.QLineEdit
	userInputField     *widgets.QLineEdit
	passwordInputField *widgets.QLineEdit
	portInputField     *widgets.QLineEdit
	connectButton      *widgets.QPushButton
	errorLabel         *widgets.QLabel
	sqlLabel           *widgets.QLabel
	sqlEntry           *widgets.QTextEdit
	executeButton      *widgets.QPushButton
	statusLabel        *widgets.QLabel
	resultLabel        *widgets.QLabel
	resultText         *widgets.QTextEdit
	messagesLabel      *widgets.QLabel
	messagesText       *widgets.QTextEdit
	exitButton         *widgets.QPushButton
}

func NewMainWindow() *MainWindow {
	window := &MainWindow{
		QMainWindow:        widgets.NewQMainWindow(nil, 0),
		db:                 &sql.DB{},
		hostLabel:          &widgets.QLabel{},
		portLabel:          &widgets.QLabel{},
		userLabel:          &widgets.QLabel{},
		passwordLabel:      &widgets.QLabel{},
		hostInputField:     &widgets.QLineEdit{},
		userInputField:     &widgets.QLineEdit{},
		passwordInputField: &widgets.QLineEdit{},
		portInputField:     &widgets.QLineEdit{},
		connectButton:      &widgets.QPushButton{},
		errorLabel:         &widgets.QLabel{},
		sqlLabel:           &widgets.QLabel{},
		sqlEntry:           &widgets.QTextEdit{},
		executeButton:      &widgets.QPushButton{},
		statusLabel:        &widgets.QLabel{},
		resultLabel:        &widgets.QLabel{},
		resultText:         &widgets.QTextEdit{},
		messagesLabel:      &widgets.QLabel{},
		messagesText:       &widgets.QTextEdit{},
		exitButton:         &widgets.QPushButton{},
	}

	Appversion := loadVersion()

	window.SetWindowTitle(fmt.Sprintf("SQLMason %s\n", Appversion))
	window.SetGeometry(core.NewQRect4(0, 0, 800, 800))

	window.SetWindowIcon(gui.NewQIcon5("Images/Logo.png"))
	window.SetFixedSize2(800, 800)

	window.initUI()
	window.hideElements()

	return window
}

func (w *MainWindow) initUI() {
	w.hostLabel = widgets.NewQLabel2("Host:", nil, 0)
	w.hostLabel.SetAlignment(core.Qt__AlignCenter)
	w.hostLabel.SetFont(gui.NewQFont2("Arial", 12, 1, false))

	w.portLabel = widgets.NewQLabel2("Port:", nil, 0)
	w.portLabel.SetAlignment(core.Qt__AlignCenter)
	w.portLabel.SetFont(gui.NewQFont2("Arial", 12, 1, false))

	w.userLabel = widgets.NewQLabel2("User:", nil, 0)
	w.userLabel.SetAlignment(core.Qt__AlignCenter)
	w.userLabel.SetFont(gui.NewQFont2("Arial", 12, 1, false))

	w.passwordLabel = widgets.NewQLabel2("Password:", nil, 0)
	w.passwordLabel.SetAlignment(core.Qt__AlignCenter)
	w.passwordLabel.SetFont(gui.NewQFont2("Arial", 12, 1, false))

	w.hostInputField = widgets.NewQLineEdit(nil)
	w.hostInputField.SetFont(gui.NewQFont2("Arial", 16, 1, false))
	w.hostInputField.SetPlaceholderText("localhost")

	w.userInputField = widgets.NewQLineEdit(nil)
	w.userInputField.SetFont(gui.NewQFont2("Arial", 16, 1, false))

	w.passwordInputField = widgets.NewQLineEdit(nil)
	w.passwordInputField.SetFont(gui.NewQFont2("Arial", 16, 1, false))

	w.portInputField = widgets.NewQLineEdit(nil)
	w.portInputField.SetFont(gui.NewQFont2("Arial", 16, 1, false))
	w.portInputField.SetPlaceholderText("3306")

	w.connectButton = widgets.NewQPushButton2("Connect to database", nil)
	w.connectButton.ConnectClicked(w.buttonClicked)
	w.connectButton.SetStyleSheet("background-color: #3a86ff; color: white; font-size: 18px;")

	w.errorLabel = widgets.NewQLabel(nil, 0)
	w.errorLabel.SetStyleSheet("color: red")

	w.sqlLabel = widgets.NewQLabel2("Enter SQL code:", nil, 0)
	w.sqlLabel.SetAlignment(core.Qt__AlignCenter)
	w.sqlLabel.SetFont(gui.NewQFont2("Arial", 16, 1, false))

	w.sqlEntry = widgets.NewQTextEdit(nil)
	w.sqlEntry.SetFont(gui.NewQFont2("Arial", 18, 1, false))

	w.exitButton = widgets.NewQPushButton2("Exit", nil)
	w.exitButton.ConnectClicked(w.exitDatabase)

	w.executeButton = widgets.NewQPushButton2("Execute SQL", nil)
	w.executeButton.ConnectClicked(w.executeSQL)
	w.executeButton.SetStyleSheet("background-color: #3a86ff; color: white; font-size: 14px;")

	w.statusLabel = widgets.NewQLabel(nil, 0)
	w.statusLabel.SetAlignment(core.Qt__AlignCenter)
	w.statusLabel.SetFont(gui.NewQFont2("Arial", 16, 1, false))

	w.resultLabel = widgets.NewQLabel2("Results:", nil, 0)
	w.resultLabel.SetAlignment(core.Qt__AlignCenter)
	w.resultLabel.SetFont(gui.NewQFont2("Arial", 16, 1, false))

	w.resultText = widgets.NewQTextEdit(nil)
	w.resultText.SetReadOnly(true)

	w.messagesLabel = widgets.NewQLabel2("Messages:", nil, 0)
	w.messagesLabel.SetAlignment(core.Qt__AlignCenter)
	w.messagesLabel.SetFont(gui.NewQFont2("Arial", 16, 1, false))

	w.messagesText = widgets.NewQTextEdit(nil)
	w.messagesText.SetReadOnly(true)

	layout := widgets.NewQVBoxLayout()
	layout.SetSpacing(10)
	layout.AddWidget(w.hostLabel, 0, 0)
	layout.AddWidget(w.hostInputField, 0, 0)
	layout.AddWidget(w.portLabel, 0, 0)
	layout.AddWidget(w.portInputField, 0, 0)
	layout.AddWidget(w.userLabel, 0, 0)
	layout.AddWidget(w.userInputField, 0, 0)
	layout.AddWidget(w.passwordLabel, 0, 0)
	layout.AddWidget(w.passwordInputField, 0, 0)
	layout.AddWidget(w.connectButton, 0, 0)
	layout.AddWidget(w.errorLabel, 0, 0)
	layout.AddWidget(w.sqlLabel, 0, 0)
	layout.AddWidget(w.sqlEntry, 0, 0)
	layout.AddWidget(w.executeButton, 0, 0)
	layout.AddWidget(w.statusLabel, 0, 0)
	layout.AddWidget(w.resultLabel, 0, 0)
	layout.AddWidget(w.resultText, 0, 0)
	layout.AddWidget(w.messagesLabel, 0, 0)
	layout.AddWidget(w.messagesText, 0, 0)
	layout.AddWidget(w.exitButton, 0, 0)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	w.SetCentralWidget(widget)
}

func (w *MainWindow) hideElements() {
	w.sqlLabel.Hide()
	w.sqlEntry.Hide()
	w.executeButton.Hide()
	w.resultLabel.Hide()
	w.resultText.Hide()
	w.messagesLabel.Hide()
	w.messagesText.Hide()
	w.exitButton.Hide()
	w.SetStyleSheet("background-color: light gray")
	w.SetFixedSize2(800, 400)
}

func (w *MainWindow) showElements() {
	w.sqlLabel.Show()
	w.sqlEntry.Show()
	w.executeButton.Show()
	w.resultLabel.Show()
	w.resultText.Show()
	w.messagesLabel.Show()
	w.messagesText.Show()
	w.exitButton.Show()
	w.SetFixedSize2(800, 800)
}

func (w *MainWindow) buttonClicked(_ bool) {
	host := w.hostInputField.Text()
	if host == "" {
		host = "127.0.0.1"
	}
	user := w.userInputField.Text()
	password := w.passwordInputField.Text()
	portText := w.portInputField.Text()
	port := 3306
	if portText != "" {
		var err error
		port, err = strconv.Atoi(portText)
		if err != nil {
			w.displayMessage(fmt.Sprintf("Invalid port: %s", err))
			logError(err)
			return
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, password, host, port)
	var err error
	w.db, err = sql.Open("mysql", dsn)
	if err != nil {
		w.displayMessage(fmt.Sprintf("Connection error: %s", err))
		w.errorLabel.SetText(err.Error())
		logError(err)
		return
	}

	err = w.db.Ping()
	if err != nil {
		w.displayMessage(fmt.Sprintf("Connection error: %s", err))
		w.errorLabel.SetText(err.Error())
		logError(err)
		return
	}

	w.hostLabel.Hide()
	w.userLabel.Hide()
	w.passwordLabel.Hide()
	w.portLabel.Hide()
	w.hostInputField.Hide()
	w.userInputField.Hide()
	w.passwordInputField.Hide()
	w.portInputField.Hide()
	w.connectButton.Hide()
	w.errorLabel.Hide()
	w.messagesText.Clear()

	w.showElements()
}

func (w *MainWindow) executeSQL(_ bool) {
	w.messagesText.Clear()
	w.statusLabel.SetText("")
	w.statusLabel.SetStyleSheet("")

	if w.db != nil {
		sqlCode := w.sqlEntry.ToPlainText()
		go func() {
			rows, err := w.db.Query(sqlCode)
			if err != nil {
				w.displayMessage(fmt.Sprintf("SQL execution error: %s", err))
				logError(err)
				return
			}
			defer rows.Close()

			w.displayResults(rows)
			w.statusLabel.SetText("SQL executed successfully")
			w.statusLabel.SetStyleSheet("color: green")
		}()
	} else {
		w.statusLabel.SetText("Not connected to the database")
		w.statusLabel.SetStyleSheet("color: red")
	}
}

func (w *MainWindow) displayResults(rows *sql.Rows) {
	w.resultText.Clear()

	columns, err := rows.Columns()
	if err != nil {
		w.displayMessage(fmt.Sprintf("Error getting columns: %s", err))
		logError(err)
		return
	}

	w.resultText.Append(strings.Join(columns, " | "))

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			w.displayMessage(fmt.Sprintf("Error scanning row: %s", err))
			logError(err)
			return
		}

		var rowData []string
		for _, col := range values {
			if col == nil {
				rowData = append(rowData, "NULL")
			} else {
				rowData = append(rowData, string(col))
			}
		}
		w.resultText.Append(strings.Join(rowData, " | "))
		w.resultText.Append(strings.Repeat("-", 40))
	}
}

func (w *MainWindow) displayMessage(message string) {
	w.messagesText.Append(message)
}

func (w *MainWindow) exitDatabase(_ bool) {
	if w.db != nil {
		w.db.Close()
		w.db = nil
	}
	w.hideElements()
	w.hostLabel.Show()
	w.userLabel.Show()
	w.passwordLabel.Show()
	w.portLabel.Show()
	w.hostInputField.Show()
	w.userInputField.Show()
	w.passwordInputField.Show()
	w.portInputField.Show()
	w.connectButton.Show()
	w.SetFixedSize2(800, 400)
}

func main() {
	app := widgets.NewQApplication(len([]string{}), []string{})
	mainWindow := NewMainWindow()
	mainWindow.Show()
	app.Exec()
}
