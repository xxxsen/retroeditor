// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIGamelistForm struct {
	GridLayoutWidget4 *widgets.QWidget
	GridLayout4 *widgets.QGridLayout
	GridLayout2 *widgets.QGridLayout
	VVideo *widgets.QWidget
	Label15 *widgets.QLabel
	PicMarquee *widgets.QLabel
	Label13 *widgets.QLabel
	Label14 *widgets.QLabel
	PicImage *widgets.QLabel
	LstGame *widgets.QListWidget
	LayoutWidget *widgets.QWidget
	GridLayout7 *widgets.QGridLayout
	GridLayout5 *widgets.QGridLayout
	GridLayout *widgets.QGridLayout
	LeRate *widgets.QLineEdit
	LePlayer *widgets.QLineEdit
	Label4 *widgets.QLabel
	Label11 *widgets.QLabel
	Label *widgets.QLabel
	Label2 *widgets.QLabel
	LeLastPlay *widgets.QLineEdit
	LeDev *widgets.QLineEdit
	Label8 *widgets.QLabel
	Label10 *widgets.QLabel
	Label12 *widgets.QLabel
	Label3 *widgets.QLabel
	LeDesc *widgets.QTextEdit
	LePlayCnt *widgets.QLineEdit
	Label9 *widgets.QLabel
	Label6 *widgets.QLabel
	Label5 *widgets.QLabel
	LePub *widgets.QLineEdit
	LeReleaseDate *widgets.QLineEdit
	Label7 *widgets.QLabel
	LeLang *widgets.QLineEdit
	LeGenre *widgets.QLineEdit
	LeName *widgets.QLineEdit
	LePath *widgets.QLineEdit
	GridLayout3 *widgets.QGridLayout
	BtnDelete *widgets.QPushButton
	BtnMod *widgets.QPushButton
	BtnCreate *widgets.QPushButton
	BtnWriteFile *widgets.QPushButton
}

func (this *UIGamelistForm) SetupUI(Form *widgets.QWidget) {
	Form.SetObjectName("Form")
	Form.SetGeometry(core.NewQRect4(0, 0, 1280, 720))
	Form.SetMinimumSize(core.NewQSize2(1280, 720))
	this.GridLayoutWidget4 = widgets.NewQWidget(Form, core.Qt__Widget)
	this.GridLayoutWidget4.SetObjectName("GridLayoutWidget4")
	this.GridLayoutWidget4.SetGeometry(core.NewQRect4(0, 20, 848, 616))
	this.GridLayout4 = widgets.NewQGridLayout(this.GridLayoutWidget4)
	this.GridLayout4.SetObjectName("gridLayout_4")
	this.GridLayout4.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout4.SetSpacing(0)
	this.GridLayout2 = widgets.NewQGridLayout(this.GridLayoutWidget4)
	this.GridLayout2.SetObjectName("gridLayout_2")
	this.GridLayout2.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout2.SetSpacing(0)
	this.VVideo = widgets.NewQWidget(this.GridLayoutWidget4, core.Qt__Widget)
	this.VVideo.SetObjectName("VVideo")
	this.VVideo.SetMinimumSize(core.NewQSize2(200, 200))
	this.GridLayout2.AddWidget3(this.VVideo, 2, 1, 1, 1, 0)
	this.Label15 = widgets.NewQLabel(this.GridLayoutWidget4, core.Qt__Widget)
	this.Label15.SetObjectName("Label15")
	this.GridLayout2.AddWidget3(this.Label15, 2, 0, 1, 1, 0)
	this.PicMarquee = widgets.NewQLabel(this.GridLayoutWidget4, core.Qt__Widget)
	this.PicMarquee.SetObjectName("PicMarquee")
	this.PicMarquee.SetMinimumSize(core.NewQSize2(200, 200))
	this.GridLayout2.AddWidget3(this.PicMarquee, 1, 1, 1, 1, 0)
	this.Label13 = widgets.NewQLabel(this.GridLayoutWidget4, core.Qt__Widget)
	this.Label13.SetObjectName("Label13")
	this.GridLayout2.AddWidget3(this.Label13, 0, 0, 1, 1, 0)
	this.Label14 = widgets.NewQLabel(this.GridLayoutWidget4, core.Qt__Widget)
	this.Label14.SetObjectName("Label14")
	this.GridLayout2.AddWidget3(this.Label14, 1, 0, 1, 1, 0)
	this.PicImage = widgets.NewQLabel(this.GridLayoutWidget4, core.Qt__Widget)
	this.PicImage.SetObjectName("PicImage")
	this.PicImage.SetMinimumSize(core.NewQSize2(200, 200))
	this.GridLayout2.AddWidget3(this.PicImage, 0, 1, 1, 1, 0)
	this.GridLayout4.AddLayout2(this.GridLayout2, 0, 2, 1, 1, 0)
	this.LstGame = widgets.NewQListWidget(this.GridLayoutWidget4)
	this.LstGame.SetObjectName("LstGame")
	this.GridLayout4.AddWidget3(this.LstGame, 0, 1, 1, 1, 0)
	this.LayoutWidget = widgets.NewQWidget(Form, core.Qt__Widget)
	this.LayoutWidget.SetObjectName("LayoutWidget")
	this.LayoutWidget.SetGeometry(core.NewQRect4(910, 30, 346, 584))
	this.GridLayout7 = widgets.NewQGridLayout(this.LayoutWidget)
	this.GridLayout7.SetObjectName("gridLayout_7")
	this.GridLayout7.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout7.SetSpacing(0)
	this.GridLayout5 = widgets.NewQGridLayout(this.LayoutWidget)
	this.GridLayout5.SetObjectName("gridLayout_5")
	this.GridLayout5.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout5.SetSpacing(0)
	this.GridLayout = widgets.NewQGridLayout(this.LayoutWidget)
	this.GridLayout.SetObjectName("gridLayout")
	this.GridLayout.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout.SetSpacing(0)
	this.LeRate = widgets.NewQLineEdit(this.LayoutWidget)
	this.LeRate.SetObjectName("LeRate")
	this.GridLayout.AddWidget3(this.LeRate, 6, 3, 1, 1, 0)
	this.LePlayer = widgets.NewQLineEdit(this.LayoutWidget)
	this.LePlayer.SetObjectName("LePlayer")
	this.GridLayout.AddWidget3(this.LePlayer, 11, 3, 1, 1, 0)
	this.Label4 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label4.SetObjectName("Label4")
	this.GridLayout.AddWidget3(this.Label4, 4, 1, 1, 1, 0)
	this.Label11 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label11.SetObjectName("Label11")
	this.GridLayout.AddWidget3(this.Label11, 11, 1, 1, 1, 0)
	this.Label = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.GridLayout.AddWidget3(this.Label, 1, 1, 1, 1, 0)
	this.Label2 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.GridLayout.AddWidget3(this.Label2, 0, 1, 1, 1, 0)
	this.LeLastPlay = widgets.NewQLineEdit(this.LayoutWidget)
	this.LeLastPlay.SetObjectName("LeLastPlay")
	this.GridLayout.AddWidget3(this.LeLastPlay, 4, 3, 1, 1, 0)
	this.LeDev = widgets.NewQLineEdit(this.LayoutWidget)
	this.LeDev.SetObjectName("LeDev")
	this.GridLayout.AddWidget3(this.LeDev, 8, 3, 1, 1, 0)
	this.Label8 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label8.SetObjectName("Label8")
	this.GridLayout.AddWidget3(this.Label8, 8, 1, 1, 1, 0)
	this.Label10 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label10.SetObjectName("Label10")
	this.GridLayout.AddWidget3(this.Label10, 10, 1, 1, 1, 0)
	this.Label12 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label12.SetObjectName("Label12")
	this.GridLayout.AddWidget3(this.Label12, 12, 1, 1, 1, 0)
	this.Label3 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label3.SetObjectName("Label3")
	this.GridLayout.AddWidget3(this.Label3, 3, 1, 1, 1, 0)
	this.LeDesc = widgets.NewQTextEdit(this.LayoutWidget)
	this.LeDesc.SetObjectName("LeDesc")
	this.GridLayout.AddWidget3(this.LeDesc, 12, 3, 1, 1, 0)
	this.LePlayCnt = widgets.NewQLineEdit(this.LayoutWidget)
	this.LePlayCnt.SetObjectName("LePlayCnt")
	this.GridLayout.AddWidget3(this.LePlayCnt, 3, 3, 1, 1, 0)
	this.Label9 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label9.SetObjectName("Label9")
	this.GridLayout.AddWidget3(this.Label9, 9, 1, 1, 1, 0)
	this.Label6 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label6.SetObjectName("Label6")
	this.GridLayout.AddWidget3(this.Label6, 6, 1, 1, 1, 0)
	this.Label5 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label5.SetObjectName("Label5")
	this.GridLayout.AddWidget3(this.Label5, 5, 1, 1, 1, 0)
	this.LePub = widgets.NewQLineEdit(this.LayoutWidget)
	this.LePub.SetObjectName("LePub")
	this.GridLayout.AddWidget3(this.LePub, 9, 3, 1, 1, 0)
	this.LeReleaseDate = widgets.NewQLineEdit(this.LayoutWidget)
	this.LeReleaseDate.SetObjectName("LeReleaseDate")
	this.GridLayout.AddWidget3(this.LeReleaseDate, 7, 3, 1, 1, 0)
	this.Label7 = widgets.NewQLabel(this.LayoutWidget, core.Qt__Widget)
	this.Label7.SetObjectName("Label7")
	this.GridLayout.AddWidget3(this.Label7, 7, 1, 1, 1, 0)
	this.LeLang = widgets.NewQLineEdit(this.LayoutWidget)
	this.LeLang.SetObjectName("LeLang")
	this.GridLayout.AddWidget3(this.LeLang, 5, 3, 1, 1, 0)
	this.LeGenre = widgets.NewQLineEdit(this.LayoutWidget)
	this.LeGenre.SetObjectName("LeGenre")
	this.GridLayout.AddWidget3(this.LeGenre, 10, 3, 1, 1, 0)
	this.LeName = widgets.NewQLineEdit(this.LayoutWidget)
	this.LeName.SetObjectName("LeName")
	this.GridLayout.AddWidget3(this.LeName, 1, 3, 1, 1, 0)
	this.LePath = widgets.NewQLineEdit(this.LayoutWidget)
	this.LePath.SetObjectName("LePath")
	this.LePath.SetEnabled(false)
	this.GridLayout.AddWidget3(this.LePath, 0, 3, 1, 1, 0)
	this.GridLayout5.AddLayout2(this.GridLayout, 0, 0, 1, 1, 0)
	this.GridLayout3 = widgets.NewQGridLayout(this.LayoutWidget)
	this.GridLayout3.SetObjectName("gridLayout_3")
	this.GridLayout3.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout3.SetSpacing(0)
	this.BtnDelete = widgets.NewQPushButton(this.LayoutWidget)
	this.BtnDelete.SetObjectName("BtnDelete")
	this.GridLayout3.AddWidget3(this.BtnDelete, 0, 0, 1, 1, 0)
	this.BtnMod = widgets.NewQPushButton(this.LayoutWidget)
	this.BtnMod.SetObjectName("BtnMod")
	this.GridLayout3.AddWidget3(this.BtnMod, 0, 2, 1, 1, 0)
	this.BtnCreate = widgets.NewQPushButton(this.LayoutWidget)
	this.BtnCreate.SetObjectName("BtnCreate")
	this.GridLayout3.AddWidget3(this.BtnCreate, 0, 1, 1, 1, 0)
	this.BtnWriteFile = widgets.NewQPushButton(this.LayoutWidget)
	this.BtnWriteFile.SetObjectName("BtnWriteFile")
	this.GridLayout3.AddWidget3(this.BtnWriteFile, 0, 3, 1, 1, 0)
	this.GridLayout5.AddLayout2(this.GridLayout3, 1, 0, 1, 1, 0)
	this.GridLayout7.AddLayout2(this.GridLayout5, 0, 0, 1, 1, 0)


    this.RetranslateUi(Form)

}

func (this *UIGamelistForm) RetranslateUi(Form *widgets.QWidget) {
    _translate := core.QCoreApplication_Translate
	Form.SetWindowTitle(_translate("Form", "Form", "", -1))
	this.Label15.SetText(_translate("Form", "预览", "", -1))
	this.PicMarquee.SetText(_translate("Form", "", "", -1))
	this.Label13.SetText(_translate("Form", "封面", "", -1))
	this.Label14.SetText(_translate("Form", "相框", "", -1))
	this.PicImage.SetText(_translate("Form", "", "", -1))
	this.Label4.SetText(_translate("Form", "最后游玩时间", "", -1))
	this.Label11.SetText(_translate("Form", "支持人数", "", -1))
	this.Label.SetText(_translate("Form", "游戏名:", "", -1))
	this.Label2.SetText(_translate("Form", "游戏路径", "", -1))
	this.Label8.SetText(_translate("Form", "开发商", "", -1))
	this.Label10.SetText(_translate("Form", "游戏类型", "", -1))
	this.Label12.SetText(_translate("Form", "描述", "", -1))
	this.Label3.SetText(_translate("Form", "游玩次数", "", -1))
	this.Label9.SetText(_translate("Form", "发行商", "", -1))
	this.Label6.SetText(_translate("Form", "好评率", "", -1))
	this.Label5.SetText(_translate("Form", "语言", "", -1))
	this.Label7.SetText(_translate("Form", "发行时间", "", -1))
	this.BtnDelete.SetText(_translate("Form", "删除游戏(&D)", "", -1))
	this.BtnMod.SetText(_translate("Form", "修改游戏(&S)", "", -1))
	this.BtnCreate.SetText(_translate("Form", "创建游戏(&N)", "", -1))
	this.BtnWriteFile.SetText(_translate("Form", "写入文件(&W)", "", -1))
}
