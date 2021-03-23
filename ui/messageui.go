package ui

import (
	"github.com/therecipe/qt/widgets"
)

func FatalError(err error) {
	FatalMessage(err.Error())
}

func FatalMessage(errmsg string) {
	box := widgets.NewQMessageBox(nil)
	box.SetText(errmsg)
	box.Show()
	//TODO:should exit program...
}
