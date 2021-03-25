package ui

import (
	"retroeditor/fs"
	"retroeditor/parser"
	"retroeditor/uigen"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

//Gamedit Gamedit
type GameditUI struct {
	*widgets.QDialog
	form     *uigen.UIGameditDialog
	gamelist *GameListUI
	mode     int32
}

const (
	UI_GeModeCreate = iota
	UI_GeModeModify
)

func NewGameditUI(name string, gamelist *GameListUI, mode int32, data *parser.GameListItem) *GameditUI {
	u := &GameditUI{QDialog: widgets.NewQDialog(nil, 0), gamelist: gamelist, mode: mode}
	u.form = &uigen.UIGameditDialog{}
	u.form.SetupUI(u.QDialog)
	u.QDialog.SetWindowTitle(name)
	u.init()
	if mode == UI_GeModeModify {
		u.form.LePath.SetEnabled(false)
	} else {
		u.form.LePath.SetEnabled(true)
	}
	u.setDefaultData(data)
	return u
}

func (u *GameditUI) setDefaultData(m *parser.GameListItem) {
	if m == nil {
		return
	}
	u.form.LePath.SetEnabled(false)
	u.form.LeName.SetText(m.Name)
	u.form.LePath.SetText(m.Path)
	u.form.LePlayCnt.SetText(m.PlayCount)
	u.form.LeLastPlay.SetText(m.LastPlayed)
	u.form.LeLang.SetText(m.Lang)
	u.form.LeRate.SetText(m.Rating)
	u.form.LeReleaseDate.SetText(m.ReleaseDate)
	u.form.LeDev.SetText(m.Developer)
	u.form.LePub.SetText(m.Publisher)
	u.form.LeGenre.SetText(m.Genre)
	u.form.LePlayer.SetText(m.Players)
	u.form.LeDesc.SetText(m.Desc)
	u.loadImage(m.Image)
	u.loadMarquee(m.Marquee)
	u.loadVideo(m.Video)
}

func (u *GameditUI) loadVideo(video string) {
	//TODO:
}

func (u *GameditUI) init() {
	desktop := widgets.QApplication_Desktop()
	u.Move2((desktop.Width()-u.Width())/2, (desktop.Height()-u.Height())/2)

	//
	u.form.BtnCancel.ConnectClicked(func(bool) { u.Close() })
	u.form.BtnSave.ConnectClicked(func(bool) { u.onSave() })
	u.form.LePath.ConnectMousePressEvent(u.onPathClick)
	u.form.PicImage.ConnectMousePressEvent(u.onImageClick)
	u.form.PicMarquee.ConnectMousePressEvent(u.onMarqueeClick)

	//输入类型限制
	u.form.LePlayCnt.SetValidator(gui.NewQIntValidator2(0, 100000, u))
	u.form.LePlayer.SetValidator(gui.NewQRegExpValidator2(core.NewQRegExp2("^\\d+-\\d+|\\d+$", 0, 0), nil))
}

func (u *GameditUI) onSave() {
	if len(u.form.LeName.Text()) == 0 || len(u.form.LePath.Text()) == 0 {
		NoticeMessagef("需要填写游戏名/游戏路径!")
		return
	}
	u.AcceptDefault()
}

func (u *GameditUI) onImageClick(*gui.QMouseEvent) {
	file := widgets.QFileDialog_GetOpenFileName(nil, "选择封面文件:",
		u.gamelist.gameloc, "*.jpg *.png *.jpeg  *.gif", "*.*", widgets.QFileDialog__ReadOnly)
	if len(file) == 0 {
		return
	}
	ok := fs.IsParentDir(u.gamelist.gameloc, file)
	if !ok {
		NoticeMessagef("图片:%s 不存在roms目录中, 请手动复制进去。", file)
		return
	}
	sub := "./" + fs.TrimPath(u.gamelist.gameloc, file)
	u.loadImage(sub)
}

func (u *GameditUI) onMarqueeClick(*gui.QMouseEvent) {
	file := widgets.QFileDialog_GetOpenFileName(nil, "选择封面文件:",
		u.gamelist.gameloc, "*.jpg *.png *.jpeg  *.gif", "*.*", widgets.QFileDialog__ReadOnly)
	if len(file) == 0 {
		return
	}
	ok := fs.IsParentDir(u.gamelist.gameloc, file)
	if !ok {
		NoticeMessagef("图片:%s 不存在roms目录中, 请手动复制进去。", file)
		return
	}
	sub := "./" + fs.TrimPath(u.gamelist.gameloc, file)
	u.loadMarquee(sub)
}

func (u *GameditUI) loadImage(img string) {
	u.form.PicImage.SetToolTip(img)
	if len(img) != 0 {
		qp := gui.NewQPixmap3(fs.MergePath(u.gamelist.gameloc, img), "", 0)
		qp = qp.Scaled2(200, 200, core.Qt__KeepAspectRatio, core.Qt__FastTransformation)
		u.form.PicImage.SetPixmap(qp)
	}
}

func (u *GameditUI) loadMarquee(img string) {
	u.form.PicMarquee.SetToolTip(img)
	if len(img) != 0 {
		qp := gui.NewQPixmap3(fs.MergePath(u.gamelist.gameloc, img), "", 0)
		qp = qp.Scaled2(200, 200, core.Qt__KeepAspectRatio, core.Qt__FastTransformation)
		u.form.PicMarquee.SetBaseSize2(100, 100)
		u.form.PicMarquee.SetPixmap(qp)
	}
}

func (u *GameditUI) onPathClick(*gui.QMouseEvent) {
	if u.mode == UI_GeModeModify {
		return
	}

	file := widgets.QFileDialog_GetOpenFileName(nil, "选择游戏文件:",
		u.gamelist.gameloc, "*.*", "*.*", widgets.QFileDialog__ReadOnly)
	if len(file) == 0 {
		return
	}
	ok := fs.IsParentDir(u.gamelist.gameloc, file)
	if !ok {
		NoticeMessagef("游戏:%s 不存在roms目录中, 请手动复制进去。", file)
		return
	}
	sub := "./" + fs.TrimPath(u.gamelist.gameloc, file)
	if _, ok := u.gamelist.glp.Get(sub); ok {
		NoticeMessagef("游戏:%s 已存在列表中!", sub)
		return
	}

	u.form.LePath.SetText(sub)
}
