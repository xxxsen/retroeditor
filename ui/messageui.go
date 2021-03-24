package ui

import (
	"fmt"

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

func NoticeMessage(msg string) {
	box := widgets.NewQMessageBox(nil)
	box.SetText(msg)
	box.Show()
}

func NoticeMessagef(fmtx string, args ...interface{}) {
	msg := fmt.Sprintf(fmtx, args...)
	NoticeMessage(msg)
}
