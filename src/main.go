package app

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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
	versionLabel  *widgets.QLabel
	creatorLabel  *widgets.QLabel

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
	exitAppButton     *widgets.QPushButton

	dbTypeComboBox *widgets.QComboBox
}

func newMainWindow() *MainWindow {
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
		titleLabel:    widgets.NewQLabel(nil, 0),
		versionLabel:  widgets.NewQLabel(nil, 0),
		creatorLabel:  widgets.NewQLabel(nil, 0),

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
		exitAppButton:     widgets.NewQPushButton(nil),
	}

	window.SetWindowTitle("SQLMason")
	window.SetGeometry(core.NewQRect4(0, 30, 800, 800))
	window.SetWindowIcon(gui.NewQIcon5("src/public/logomark.svg"))
	window.SetFixedSize2(800, 800)

	pixmap := gui.NewQPixmap3("src/public/logotype.svg", "", core.Qt__AutoColor)
	pixmap = pixmap.Scaled2(300, 200, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)

	window.titleLabel.SetPixmap(pixmap)

	window.initUI()
	window.firstRun()

	return window
}

func (w *MainWindow) initUI() {
	var codeFont = gui.QFontDatabase_SystemFont(gui.QFontDatabase__FixedFont)

	w.hostLabel = widgets.NewQLabel2("Host:", nil, 0)
	w.hostLabel.SetAlignment(core.Qt__AlignCenter)

	w.portLabel = widgets.NewQLabel2("Port:", nil, 0)
	w.portLabel.SetAlignment(core.Qt__AlignCenter)

	w.userLabel = widgets.NewQLabel2("User:", nil, 0)
	w.userLabel.SetAlignment(core.Qt__AlignCenter)

	w.passwordLabel = widgets.NewQLabel2("Password:", nil, 0)
	w.passwordLabel.SetAlignment(core.Qt__AlignCenter)

	w.statusLabel = widgets.NewQLabel(nil, 0)
	w.statusLabel.SetAlignment(core.Qt__AlignCenter)
	w.statusLabel.SetStyleSheet("QLabel { color : green; }")

	w.resultLabel = widgets.NewQLabel2("Results:", nil, 0)
	w.resultLabel.SetAlignment(core.Qt__AlignCenter)

	w.errorLabel = widgets.NewQLabel(nil, 0)
	w.errorLabel.SetStyleSheet("QLabel { color : red; }")
	w.errorLabel.SetFont(codeFont)

	w.sqlLabel = widgets.NewQLabel2("Enter SQL code:", nil, 0)
	w.sqlLabel.SetAlignment(core.Qt__AlignCenter)

	w.messagesLabel = widgets.NewQLabel2("Messages:", nil, 0)
	w.messagesLabel.SetAlignment(core.Qt__AlignCenter)

	w.versionLabel = widgets.NewQLabel2(Version, nil, 0)

	w.creatorLabel = widgets.NewQLabel2("Created by: @codeforge11", nil, 0)

	w.hostInputField.SetPlaceholderText("localhost")

	w.portInputField.SetPlaceholderText("3306")
	portValidator := gui.NewQIntValidator(nil)
	w.portInputField.SetValidator(portValidator)

	w.passwordInputField.SetEchoMode(widgets.QLineEdit__PasswordEchoOnEdit)

	w.connectButton = widgets.NewQPushButton2("Connect to database", nil)
	w.connectButton.ConnectClicked(w.buttonClicked)

	w.exitButton = widgets.NewQPushButton2("Back", nil)
	w.exitButton.ConnectClicked(w.exitDatabase)

	w.executeButton = widgets.NewQPushButton2("Execute SQL", nil)
	w.executeButton.ConnectClicked(w.executeSQL)

	w.connecttodbButton = widgets.NewQPushButton2("Connect to database", nil)
	w.connecttodbButton.ConnectClicked(w.buttonClicked2)

	w.returnButton = widgets.NewQPushButton2("Return", nil)
	w.returnButton.ConnectClicked(w.returnClicked)

	w.exitAppButton = widgets.NewQPushButton2("Exit", nil)
	w.exitAppButton.ConnectClicked(w.exitApp)

	w.resultText = widgets.NewQTextEdit(nil)
	w.resultText.SetReadOnly(true)
	w.resultText.SetFont(codeFont)

	w.messagesText = widgets.NewQTextEdit(nil)
	w.messagesText.SetReadOnly(true)
	w.messagesText.SetFont(codeFont)

	w.sqlEntry = widgets.NewQTextEdit(nil)
	w.sqlEntry.SetFont(codeFont)

	w.dbTypeComboBox = widgets.NewQComboBox(nil)
	w.dbTypeComboBox.AddItems([]string{"MySQL/MariaDB", "PostgreSQL", "Microsoft SQL Server"})

	layout := widgets.NewQVBoxLayout()
	layout.SetSpacing(10)

	layout.AddWidget(w.titleLabel, 0, core.Qt__AlignTop|core.Qt__AlignCenter)
	layout.AddWidget(w.connecttodbButton, 0, core.Qt__AlignCenter)
	layout.AddWidget(w.exitAppButton, 0, core.Qt__AlignCenter)
	layout.AddWidget(w.hostLabel, 0, 0)
	layout.AddWidget(w.hostInputField, 0, 0)
	layout.AddWidget(w.portLabel, 0, 0)
	layout.AddWidget(w.portInputField, 0, 0)
	layout.AddWidget(w.userLabel, 0, 0)
	layout.AddWidget(w.userInputField, 0, 0)
	layout.AddWidget(w.passwordLabel, 0, 0)
	layout.AddWidget(w.passwordInputField, 0, 0)
	layout.AddWidget(w.dbTypeComboBox, 0, 0)

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

	creatorandversionlayout := widgets.NewQHBoxLayout()
	creatorandversionlayout.AddWidget(w.creatorLabel, 0, core.Qt__AlignLeft|core.Qt__AlignBottom)
	creatorandversionlayout.AddWidget(w.versionLabel, 0, core.Qt__AlignRight|core.Qt__AlignBottom)

	layout.AddLayout(creatorandversionlayout, 0)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	w.SetCentralWidget(widget)
}

func (w *MainWindow) firstRun() {

	color.HiGreen("App is running")

	w.versionLabel.Show()
	w.titleLabel.Show()
	w.exitAppButton.Show()

	w.dbTypeComboBox.Hide()
	w.hostLabel.Hide()
	w.portLabel.Hide()
	w.userLabel.Hide()
	w.passwordLabel.Hide()
	w.errorLabel.Hide()
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
	w.SetFixedSize2(700, 400)
}

func (w *MainWindow) buttonClicked2(checked bool) {
	w.connecttodbButton.Hide()
	w.titleLabel.Hide()
	w.versionLabel.Hide()
	w.exitAppButton.Hide()
	w.creatorLabel.Hide()

	w.dbTypeComboBox.Show()
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

func (w *MainWindow) showElementsafterConnect() {

	w.dbTypeComboBox.Hide()
	w.hostLabel.Hide()
	w.errorLabel.Hide()
	w.returnButton.Hide()
	w.userLabel.Hide()
	w.passwordLabel.Hide()
	w.portLabel.Hide()
	w.hostInputField.Hide()
	w.userInputField.Hide()
	w.passwordInputField.Hide()
	w.portInputField.Hide()
	w.connectButton.Hide()

	w.sqlLabel.Show()
	w.sqlEntry.Show()
	w.executeButton.Show()
	w.resultLabel.Show()
	w.resultText.Show()
	w.messagesLabel.Show()
	w.messagesText.Show()
	w.exitButton.Show()

	w.messagesText.Clear()
	w.SetFixedSize2(800, 800)
}

func (w *MainWindow) buttonClicked(_ bool) { //connect to db
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

	dbType := w.dbTypeComboBox.CurrentText()
	var dsn string
	var err error

	switch dbType {
	case "MySQL/MariaDB":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, password, host, port)
		w.db, err = sql.Open("mysql", dsn)
	case "PostgreSQL":
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/?sslmode=disable", user, password, host, port)
		w.db, err = sql.Open("postgres", dsn)
	case "Microsoft SQL Server":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%d", user, password, host, port)
		w.db, err = sql.Open("sqlserver", dsn)
	default:
		w.displayMessage("Unsupported database selected")
		return
	}

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

	w.showElementsafterConnect()
}

func (w *MainWindow) executeSQL(_ bool) {
	w.messagesText.Clear()
	w.statusLabel.SetText("")

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
		}()
	} else {
		w.statusLabel.SetText("Not connected to the database")
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

	w.errorLabel.SetText("")

	w.sqlLabel.Hide()
	w.sqlEntry.Hide()
	w.executeButton.Hide()
	w.resultLabel.Hide()
	w.resultText.Hide()
	w.messagesLabel.Hide()
	w.messagesText.Hide()
	w.exitButton.Hide()
	w.statusLabel.Hide()

	w.dbTypeComboBox.Show()
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

func (w *MainWindow) returnClicked(_ bool) {

	w.errorLabel.SetText("")

	w.connecttodbButton.Show()
	w.titleLabel.Show()
	w.versionLabel.Show()
	w.exitAppButton.Show()
	w.creatorLabel.Show()

	w.dbTypeComboBox.Hide()
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
	w.SetFixedSize2(700, 400)
}

func (w *MainWindow) exitApp(_ bool) {
	w.Close()
}

func Main() {

	// if runtime.GOOS == "linux" {
	// 	os.Setenv("QT_QPA_PLATFORM", "xcb") // Sets x11 for linux
	// }

	app := widgets.NewQApplication(len([]string{}), []string{})

	if core.QCoreApplication_Instance() == nil {
		color.Red("Failed to initialize QCoreApplication")
	}

	mainWindow := newMainWindow()
	mainWindow.Show()
	app.Exec()

}
