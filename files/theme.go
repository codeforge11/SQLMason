package appdata

import "github.com/therecipe/qt/widgets"

func Theme(app *widgets.QApplication) {
	app.SetStyleSheet(`
        QLabel#titleLabel {
            font-size: 24px;
            color: black;
            padding: 10px;
            text-align: center;
        }
        QPushButton#connecttodbButton{
            text-align: center;
            margin-left: auto;
            margin-right: auto;
            background-color: #3a86ff;
            color: white;
            font-size: 72px; 
            padding: 40px; 
            border: none;
            border-radius: 10px; 
        }
    `)
}
