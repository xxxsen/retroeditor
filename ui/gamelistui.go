package ui

import (
	"fmt"
	"path/filepath"
	"retroeditor/parser"

	"github.com/therecipe/qt/widgets"
)

type GameListUI struct {
	*widgets.QMainWindow
	rui     *RetroUI
	gameloc string
	glp     *parser.GameListParser
}

func NewGameListUI(rui *RetroUI, gameloc string) *GameListUI {
	u := &GameListUI{rui: rui, gameloc: gameloc, glp: parser.NewGameListParser(gameloc + string(filepath.Separator) + "gamelist.xml")}
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

	err := u.glp.Parse()
	if err != nil {
		FatalError(err)
	}
	lst := widgets.NewQListWidget(nil)

	for name := range u.glp.GetAll() {
		lst.AddItem(name)
	}
	layout.AddWidget2(lst, 0, 0, 0)
}
