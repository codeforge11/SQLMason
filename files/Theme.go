package appdata

import "github.com/therecipe/qt/widgets"

func Theme(app *widgets.QApplication) {
	app.SetStyleSheet(`

        QMainWindow {
            background-color: #5FAD41;
            color: #000000; 
        }
        QLabel, QLineEdit, QTextEdit {
            background-color: #5FAD41;
            color: #000000;
        }
        QPushButton {
            background-color: #3a86ff;
            color: white;
            font-size: 14px;
        }
        QLineEdit, QTextEdit {
            border: 1px solid #3a86ff;
            padding: 5px;
        }
        QPushButton {
            background-color: #3a86ff;
            color: white;
            font-size: 14px;
        }



    `)
}
