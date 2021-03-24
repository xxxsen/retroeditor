package ui

import (
	"fmt"
	"path/filepath"
	"retroeditor/parser"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type GameListUI struct {
	*widgets.QMainWindow
	rui     *RetroUI
	gameloc string
	glp     *parser.GameListParser

	//基础数据
	lePath        *widgets.QLineEdit
	leName        *widgets.QLineEdit
	picImage      *widgets.QLabel
	vVideo        *widgets.QLabel
	lePlayCnt     *widgets.QLineEdit
	leLastPlay    *widgets.QLineEdit
	leLang        *widgets.QLineEdit
	leDesc        *widgets.QTextEdit
	leRate        *widgets.QLineEdit
	leReleaseDate *widgets.QLineEdit
	leDev         *widgets.QLineEdit
	lePub         *widgets.QLineEdit
	leGenre       *widgets.QLineEdit
	lePlayer      *widgets.QLineEdit
	picMarquee    *widgets.QLabel
}

func NewGameListUI(rui *RetroUI, gameloc string) *GameListUI {
	u := &GameListUI{
		rui:           rui,
		gameloc:       gameloc,
		glp:           parser.NewGameListParser(gameloc + string(filepath.Separator) + "gamelist.xml"),
		lePath:        widgets.NewQLineEdit(nil),
		leName:        widgets.NewQLineEdit(nil),
		picImage:      widgets.NewQLabel(nil, 0),
		vVideo:        widgets.NewQLabel(nil, 0),
		lePlayCnt:     widgets.NewQLineEdit(nil),
		leLastPlay:    widgets.NewQLineEdit(nil),
		leLang:        widgets.NewQLineEdit(nil),
		leDesc:        widgets.NewQTextEdit(nil),
		leRate:        widgets.NewQLineEdit(nil),
		leReleaseDate: widgets.NewQLineEdit(nil),
		leDev:         widgets.NewQLineEdit(nil),
		lePub:         widgets.NewQLineEdit(nil),
		leGenre:       widgets.NewQLineEdit(nil),
		lePlayer:      widgets.NewQLineEdit(nil),
		picMarquee:    widgets.NewQLabel(nil, 0),
	}
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
	//添加游戏展示list
	layout.AddWidget2(lst, 0, 0, 0)

	//添加右边的布局
	var layoutRight = widgets.NewQGridLayout2()
	layout.AddLayout(layoutRight, 0, 1, 0)
	//生成基础属性布局
	var layoutBasicInfo = widgets.NewQGridLayout2()
	layoutRight.AddLayout(layoutBasicInfo, 0, 0, 0)
	//生成图片、视频框等布局
	var layoutView = widgets.NewQGridLayout2()
	layoutRight.AddLayout(layoutView, 0, 1, 0)

	//生成基础的属性框
	u.buildBasicInfo(layoutBasicInfo)
	//生成预览框
	u.buildPreviewInfo(layoutView)
}

func (u *GameListUI) buildBasicInfo(grid *widgets.QGridLayout) {
	var lbs = []*widgets.QLabel{
		widgets.NewQLabel2("游戏名", nil, 0),
		widgets.NewQLabel2("游戏路径", nil, 0),
		widgets.NewQLabel2("游玩次数", nil, 0),
		widgets.NewQLabel2("最后游玩时间", nil, 0),
		widgets.NewQLabel2("语言", nil, 0),
		widgets.NewQLabel2("好评率", nil, 0),
		widgets.NewQLabel2("发行时间", nil, 0),
		widgets.NewQLabel2("开发商", nil, 0),
		widgets.NewQLabel2("发行商", nil, 0),
		widgets.NewQLabel2("游戏类型", nil, 0),
		widgets.NewQLabel2("支持人数", nil, 0),
		widgets.NewQLabel2("描述", nil, 0),
	}
	for index := range lbs {
		grid.AddWidget2(lbs[index], index, 0, 0)
	}

	//
	u.leName.SetEnabled(false)
	grid.AddWidget2(u.leName, 0, 1, 0)
	grid.AddWidget2(u.lePath, 1, 1, 0)
	grid.AddWidget2(u.lePlayCnt, 2, 1, 0)
	grid.AddWidget2(u.leLastPlay, 3, 1, 0)
	grid.AddWidget2(u.leLang, 4, 1, 0)
	grid.AddWidget2(u.leRate, 5, 1, 0)
	grid.AddWidget2(u.leReleaseDate, 6, 1, 0)
	grid.AddWidget2(u.leDev, 7, 1, 0)
	grid.AddWidget2(u.lePub, 8, 1, 0)
	grid.AddWidget2(u.leGenre, 9, 1, 0)
	grid.AddWidget2(u.lePlayer, 10, 1, 0)
	grid.AddWidget2(u.leDesc, 11, 1, 0)
}

func (u *GameListUI) buildPreviewInfo(grid *widgets.QGridLayout) {
	var lbs = []*widgets.QLabel{
		widgets.NewQLabel2("封面图", nil, 0),
		widgets.NewQLabel2("相框", nil, 0),
		widgets.NewQLabel2("预览", nil, 0),
	}
	for index := range lbs {
		grid.AddWidget2(lbs[index], index, 0, 0)
	}
	//
	grid.AddWidget2(u.picImage, 0, 1, 0)
	grid.AddWidget2(u.picMarquee, 1, 1, 0)
	grid.AddWidget2(u.vVideo, 2, 1, 0)
}

func (u *GameListUI) safeGet(m map[string]string, key string) string {
	if v, ok := m[key]; ok {
		return v
	}
	return ""
}

func (u *GameListUI) OnDataNotify(m map[string]string) {
	u.leName.SetText(u.safeGet(m, "Name"))
	u.lePath.SetText(u.safeGet(m, "Path"))
	u.lePlayCnt.SetText(u.safeGet(m, "PlayCount"))
	u.leLastPlay.SetText(u.safeGet(m, "LastPlayed"))
	u.leLang.SetText(u.safeGet(m, "Lang"))
	u.leRate.SetText(u.safeGet(m, "Rating"))
	u.leReleaseDate.SetText(u.safeGet(m, "ReleaseDate"))
	u.leDev.SetText(u.safeGet(m, "Developer"))
	u.lePub.SetText(u.safeGet(m, "Publisher"))
	u.leGenre.SetText(u.safeGet(m, "Genre"))
	u.lePlayer.SetText(u.safeGet(m, "Players"))
	u.leDesc.SetText(u.safeGet(m, "Desc"))
	img := u.safeGet(m, "Image")
	if len(img) != 0 {
		qp := gui.NewQPixmap3(img, "", 0)
		u.picImage.SetPixmap(qp)
	}
	marquee := u.safeGet(m, "Marquee")
	if len(marquee) != 0 {
		qp := gui.NewQPixmap3(marquee, "", 0)
		u.picMarquee.SetPixmap(qp)
	}
	//TODO:add video support
	// video := u.safeGet(m, "Video")
	// if len(video) != 0 {
	// 	player := multimedia.NewQMediaPlayer(nil, 0)
	// 	player.SetMedia(multimedia.NewQMediaContent2(core.QUrl_FromLocalFile(video)), nil)
	// 	player.Play()
	// }
}
