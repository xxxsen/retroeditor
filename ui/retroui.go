package ui

import "github.com/therecipe/qt/widgets"

type RetroUI struct {
	*widgets.QMainWindow
}

func NewRetroUI() *RetroUI {
	u := &RetroUI{}
	u.init()
	return u
}

func (r *RetroUI) init() {
	r.QMainWindow = widgets.NewQMainWindow(nil, 0)
	r.SetWindowTitle("RetroEditor")

	var centralWidget = widgets.NewQWidget(r, 0)
	var layout = widgets.NewQGridLayout2()
	centralWidget.SetLayout(layout)
	r.SetCentralWidget(centralWidget)
}
