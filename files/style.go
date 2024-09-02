package appdata

import (
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/therecipe/qt/widgets"
)

func Style(app *widgets.QApplication) {
	color.HiGreen("Style file is running")

	css, err := ioutil.ReadFile("source/style.css")
	if err != nil {
		panic(err)
	}

	app.SetStyleSheet(string(css))
}
