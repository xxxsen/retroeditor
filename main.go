package main

import (
	"os"
	"retroeditor/ui"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	window := ui.NewRetroUI()
	window.Show()
	widgets.QApplication_Exec()
}
