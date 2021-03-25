package ui

import (
	"fmt"
	"log"
	"path/filepath"
	"retroeditor/fs"
	"retroeditor/mgr"
	"retroeditor/parser"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/multimedia"
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
	vVideo        *multimedia.QVideoWidget
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

	//按钮
	btnDelete    *widgets.QPushButton
	btnAdd       *widgets.QPushButton
	btnSave      *widgets.QPushButton
	btnDiscard   *widgets.QPushButton
	btnWriteFile *widgets.QPushButton
	btnClean     *widgets.QPushButton

	//list
	lstGame *widgets.QListWidget

	//videoplayer
	videoplayer *multimedia.QMediaPlayer
}

func NewGameListUI(rui *RetroUI, gameloc string) *GameListUI {
	u := &GameListUI{
		rui:           rui,
		gameloc:       gameloc,
		glp:           parser.NewGameListParser(gameloc + string(filepath.Separator) + "gamelist.xml"),
		lePath:        widgets.NewQLineEdit(nil),
		leName:        widgets.NewQLineEdit(nil),
		picImage:      widgets.NewQLabel(nil, 0),
		vVideo:        multimedia.NewQVideoWidget(nil),
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
		videoplayer:   multimedia.NewQMediaPlayer(nil, multimedia.QMediaPlayer__VideoSurface),
	}
	u.init()
	return u
}

func (u *GameListUI) init() {
	u.QMainWindow = widgets.NewQMainWindow(u.rui, 0)
	u.SetMinimumSize2(1280, 720)
	desktop := widgets.QApplication_Desktop()
	u.Move2((desktop.Width()-u.Width())/2, (desktop.Height()-u.Height())/2)
	u.SetWindowTitle(fmt.Sprintf("GameListEditor - %s", filepath.Base(u.gameloc)))
	var centralWidget = widgets.NewQWidget(u, 0)

	//基础布局
	basicLayout := widgets.NewQGridLayout2()
	centralWidget.SetLayout(basicLayout)
	u.SetCentralWidget(centralWidget)

	err := u.glp.Parse()
	if err != nil {
		FatalError(err)
	}
	u.lstGame = widgets.NewQListWidget(nil)
	u.lstGame.ConnectItemSelectionChanged(u.onListItemSelect)

	//控制按钮布局
	controlLayout := widgets.NewQGridLayout2()
	u.buildControl(controlLayout)
	//展示布局
	layout := widgets.NewQGridLayout2()

	basicLayout.AddLayout(layout, 0, 0, 0)
	basicLayout.AddLayout(controlLayout, 1, 0, 0)

	//添加游戏展示list
	layout.AddWidget2(u.lstGame, 0, 0, 0)

	//添加右边的布局
	var layoutRight = widgets.NewQGridLayout2()
	layout.AddLayout(layoutRight, 0, 1, 0)
	//生成基础属性布局
	var layoutBasicInfo = widgets.NewQGridLayout2()

	//设置比例
	layout.SetColumnStretch(0, 20)
	layout.SetColumnStretch(1, 80)
	layoutRight.SetColumnStretch(0, 80)
	layoutRight.SetColumnStretch(1, 20)

	layoutRight.AddLayout(layoutBasicInfo, 0, 0, 0)
	//生成图片、视频框等布局
	var layoutView = widgets.NewQGridLayout2()
	layoutRight.AddLayout(layoutView, 0, 1, 0)

	//生成基础的属性框
	u.buildBasicInfo(layoutBasicInfo)
	//生成预览框
	u.buildPreviewInfo(layoutView)

	//构建list
	u.buildList()
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

	//设置编辑回调
	u.leName.ConnectTextEdited(u.onModify)
	u.lePath.ConnectTextEdited(u.onModify)
	u.lePlayCnt.ConnectTextEdited(u.onModify)
	u.leLastPlay.ConnectTextEdited(u.onModify)
	u.leLang.ConnectTextEdited(u.onModify)
	u.leRate.ConnectTextEdited(u.onModify)
	u.leReleaseDate.ConnectTextEdited(u.onModify)
	u.leDev.ConnectTextEdited(u.onModify)
	u.lePub.ConnectTextEdited(u.onModify)
	u.leGenre.ConnectTextEdited(u.onModify)
	u.lePlayer.ConnectTextEdited(u.onModify)
	u.leDesc.ConnectTextChanged(func() { u.onModify("") })

	//增加选框支持
	u.lePath.ConnectMousePressEvent(u.onPathClick)

}

func (u *GameListUI) buildControl(grid *widgets.QGridLayout) {
	u.btnDelete = widgets.NewQPushButton2("删除游戏(&D)", u)
	u.btnAdd = widgets.NewQPushButton2("添加游戏(&N)", u)
	u.btnSave = widgets.NewQPushButton2("保存修改(&S)", u)
	u.btnDiscard = widgets.NewQPushButton2("还原修改(&C)", u)
	u.btnWriteFile = widgets.NewQPushButton2("写入文件(&W)", u)
	u.btnClean = widgets.NewQPushButton2("清理无效文件(&E)", u)

	//初始不能修改
	u.btnSave.SetEnabled(false)
	u.btnDiscard.SetEnabled(false)
	u.btnDelete.SetEnabled(false)
	u.btnWriteFile.SetEnabled(false)

	//生成2层布局
	layoutNormalCtl := widgets.NewQGridLayout2()
	layoutSpecCtl := widgets.NewQGridLayout2()
	grid.AddLayout(layoutNormalCtl, 0, 0, 0)
	grid.AddLayout(layoutSpecCtl, 1, 0, 0)

	//添加到布局里面
	//普通按钮
	layoutNormalCtl.AddWidget2(u.btnDelete, 0, 0, 0)
	layoutNormalCtl.AddWidget2(u.btnAdd, 0, 1, 0)
	layoutNormalCtl.AddWidget2(u.btnSave, 0, 2, 0)
	layoutNormalCtl.AddWidget2(u.btnDiscard, 0, 3, 0)
	layoutNormalCtl.AddWidget2(u.btnWriteFile, 0, 4, 0)
	//特殊按钮
	layoutSpecCtl.AddWidget2(u.btnClean, 0, 0, 0)

	//设置回调
	u.btnSave.ConnectClicked(func(bool) { u.onDataSave() })
	u.btnDiscard.ConnectClicked(func(bool) { u.onDataDiscard() })
	u.btnAdd.ConnectClicked(func(bool) { u.onDataCreate() })
	u.btnDelete.ConnectClicked(func(bool) { u.onDataDelete() })
	u.btnWriteFile.ConnectClicked(func(bool) { u.onWriteToFile() })
	u.btnClean.ConnectClicked(func(bool) { u.onCleanInvalidFile() })
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

	//增加选框支持
	u.picImage.ConnectMousePressEvent(u.onImageClick)
	u.picMarquee.ConnectMousePressEvent(u.onMarqueeClick)
	u.vVideo.ConnectMousePressEvent(u.onVideoClick)

	//设置边框
	u.picImage.SetFrameShape(widgets.QFrame__Box)
	u.picMarquee.SetFrameShape(widgets.QFrame__Box)
	u.vVideo.Resize2(200, 200)
	u.vVideo.Show()

	//
	u.videoplayer.ConnectError2(func(err multimedia.QMediaPlayer__Error) {
		log.Printf("Play video but recv err:%v\n", err)
	})
}

func (u *GameListUI) onDataNotify(m *parser.GameListItem) {
	u.leName.SetEnabled(false)
	u.leName.SetText(m.Name)
	u.lePath.SetText(m.Path)
	u.lePlayCnt.SetText(m.PlayCount)
	u.leLastPlay.SetText(m.LastPlayed)
	u.leLang.SetText(m.Lang)
	u.leRate.SetText(m.Rating)
	u.leReleaseDate.SetText(m.ReleaseDate)
	u.leDev.SetText(m.Developer)
	u.lePub.SetText(m.Publisher)
	u.leGenre.SetText(m.Genre)
	u.lePlayer.SetText(m.Players)
	u.leDesc.SetText(m.Desc)
	u.loadImage(m.Image)
	u.loadMarquee(m.Marquee)
	u.loadVideo(m.Video)
}

func (u *GameListUI) loadVideo(video string) {
	u.vVideo.SetToolTip(video)
	if len(video) != 0 {
		u.videoplayer.Stop()
		u.videoplayer.SetVideoOutput(u.vVideo)
		uri := fs.MergePath(u.gameloc, video)
		u.videoplayer.SetMedia(multimedia.NewQMediaContent2(
			core.QUrl_FromLocalFile(uri)), nil)
		u.videoplayer.Play()
		u.vVideo.AdjustSize()
	}
}

func (u *GameListUI) loadImage(img string) {
	u.picImage.SetToolTip(img)
	if len(img) != 0 {
		qp := gui.NewQPixmap3(fs.MergePath(u.gameloc, img), "", 0)
		qp = qp.Scaled2(200, 200, core.Qt__KeepAspectRatio, core.Qt__FastTransformation)
		u.picImage.SetPixmap(qp)
	}
}

func (u *GameListUI) loadMarquee(img string) {
	u.picMarquee.SetToolTip(img)
	if len(img) != 0 {
		qp := gui.NewQPixmap3(fs.MergePath(u.gameloc, img), "", 0)
		qp = qp.Scaled2(200, 200, core.Qt__KeepAspectRatio, core.Qt__FastTransformation)
		u.picMarquee.SetBaseSize2(100, 100)
		u.picMarquee.SetPixmap(qp)
	}
}

func (u *GameListUI) onListItemSelect() {

	name := u.lstGame.CurrentItem().Text()
	item, found := u.glp.Get(name)
	if !found {
		NoticeMessagef("未找到游戏:%s", name)
		return
	}
	u.onDataNotify(item)
	u.btnSave.SetEnabled(false)
	u.btnDiscard.SetEnabled(false)
	u.btnDelete.SetEnabled(true)

}

func (u *GameListUI) onModify(string) {
	u.btnSave.SetEnabled(true)
	u.btnDiscard.SetEnabled(true)
}

func (u *GameListUI) onDataSave() {
	//
	item, found := u.glp.Get(u.leName.Text())
	if !found {
		item = &parser.GameListItem{}
	}
	item.Name = u.leName.Text()
	item.Path = u.lePath.Text()
	item.PlayCount = u.lePlayCnt.Text()
	item.LastPlayed = u.leLastPlay.Text()
	item.Lang = u.leLang.Text()
	item.Desc = u.leDesc.ToPlainText()
	item.Rating = u.leRate.Text()
	item.ReleaseDate = u.leReleaseDate.Text()
	item.Developer = u.leDev.Text()
	item.Publisher = u.lePub.Text()
	item.Genre = u.leGenre.Text()
	item.Players = u.lePlayer.Text()
	//TODO:
	item.Image = u.picImage.ToolTip()
	item.Marquee = u.picMarquee.ToolTip()
	item.Video = u.vVideo.ToolTip()

	//
	u.glp.Set(item)

	//重新设置状态
	u.btnSave.SetEnabled(false)
	u.btnDiscard.SetEnabled(false)
	u.btnWriteFile.SetEnabled(true)
}

func (u *GameListUI) findIndex(name string) int {
	for index := 0; index < u.lstGame.Count(); index++ {
		if u.lstGame.Item(index).Text() == name {
			return index
		}
	}
	return -1
}

func (u *GameListUI) onDataDiscard() {
	item, found := u.glp.Get(u.leName.Text())
	if !found {
		NoticeMessagef("游戏:%s 不存在", u.leName.Text())
		return
	}
	if len(item.Path) == 0 { //不存在路径, 那么这个是要直接删除的
		u.glp.Remove(item.Name)
		index := u.findIndex(item.Name)
		if index != -1 {
			u.lstGame.TakeItem(index)
		}
	} else {
		u.onDataNotify(item)
	}
	//重新设置状态
	u.btnSave.SetEnabled(false)
	u.btnDiscard.SetEnabled(false)
}

func (u *GameListUI) onDataCreate() {
	//
	var dialog = widgets.NewQInputDialog(nil, core.Qt__Dialog)
	dialog.SetWindowTitle("Game")
	dialog.SetLabelText("输入游戏名:")
	dialog.SetTextEchoMode(widgets.QLineEdit__Normal)
	dialog.SetInputMethodHints(core.Qt__ImhNone)
	var data string
	var ok bool = false
	dialog.ConnectAccept(func() {
		data = dialog.TextValue()
		ok = true
		dialog.AcceptDefault()
	})
	dialog.Exec()

	if !ok {
		return
	}
	u.btnSave.SetEnabled(true)
	u.lstGame.AddItem(data)
	g := &parser.GameListItem{Name: data}
	u.glp.Set(g)
	u.lstGame.SetCurrentRow(u.lstGame.Count() - 1)
	u.btnDiscard.SetEnabled(true)
}

func (u *GameListUI) onDataDelete() {
	name := u.leName.Text()
	u.glp.Remove(name)
	u.lstGame.TakeItem(u.lstGame.CurrentIndex().Row())
	u.btnWriteFile.SetEnabled(true)
}

func (u *GameListUI) buildList() {
	for _, name := range u.glp.GetList() {
		u.lstGame.AddItem(name)
	}
	u.lstGame.SetCurrentIndex(core.NewQModelIndex())
}

func (u *GameListUI) onWriteToFile() {
	err := u.glp.Save()
	u.btnWriteFile.SetEnabled(false)
	if err != nil {
		NoticeMessagef("保存配置到文件失败, 错误信息:%v", err)
		return
	}
}

func (u *GameListUI) onPathClick(*gui.QMouseEvent) {
	file := widgets.QFileDialog_GetOpenFileName(nil, "选择游戏文件:",
		u.gameloc, "*.*", "*.*", widgets.QFileDialog__ReadOnly)
	if len(file) == 0 {
		return
	}
	ok := fs.IsParentDir(u.gameloc, file)
	if !ok {
		NoticeMessagef("游戏:%s 不存在roms目录中, 请手动复制进去。", file)
		return
	}
	sub := "./" + fs.TrimPath(u.gameloc, file)
	u.lePath.SetText(sub)
	u.onModify("")
}

func (u *GameListUI) onImageClick(*gui.QMouseEvent) {
	file := widgets.QFileDialog_GetOpenFileName(nil, "选择封面文件:",
		u.gameloc, "*.jpg *.png *.jpeg  *.gif", "*.*", widgets.QFileDialog__ReadOnly)
	if len(file) == 0 {
		return
	}
	ok := fs.IsParentDir(u.gameloc, file)
	if !ok {
		NoticeMessagef("图片:%s 不存在roms目录中, 请手动复制进去。", file)
		return
	}
	sub := "./" + fs.TrimPath(u.gameloc, file)
	u.loadImage(sub)
	u.onModify("")
}

func (u *GameListUI) onMarqueeClick(*gui.QMouseEvent) {
	file := widgets.QFileDialog_GetOpenFileName(nil, "选择封面文件:",
		u.gameloc, "*.jpg *.png *.jpeg  *.gif", "*.*", widgets.QFileDialog__ReadOnly)
	if len(file) == 0 {
		return
	}
	ok := fs.IsParentDir(u.gameloc, file)
	if !ok {
		NoticeMessagef("图片:%s 不存在roms目录中, 请手动复制进去。", file)
		return
	}
	sub := "./" + fs.TrimPath(u.gameloc, file)
	u.loadMarquee(sub)
	u.onModify("")
}

func (u *GameListUI) onVideoClick(*gui.QMouseEvent) {
	file := widgets.QFileDialog_GetOpenFileName(nil, "选择封面文件:",
		u.gameloc, "*.mp4 *.wmv", "*.*", widgets.QFileDialog__ReadOnly)
	if len(file) == 0 {
		return
	}
	ok := fs.IsParentDir(u.gameloc, file)
	if !ok {
		NoticeMessagef("视频:%s 不存在roms目录中, 请手动复制进去。", file)
		return
	}
	sub := "./" + fs.TrimPath(u.gameloc, file)
	u.loadVideo(sub)
	u.onModify("")
}

func (u *GameListUI) onCleanInvalidFile() {
	//正常来说, 目录下会有很多没有关联的rom、图片文件
	//这个功能就是来清理掉这些异常文件的

	//需要删除的项(gamelist.xml)
	var needRemoveItem []string
	//需要删除的文件
	var needRemoveFile []string
	//有效的扩展
	var validExt = make(map[string]bool)

	//被使用的文件, 包括rom, image, video
	var usedUrl = make(map[string]bool)

	//1. 检查gamelist.xml的roms文件是否存在
	mp := u.glp.GetAll()
	for _, item := range mp {
		//通过配置的rom文件自动探测支持的扩展
		//探测rom
		{
			ext := strings.ToLower(filepath.Ext(item.Path))
			if len(ext) != 0 {
				validExt[ext] = true
			}
		}
		//探测image
		{
			ext := strings.ToLower(filepath.Ext(item.Image))
			if len(ext) != 0 {
				validExt[ext] = true
			}
		}
		//探测marquee
		{
			ext := strings.ToLower(filepath.Ext(item.Marquee))
			if len(ext) != 0 {
				validExt[ext] = true
			}
		}
		//探测video
		{
			ext := strings.ToLower(filepath.Ext(item.Video))
			if len(ext) != 0 {
				validExt[ext] = true
			}
		}
		//检查文件是否存在
		loc := fs.MergePath(u.gameloc, item.Path)
		exist, err := fs.IsExist(loc)
		if err != nil {
			log.Printf("Read loc:%s fail, err:%v\n", loc, err)
			continue
		}
		if !exist {
			needRemoveItem = append(needRemoveItem, item.Name)
			continue
		}
		//构建当前项用到的所有路径
		{
			fmtPath := strings.Join(fs.SplitPath(item.Path), "/")
			fmtImage := strings.Join(fs.SplitPath(item.Image), "/")
			fmtMarquee := strings.Join(fs.SplitPath(item.Marquee), "/")
			fmtVideo := strings.Join(fs.SplitPath(item.Video), "/")
			if len(fmtPath) != 0 {
				usedUrl[fmtPath] = true
			}
			if len(fmtImage) != 0 {
				usedUrl[fmtImage] = true
			}
			if len(fmtMarquee) != 0 {
				usedUrl[fmtMarquee] = true
			}
			if len(fmtVideo) != 0 {
				usedUrl[fmtVideo] = true
			}
		}
	}
	//检查目录下的rom, 图片, 视频是否都在gamelist.xml中被引用
	fileMgr := mgr.NewFileMgr(u.gameloc)
	listFiles, err := fileMgr.Search(validExt)
	if err != nil {
		NoticeMessagef("读取文件列表失败, 错误信息:%v", err)
		return
	}
	for _, item := range listFiles {
		fmtPath := fs.TrimPath(u.gameloc, item)
		if _, ok := usedUrl[fmtPath]; !ok {
			log.Printf("Found url:%s not used, fmt:%s\n", item, fmtPath)
			needRemoveFile = append(needRemoveFile, fmtPath)
		}
	}
	//TODO: 执行清理

	//TODO:提供自动去重能力
	//TODO:提供自动添加不存在的gamelist.xml能力
}
