package appdata

import (
	"github.com/fatih/color"
	"github.com/therecipe/qt/widgets"
)

func Style(app *widgets.QApplication) {
	color.HiGreen("Style file is running")

	app.SetStyleSheet(`
        QLabel#titleLabel {
            font-size: 24px;
            color: black;
            padding: 10px;
            text-align: center;
        }
        QPushButton#connecttodbButton {
            font-family: "Open Sans", sans-serif;
            font-size: 16px;
            text-decoration: none;
            text-transform: uppercase;
            color: #000;
            border: 3px solid;
            padding: 0.25em 0.5em;
            position: relative;
        }

        QPushButton#connecttodbButton:hover {
            background-color: #4CAF50;
            color: #fff;
        }
        QPushButton#connecttodbButton:active{
            top: 5px;
            left: 5px;
        }

        QPushButton#connectButton {
            background-color: #3a86ff;
            color: black; 
            font-size: 18px;
        }
        QPushButton#executeButton {
            background-color: #3a86ff; 
            color: black; 
            font-size: 14px;
        }
        QPushButton#returnButton{
            font-size: 18px;
        }
        QLabel#errorLabel{
            color: red
        }
    `)
}
