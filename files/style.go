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

       .button-64 {
            align-items: center;
            background-image: linear-gradient(144deg,#AF40FF, #5B42F3 50%,#00DDEB);
            border: 0;
            border-radius: 8px;
            color: #FFFFFF;
            display: flex;
            font-family: Phantomsans, sans-serif;
            font-size: 20px;
            justify-content: center;
            line-height: 1em;
            max-width: 100%;
            min-width: 140px;
            padding: 3px;
            text-decoration: none;
        }

       .button-64:active,
       .button-64:hover {
            outline: 0;
        }

       .button-64 span {
            background-color: rgb(5, 6, 45);
            padding: 16px 24px;
            border-radius: 6px;
            width: 100%;
            height: 100%;
        }

       .button-64:hover span {
            background: none;
        }

        @media (min-width: 768px) {
           .button-54 {
                padding: 0.25em 0.75em;
            }
            QPushButton#connectButton {
                background-color: #3a86ff;
                color: white; 
                font-size: 36px;
            }
            QPushButton#executeButton {
                background-color: #3a86ff; 
                color: white; 
                font-size: 14px;
            }
        }
    `)
}
