package ui

import (
	"path/filepath"
	"retroeditor/mgr"

	"github.com/therecipe/qt/widgets"
)

type RetroUI struct {
	*widgets.QMainWindow
	dirmgr *mgr.DirManager
}

func NewRetroUI(base string) *RetroUI {
	u := &RetroUI{dirmgr: mgr.NewDirManager(base)}
	u.init()
	return u
}

func (r *RetroUI) init() {
	r.QMainWindow = widgets.NewQMainWindow(nil, 0)
	desktop := widgets.QApplication_Desktop()
	r.Move2((desktop.Width()-r.Width())/2, (desktop.Height()-r.Height())/2)
	r.SetWindowTitle("RetroEditor")

	var centralWidget = widgets.NewQWidget(r, 0)
	var layout = widgets.NewQGridLayout2()
	centralWidget.SetLayout(layout)
	r.SetCentralWidget(centralWidget)
	dirs, err := r.dirmgr.GetDirs()
	if err != nil {
		FatalError(err)
	}
	//10x10
	nextline := 0
	for index, name := range dirs {
		if index != 0 && index%10 == 0 {
			nextline++
		}
		btn := widgets.NewQPushButton2(name, nil)
		btn.ConnectClicked(r.btnCallback(btn))
		layout.AddWidget2(btn, nextline, index%10, 0)
	}
}

func (r *RetroUI) btnCallback(btn *widgets.QPushButton) func(ch bool) {
	return func(bool) {
		gamedir := r.dirmgr.GetBase() + string(filepath.Separator) + btn.Text()
		gameui := NewGameListUI(gamedir)
		gameui.Show()
	}
}
