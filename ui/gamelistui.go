package ui

import (
	"fmt"
	"path/filepath"

	"github.com/therecipe/qt/widgets"
)

type GameListUI struct {
	*widgets.QMainWindow
	rui     *RetroUI
	gameloc string
}

func NewGameListUI(rui *RetroUI, gameloc string) *GameListUI {
	u := &GameListUI{rui: rui, gameloc: gameloc}
	u.init()
	return u
}

func (u *GameListUI) init() {
	u.QMainWindow = widgets.NewQMainWindow(u.rui, 0)
	u.SetWindowTitle(fmt.Sprintf("GameListEditor - %s", filepath.Base(u.gameloc)))

	var centralWidget = widgets.NewQWidget(u, 0)
	var layout = widgets.NewQGridLayout2()
	centralWidget.SetLayout(layout)
	u.SetCentralWidget(centralWidget)
}
