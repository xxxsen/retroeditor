package ui

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"retroeditor/fs"
	"retroeditor/mgr"
	"retroeditor/parser"
	"retroeditor/uigen"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//GameRebuild GameRebuild
type GameRebuildUI struct {
	*widgets.QDialog
	gamelist *GameListUI
	form     *uigen.UIGamecleanDialog

	//只从有效ext列表里面扫描, 避免扫到异常的文件
	eraseXmlList   []string
	addXmlMap      map[string]string
	eraseMediaList []string
	eraseRomList   []string
}

//NewGameRebuildUI NewGameRebuildUI
func NewGameRebuildUI(gamelist *GameListUI) *GameRebuildUI {
	u := &GameRebuildUI{QDialog: widgets.NewQDialog(nil, 0), gamelist: gamelist,
		addXmlMap: make(map[string]string)}
	u.form = &uigen.UIGamecleanDialog{}
	u.form.SetupUI(u.QDialog)
	u.init()
	return u
}

//
func (u *GameRebuildUI) init() {
	desktop := widgets.QApplication_Desktop()
	u.Move2((desktop.Width()-u.Width())/2, (desktop.Height()-u.Height())/2)

	u.form.BtnExec.SetEnabled(false)
	u.form.BtnCancel.ConnectClicked(func(bool) { u.Close() })
	u.form.BtnScan.ConnectClicked(u.onScan)
	u.form.BtnExec.ConnectClicked(u.onExec)

	u.form.CbAddXml.ConnectClicked(func(ck bool) {
		if ck {
			u.form.CbCleanRom.SetChecked(false)
		}
	})
	u.form.CbCleanRom.ConnectClicked(func(ck bool) {
		if ck {
			u.form.CbAddXml.SetChecked(false)
		}
	})

	u.buildSupportExt()
}

func (u *GameRebuildUI) buildSupportExt() {
	//生成当前有效的扩展名
	m := u.gamelist.glp.GetAll()

	romsExt := make(map[string]bool)
	picExt := make(map[string]bool)
	videoExt := make(map[string]bool)
	var romExtList []string
	var picExtList []string
	var videoExtList []string

	for _, item := range m {
		romsExt[filepath.Ext(item.Path)] = true
		picExt[filepath.Ext(item.Image)] = true
		picExt[filepath.Ext(item.Marquee)] = true
		videoExt[filepath.Ext(item.Video)] = true
	}

	for ext := range romsExt {
		if len(ext) == 0 {
			continue
		}
		romExtList = append(romExtList, ext)
	}
	for ext := range picExt {
		if len(ext) == 0 {
			continue
		}
		picExtList = append(picExtList, ext)
	}
	for ext := range videoExt {
		if len(ext) == 0 {
			continue
		}
		videoExtList = append(videoExtList, ext)
	}
	u.form.LeRomExt.SetText(strings.Join(romExtList, ";"))
	u.form.LePicExt.SetText(strings.Join(picExtList, ";"))
	u.form.LeVideoExt.SetText(strings.Join(videoExtList, ";"))
}

func (u *GameRebuildUI) onScan(bool) {
	defer u.form.BtnExec.SetEnabled(true)
	u.form.BtnScan.SetEnabled(false)
	defer u.form.BtnScan.SetEnabled(true)
	u.form.LstAdd.Clear()
	u.form.LstClean.Clear()
	u.form.LstCleanMedia.Clear()
	u.form.LstCleanRom.Clear()
	if u.form.CbCleanXml.CheckState() == core.Qt__Checked {
		u.scanCleanXML()
	}
	if u.form.CbAddXml.CheckState() == core.Qt__Checked {
		u.scanAddXML()
	}
	if u.form.CbCleanMedia.CheckState() == core.Qt__Checked {
		u.scanCleanMedia()
	}
	if u.form.CbCleanRom.CheckState() == core.Qt__Checked {
		u.scanCleanRom()
	}
	NoticeMessagef("扫描完成!")
}

func (u *GameRebuildUI) isValidExt(ext string) bool {
	if len(ext) == 0 || ext == "." || !strings.HasPrefix(ext, ".") {
		return false
	}
	return true
}

func (u *GameRebuildUI) scanCleanRom() {
	exts := strings.Split(u.form.LeRomExt.Text(), ";")
	filters := make(map[string]bool)
	for _, ext := range exts {
		if !u.isValidExt(ext) {
			continue
		}
		filters[ext] = true
	}
	if len(filters) == 0 {
		NoticeMessagef("未找到rom扩展名信息, 跳过rom清理扫描")
		return
	}

	fm := mgr.NewFileMgr(u.gamelist.gameloc)

	lst, err := fm.Search(filters)
	if err != nil {
		NoticeMessagef("从文件搜索rom信息异常, 错误:%v", err)
		return
	}
	m := u.gamelist.glp.GetAll()
	for _, item := range lst {
		sub := "./" + fs.TrimPath(u.gamelist.gameloc, item)
		if _, ok := m[sub]; ok {
			continue
		}
		u.eraseRomList = append(u.eraseRomList, item)
		u.form.LstCleanRom.AddItem(sub)
	}
}

func (u *GameRebuildUI) scanCleanMedia() {
	exts := strings.Split(u.form.LePicExt.Text()+";"+u.form.LeVideoExt.Text(), ";")
	filters := make(map[string]bool)
	for _, ext := range exts {
		if !u.isValidExt(ext) {
			continue
		}
		filters[ext] = true
	}
	if len(filters) == 0 {
		NoticeMessagef("未找到图片/视频扩展名信息, 跳过媒体清理扫描")
		return
	}

	fm := mgr.NewFileMgr(u.gamelist.gameloc)
	lst, err := fm.Search(filters)
	if err != nil {
		NoticeMessagef("从文件搜索媒体信息异常, 错误:%v", err)
		return
	}
	existMedia := make(map[string]bool)
	m := u.gamelist.glp.GetAll()
	for _, item := range m {
		if len(item.Image) != 0 {
			existMedia[item.Image] = true
		}
		if len(item.Marquee) != 0 {
			existMedia[item.Marquee] = true
		}
		if len(item.Video) != 0 {
			existMedia[item.Video] = true
		}
	}

	for _, item := range lst {
		sub := "./" + fs.TrimPath(u.gamelist.gameloc, item)
		if _, ok := existMedia[sub]; ok {
			continue
		}
		u.eraseMediaList = append(u.eraseMediaList, item)
		u.form.LstCleanMedia.AddItem(item)
	}
}

func (u *GameRebuildUI) scanAddXML() {
	exts := strings.Split(u.form.LeRomExt.Text(), ";")
	filters := make(map[string]bool)
	for _, ext := range exts {
		if !u.isValidExt(ext) {
			continue
		}
		filters[ext] = true
	}
	if len(filters) == 0 {
		NoticeMessagef("未找到rom扩展名信息, 跳过rom扫描")
		return
	}

	fm := mgr.NewFileMgr(u.gamelist.gameloc)

	lst, err := fm.Search(filters)
	if err != nil {
		NoticeMessagef("从文件搜索rom信息异常, 错误:%v", err)
		return
	}
	for _, item := range lst {
		sub := "./" + fs.TrimPath(u.gamelist.gameloc, item)
		if _, ok := u.gamelist.glp.Get(sub); ok {
			continue
		}
		fullname := path.Base(item)
		extname := path.Ext(item)
		nameonly := fullname[0 : len(fullname)-len(extname)]
		u.addXmlMap[sub] = nameonly
		u.form.LstAdd.AddItem(sub)
	}
}

func (u *GameRebuildUI) scanCleanXML() {
	m := u.gamelist.glp.GetAll()
	for fp := range m {
		loc := fs.MergePath(u.gamelist.gameloc, fp)
		exist, err := fs.IsExist(loc)
		if err != nil {
			log.Printf("Check exist fail, path:%s, err:%v", loc, err)
			continue
		}
		if !exist {
			u.eraseXmlList = append(u.eraseXmlList, fp)
			u.form.LstClean.AddItem(fp)
		}
	}
}

func (u *GameRebuildUI) onExec(bool) {
	defer u.form.BtnExec.SetEnabled(false)

	wouldUpdate := false
	if u.form.CbCleanXml.CheckState() == core.Qt__Checked {
		wouldUpdate = true
		for _, fp := range u.eraseXmlList {
			u.gamelist.glp.Remove(fp)
		}
	}
	if u.form.CbAddXml.CheckState() == core.Qt__Checked {
		wouldUpdate = true
		for fp, name := range u.addXmlMap {
			item := &parser.GameListItem{Name: name, Path: fp}
			u.gamelist.glp.Set(item)
		}
	}
	if u.form.CbCleanMedia.CheckState() == core.Qt__Checked {
		for _, item := range u.eraseMediaList {
			err := os.Remove(item)
			if err != nil {
				log.Printf("Remove media item:%s fail, err:%v", item, err)
				continue
			}
			log.Printf("Remove media file:%s", item)
		}
	}
	if u.form.CbCleanRom.CheckState() == core.Qt__Checked {
		for _, item := range u.eraseRomList {
			err := os.Remove(item)
			if err != nil {
				log.Printf("Remove rom item:%s fail, err:%v", item, err)
				continue
			}
			log.Printf("Remove rom file:%s", item)
		}
	}
	//
	if wouldUpdate {
		u.gamelist.form.LstGame.SetCurrentRow(-1)
		u.gamelist.form.LstGame.Clear()
		u.gamelist.form.LstGame.DisconnectItemSelectionChanged()
		lst := u.gamelist.glp.GetList()
		for _, item := range lst {
			u.gamelist.form.LstGame.AddItem(item)
		}
		u.gamelist.form.LstGame.ConnectItemSelectionChanged(u.gamelist.onListItemSelect)
	}
	//保存到文件
	u.gamelist.glp.Save()
	NoticeMessagef("重建完成, 清理XML:%d项, 添加XML:%d项, 清理媒体文件:%d项, 清理rom文件:%d项",
		len(u.eraseXmlList), len(u.addXmlMap), len(u.eraseMediaList), len(u.eraseRomList))
}
