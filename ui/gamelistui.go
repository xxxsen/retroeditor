package ui

import (
	"fmt"
	"path/filepath"
	"retroeditor/fs"
	"retroeditor/parser"
	"retroeditor/uigen"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type GameListUI struct {
	//
	*widgets.QWidget
	form *uigen.UIGamelistForm

	//
	gameloc string
	glp     *parser.GameListParser
}

func NewGameListUI(gameloc string) *GameListUI {
	u := &GameListUI{gameloc: gameloc, QWidget: widgets.NewQWidget(nil, 0), glp: parser.NewGameListParser(gameloc + "/" + "gamelist.xml")}
	u.form = &uigen.UIGamelistForm{}
	u.form.SetupUI(u.QWidget)
	u.SetWindowTitle(fmt.Sprintf("GameListEditor - %s", filepath.Base(u.gameloc)))
	u.init()
	return u
}

func (u *GameListUI) init() {
	desktop := widgets.QApplication_Desktop()
	u.Move2((desktop.Width()-u.Width())/2, (desktop.Height()-u.Height())/2)

	//解析列表
	err := u.glp.Parse()
	if err != nil {
		FatalError(err)
	}
	//列表配置数据
	lst := u.glp.GetList()
	for _, item := range lst {
		u.form.LstGame.AddItem(item)
	}
	u.form.LstGame.ConnectItemSelectionChanged(u.onListItemSelect)
	if u.form.LstGame.Count() > 0 {
		u.form.LstGame.SetCurrentRow(0)
	}

	//设置按钮回调
	u.form.BtnDelete.ConnectClicked(u.onGameDelete)
	u.form.BtnCreate.ConnectClicked(u.onGameCreate)
	u.form.BtnMod.ConnectClicked(u.onGameMod)
	u.form.BtnRebuild.ConnectClicked(u.onRebuildGameList)
	u.form.BtnWriteFile.ConnectClicked(u.onWriteToFile)
}

func (u *GameListUI) onWriteToFile(bool) {
	err := u.glp.Save()
	if err != nil {
		NoticeMessagef("数据写入失败, 错误:%v", err)
		return
	}
	NoticeMessagef("数据写入成功!")
}

func (u *GameListUI) onRebuildGameList(bool) {
	dig := NewGameRebuildUI(u)
	dig.Exec()
}

func (u *GameListUI) onGameMod(bool) {
	if u.form.LstGame.CurrentRow() < 0 {
		return
	}
	old, _ := u.glp.Get(u.form.LstGame.CurrentItem().Text())
	editor := NewGameditUI(fmt.Sprintf("修改 - %s", old.Name), u, UI_GeModeModify, old)
	rs := editor.Exec()
	if rs != int(widgets.QDialog__Accepted) {
		return
	}
	item := u.readPropertyFromEditor(editor)
	u.glp.Set(item)
	u.onDataNotify(item)
}

func (u *GameListUI) readPropertyFromEditor(editor *GameditUI) *parser.GameListItem {
	item := &parser.GameListItem{}
	//设置属性
	item.Name = editor.form.LeName.Text()
	item.Path = editor.form.LePath.Text()
	item.PlayCount = editor.form.LePlayCnt.Text()
	item.LastPlayed = editor.form.LeLastPlay.Text()
	item.Lang = editor.form.LeLang.Text()
	item.Rating = editor.form.LeRate.Text()
	item.ReleaseDate = editor.form.LeReleaseDate.Text()
	item.Developer = editor.form.LeDev.Text()
	item.Publisher = editor.form.LePub.Text()
	item.Genre = editor.form.LeGenre.Text()
	item.Players = editor.form.LePlayer.Text()
	item.Desc = editor.form.LeDesc.ToPlainText()
	item.Image = editor.form.PicImage.ToolTip()
	item.Marquee = editor.form.PicMarquee.ToolTip()
	item.Video = editor.form.VVideo.ToolTip()
	return item
}

func (u *GameListUI) onGameCreate(bool) {
	editor := NewGameditUI("新游戏", u, UI_GeModeCreate, nil)
	rs := editor.Exec()
	if rs != int(widgets.QDialog__Accepted) {
		return
	}
	item := u.readPropertyFromEditor(editor)
	u.glp.Set(item)
	u.form.LstGame.AddItem(item.Path)
	u.form.LstGame.SetCurrentRow(u.form.LstGame.Count() - 1)
}

func (u *GameListUI) onGameDelete(bool) {
	if u.form.LstGame.CurrentRow() < 0 {
		return
	}
	name := u.form.LePath.Text()
	u.glp.Remove(name)
	u.form.LstGame.TakeItem(u.form.LstGame.CurrentIndex().Row())
}

func (u *GameListUI) onListItemSelect() {
	if u.form.LstGame.CurrentRow() == -1 {
		return
	}
	name := u.form.LstGame.CurrentItem().Text()
	item, found := u.glp.Get(name)
	if !found {
		NoticeMessagef("未找到游戏:%s", name)
		return
	}
	u.onDataNotify(item)
}

func (u *GameListUI) onDataNotify(m *parser.GameListItem) {
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

func (u *GameListUI) loadVideo(video string) {
	//TODO:
}

func (u *GameListUI) loadImage(img string) {
	u.form.PicImage.SetToolTip(img)
	if len(img) != 0 {
		qp := gui.NewQPixmap3(fs.MergePath(u.gameloc, img), "", 0)
		qp = qp.Scaled2(200, 200, core.Qt__KeepAspectRatio, core.Qt__FastTransformation)
		u.form.PicImage.SetPixmap(qp)
	} else {
		u.form.PicImage.Clear()
	}
}

func (u *GameListUI) loadMarquee(img string) {
	u.form.PicMarquee.SetToolTip(img)
	if len(img) != 0 {
		qp := gui.NewQPixmap3(fs.MergePath(u.gameloc, img), "", 0)
		qp = qp.Scaled2(200, 200, core.Qt__KeepAspectRatio, core.Qt__FastTransformation)
		u.form.PicMarquee.SetBaseSize2(100, 100)
		u.form.PicMarquee.SetPixmap(qp)
	} else {
		u.form.PicMarquee.Clear()
	}
}
