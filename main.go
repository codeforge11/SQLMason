package main

import (
	appdata "SQLMason/files"
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"

	_ "github.com/go-sql-driver/mysql"
)

type MainWindow struct {
	*widgets.QMainWindow

	db *sql.DB

	titleLabel    *widgets.QLabel
	hostLabel     *widgets.QLabel
	portLabel     *widgets.QLabel
	userLabel     *widgets.QLabel
	passwordLabel *widgets.QLabel
	errorLabel    *widgets.QLabel
	sqlLabel      *widgets.QLabel
	statusLabel   *widgets.QLabel
	resultLabel   *widgets.QLabel
	messagesLabel *widgets.QLabel

	hostInputField     *widgets.QLineEdit
	userInputField     *widgets.QLineEdit
	passwordInputField *widgets.QLineEdit
	portInputField     *widgets.QLineEdit

	sqlEntry     *widgets.QTextEdit
	resultText   *widgets.QTextEdit
	messagesText *widgets.QTextEdit

	executeButton     *widgets.QPushButton
	connectButton     *widgets.QPushButton
	connecttodbButton *widgets.QPushButton
	exitButton        *widgets.QPushButton
	returnButton      *widgets.QPushButton
}

func NewMainWindow() *MainWindow {
	window := &MainWindow{
		QMainWindow: widgets.NewQMainWindow(nil, 0),

		db: &sql.DB{},

		hostLabel:     widgets.NewQLabel(nil, 0),
		portLabel:     widgets.NewQLabel(nil, 0),
		userLabel:     widgets.NewQLabel(nil, 0),
		passwordLabel: widgets.NewQLabel(nil, 0),
		statusLabel:   widgets.NewQLabel(nil, 0),
		resultLabel:   widgets.NewQLabel(nil, 0),
		errorLabel:    widgets.NewQLabel(nil, 0),
		sqlLabel:      widgets.NewQLabel(nil, 0),
		messagesLabel: widgets.NewQLabel(nil, 0),

		titleLabel: widgets.NewQLabel2(fmt.Sprintf("SQLMason %s", appdata.Version), nil, 0),

		hostInputField:     widgets.NewQLineEdit(nil),
		userInputField:     widgets.NewQLineEdit(nil),
		passwordInputField: widgets.NewQLineEdit(nil),
		portInputField:     widgets.NewQLineEdit(nil),

		sqlEntry:     widgets.NewQTextEdit(nil),
		resultText:   widgets.NewQTextEdit(nil),
		messagesText: widgets.NewQTextEdit(nil),

		connecttodbButton: widgets.NewQPushButton(nil),
		connectButton:     widgets.NewQPushButton(nil),
		executeButton:     widgets.NewQPushButton(nil),
		exitButton:        widgets.NewQPushButton(nil),
		returnButton:      widgets.NewQPushButton(nil),
	}

	window.SetWindowTitle(fmt.Sprintf("SQLMason %s", appdata.Version))
	window.SetGeometry(core.NewQRect4(0, 0, 800, 800))
	window.SetWindowIcon(gui.NewQIcon5("source/Images/Logo.svg"))
	window.SetFixedSize2(800, 800)

	window.initUI()
	window.firstrun()

	return window
}

func (w *MainWindow) initUI() {

	w.titleLabel.SetObjectName("titleLabel")

	w.hostLabel = widgets.NewQLabel2("Host:", nil, 0)
	w.hostLabel.SetAlignment(core.Qt__AlignCenter)
	w.hostLabel.SetFont(gui.NewQFont2("Inter", 12, 1, false))

	w.portLabel = widgets.NewQLabel2("Port:", nil, 0)
	w.portLabel.SetAlignment(core.Qt__AlignCenter)
	w.portLabel.SetFont(gui.NewQFont2("Inter", 12, 1, false))

	w.userLabel = widgets.NewQLabel2("User:", nil, 0)
	w.userLabel.SetAlignment(core.Qt__AlignCenter)
	w.userLabel.SetFont(gui.NewQFont2("Inter", 12, 1, false))

	w.passwordLabel = widgets.NewQLabel2("Password:", nil, 0)
	w.passwordLabel.SetAlignment(core.Qt__AlignCenter)
	w.passwordLabel.SetFont(gui.NewQFont2("Inter", 12, 1, false))

	w.statusLabel = widgets.NewQLabel(nil, 0)
	w.statusLabel.SetAlignment(core.Qt__AlignCenter)
	w.statusLabel.SetFont(gui.NewQFont2("Inter", 16, 1, false))

	w.resultLabel = widgets.NewQLabel2("Results:", nil, 0)
	w.resultLabel.SetAlignment(core.Qt__AlignCenter)
	w.resultLabel.SetFont(gui.NewQFont2("Inter", 16, 1, false))

	w.errorLabel = widgets.NewQLabel(nil, 0)
	w.errorLabel.SetObjectName("errorLabel")

	w.sqlLabel = widgets.NewQLabel2("Enter SQL code:", nil, 0)
	w.sqlLabel.SetAlignment(core.Qt__AlignCenter)
	w.sqlLabel.SetFont(gui.NewQFont2("Inter", 18, 1, false))

	w.messagesLabel = widgets.NewQLabel2("Messages:", nil, 0)
	w.messagesLabel.SetAlignment(core.Qt__AlignCenter)
	w.messagesLabel.SetFont(gui.NewQFont2("Inter", 16, 1, false))

	w.hostInputField.SetFont(gui.NewQFont2("Inter", 16, 1, false))
	w.hostInputField.SetPlaceholderText("localhost")

	w.userInputField.SetFont(gui.NewQFont2("Inter", 16, 1, false))

	w.passwordInputField.SetFont(gui.NewQFont2("Inter", 16, 1, false))

	w.portInputField.SetFont(gui.NewQFont2("Inter", 16, 1, false))
	w.portInputField.SetPlaceholderText("3306")

	w.connectButton = widgets.NewQPushButton2("Connect to database", nil)
	w.connectButton.ConnectClicked(w.buttonClicked)
	w.connectButton.SetObjectName("connectButton")

	w.exitButton = widgets.NewQPushButton2("Back", nil)
	w.exitButton.ConnectClicked(w.exitDatabase)

	w.executeButton = widgets.NewQPushButton2("Execute SQL", nil)
	w.executeButton.ConnectClicked(w.executeSQL)
	w.executeButton.SetObjectName("executeButton")

	w.connecttodbButton = widgets.NewQPushButton2("Connect to database", nil)
	w.connecttodbButton.ConnectClicked(w.buttonClicked2)
	w.connecttodbButton.SetObjectName("connecttodbButton")

	w.returnButton = widgets.NewQPushButton2("Return", nil)
	w.returnButton.ConnectClicked(w.returnclicket)
	w.returnButton.SetObjectName("returnButton")

	w.resultText = widgets.NewQTextEdit(nil)
	w.resultText.SetReadOnly(true)

	w.messagesText = widgets.NewQTextEdit(nil)
	w.messagesText.SetReadOnly(true)

	w.sqlEntry = widgets.NewQTextEdit(nil)
	w.sqlEntry.SetFont(gui.NewQFont2("Inter", 18, 1, false))

	layout := widgets.NewQVBoxLayout()
	layout.SetSpacing(10)

	layout.AddWidget(w.titleLabel, 0, core.Qt__AlignTop|core.Qt__AlignCenter)
	layout.AddWidget(w.connecttodbButton, 0, core.Qt__AlignCenter)
	layout.AddWidget(w.hostLabel, 0, 0)
	layout.AddWidget(w.hostInputField, 0, 0)
	layout.AddWidget(w.portLabel, 0, 0)
	layout.AddWidget(w.portInputField, 0, 0)
	layout.AddWidget(w.userLabel, 0, 0)
	layout.AddWidget(w.userInputField, 0, 0)
	layout.AddWidget(w.passwordLabel, 0, 0)
	layout.AddWidget(w.passwordInputField, 0, 0)

	connectLayout := widgets.NewQHBoxLayout()
	connectLayout.AddWidget(w.connectButton, 0, 0)
	connectLayout.AddWidget(w.returnButton, 0, 0)

	layout.AddLayout(connectLayout, 0)

	layout.AddWidget(w.sqlLabel, 0, 0)
	layout.AddWidget(w.sqlEntry, 0, 0)
	layout.AddWidget(w.executeButton, 0, 0)
	layout.AddWidget(w.statusLabel, 0, 0)
	layout.AddWidget(w.resultLabel, 0, 0)
	layout.AddWidget(w.resultText, 0, 0)
	layout.AddWidget(w.messagesLabel, 0, 0)
	layout.AddWidget(w.messagesText, 0, 0)
	layout.AddWidget(w.exitButton, 0, 0)
	layout.AddWidget(w.errorLabel, 0, 0)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	w.SetCentralWidget(widget)
}

func (w *MainWindow) firstrun() {

	color.HiGreen("App is running")

	w.titleLabel.Show()
	w.hostLabel.Hide()
	w.portLabel.Hide()
	w.userLabel.Hide()
	w.passwordLabel.Hide()
	w.hostInputField.Hide()
	w.userInputField.Hide()
	w.passwordInputField.Hide()
	w.portInputField.Hide()
	w.connectButton.Hide()
	w.sqlLabel.Hide()
	w.sqlEntry.Hide()
	w.executeButton.Hide()
	w.resultLabel.Hide()
	w.resultText.Hide()
	w.messagesLabel.Hide()
	w.messagesText.Hide()
	w.exitButton.Hide()
	w.returnButton.Hide()
	w.SetStyleSheet("background-color: light gray")
	w.SetFixedSize2(700, 400)
}

func (w *MainWindow) buttonClicked2(checked bool) {
	w.connecttodbButton.Hide()
	w.titleLabel.Hide()

	w.hostLabel.Show()
	w.hostInputField.Show()
	w.portLabel.Show()
	w.portInputField.Show()
	w.userLabel.Show()
	w.userInputField.Show()
	w.passwordLabel.Show()
	w.passwordInputField.Show()
	w.connectButton.Show()
	w.returnButton.Show()
	w.errorLabel.Show()

	w.SetFixedSize2(800, 440)
}

func (w *MainWindow) showElements() {

	w.hostLabel.Hide()
	w.hostInputField.Hide()
	w.portLabel.Hide()
	w.portInputField.Hide()
	w.userLabel.Hide()
	w.userInputField.Hide()
	w.passwordLabel.Hide()
	w.passwordInputField.Hide()
	w.connectButton.Hide()
	w.errorLabel.Hide()

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

	w.returnButton.Hide()
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
	w.errorLabel.SetText(message)
	w.errorLabel.Show()
	timer := core.NewQTimer(nil)
	timer.SetSingleShot(true)
	timer.ConnectTimeout(w.clearErrorLabel)
	timer.Start(5000)
}

func (w *MainWindow) clearErrorLabel() {
	w.errorLabel.SetText("")
}

func (w *MainWindow) exitDatabase(_ bool) {
	if w.db != nil {
		w.db.Close()
		w.db = nil
	}

	w.sqlLabel.Hide()
	w.sqlEntry.Hide()
	w.executeButton.Hide()
	w.resultLabel.Hide()
	w.resultText.Hide()
	w.messagesLabel.Hide()
	w.messagesText.Hide()
	w.exitButton.Hide()
	w.statusLabel.Hide()

	w.returnButton.Show()
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

func (w *MainWindow) returnclicket(_ bool) {

	w.errorLabel.SetText("")

	w.connecttodbButton.Show()
	w.titleLabel.Show()

	w.errorLabel.Hide()
	w.hostLabel.Hide()
	w.portLabel.Hide()
	w.userLabel.Hide()
	w.passwordLabel.Hide()
	w.hostInputField.Hide()
	w.userInputField.Hide()
	w.passwordInputField.Hide()
	w.portInputField.Hide()
	w.connectButton.Hide()
	w.sqlLabel.Hide()
	w.sqlEntry.Hide()
	w.executeButton.Hide()
	w.resultLabel.Hide()
	w.resultText.Hide()
	w.messagesLabel.Hide()
	w.messagesText.Hide()
	w.exitButton.Hide()
	w.returnButton.Hide()
	w.SetStyleSheet("background-color: light gray")
	w.SetFixedSize2(700, 400)
}

func main() {

	if runtime.GOOS == "linux" {
		os.Setenv("QT_QPA_PLATFORM", "xcb") // Sets x11 for linux
	}

	app := widgets.NewQApplication(len([]string{}), []string{})
	appdata.Style(app)

	if core.QCoreApplication_Instance() == nil {
		color.Red("Failed to initialize QCoreApplication")
	}

	mainWindow := NewMainWindow()
	mainWindow.Show()
	app.Exec()

}
