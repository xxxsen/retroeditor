// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIGamecleanDialog struct {
	GridLayout3 *widgets.QGridLayout
	GridLayout6 *widgets.QGridLayout
	Label *widgets.QLabel
	LeRomExt *widgets.QLineEdit
	GridLayout7 *widgets.QGridLayout
	Label2 *widgets.QLabel
	LePicExt *widgets.QLineEdit
	GridLayout8 *widgets.QGridLayout
	Label3 *widgets.QLabel
	LeVideoExt *widgets.QLineEdit
	GridLayout *widgets.QGridLayout
	LstClean *widgets.QListWidget
	CbCleanXml *widgets.QCheckBox
	GridLayout2 *widgets.QGridLayout
	LstAdd *widgets.QListWidget
	CbAddXml *widgets.QCheckBox
	GridLayout4 *widgets.QGridLayout
	LstCleanMedia *widgets.QListWidget
	CbCleanMedia *widgets.QCheckBox
	GridLayout9 *widgets.QGridLayout
	LstCleanRom *widgets.QListWidget
	CbCleanRom *widgets.QCheckBox
	GridLayout5 *widgets.QGridLayout
	BtnCancel *widgets.QPushButton
	BtnScan *widgets.QPushButton
	BtnExec *widgets.QPushButton
}

func (this *UIGamecleanDialog) SetupUI(Dialog *widgets.QDialog) {
	Dialog.SetObjectName("Dialog")
	Dialog.SetGeometry(core.NewQRect4(0, 0, 958, 600))
	this.GridLayout3 = widgets.NewQGridLayout(Dialog)
	this.GridLayout3.SetObjectName("gridLayout_3")
	this.GridLayout3.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout3.SetSpacing(0)
	this.GridLayout6 = widgets.NewQGridLayout(Dialog)
	this.GridLayout6.SetObjectName("gridLayout_6")
	this.GridLayout6.SetContentsMargins(8, 8, 8, 8)
	this.GridLayout6.SetSpacing(0)
	this.Label = widgets.NewQLabel(Dialog, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.GridLayout6.AddWidget3(this.Label, 0, 0, 1, 1, 0)
	this.LeRomExt = widgets.NewQLineEdit(Dialog)
	this.LeRomExt.SetObjectName("LeRomExt")
	this.GridLayout6.AddWidget3(this.LeRomExt, 0, 1, 1, 1, 0)
	this.GridLayout3.AddLayout2(this.GridLayout6, 0, 0, 1, 2, 0)
	this.GridLayout7 = widgets.NewQGridLayout(Dialog)
	this.GridLayout7.SetObjectName("gridLayout_7")
	this.GridLayout7.SetContentsMargins(8, 8, 8, 8)
	this.GridLayout7.SetSpacing(0)
	this.Label2 = widgets.NewQLabel(Dialog, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.GridLayout7.AddWidget3(this.Label2, 0, 0, 1, 1, 0)
	this.LePicExt = widgets.NewQLineEdit(Dialog)
	this.LePicExt.SetObjectName("LePicExt")
	this.GridLayout7.AddWidget3(this.LePicExt, 0, 1, 1, 1, 0)
	this.GridLayout3.AddLayout2(this.GridLayout7, 0, 2, 1, 2, 0)
	this.GridLayout8 = widgets.NewQGridLayout(Dialog)
	this.GridLayout8.SetObjectName("gridLayout_8")
	this.GridLayout8.SetContentsMargins(8, 8, 8, 8)
	this.GridLayout8.SetSpacing(0)
	this.Label3 = widgets.NewQLabel(Dialog, core.Qt__Widget)
	this.Label3.SetObjectName("Label3")
	this.GridLayout8.AddWidget3(this.Label3, 0, 0, 1, 1, 0)
	this.LeVideoExt = widgets.NewQLineEdit(Dialog)
	this.LeVideoExt.SetObjectName("LeVideoExt")
	this.GridLayout8.AddWidget3(this.LeVideoExt, 0, 1, 1, 1, 0)
	this.GridLayout3.AddLayout2(this.GridLayout8, 0, 4, 1, 2, 0)
	this.GridLayout = widgets.NewQGridLayout(Dialog)
	this.GridLayout.SetObjectName("gridLayout")
	this.GridLayout.SetContentsMargins(8, 0, 8, 0)
	this.GridLayout.SetSpacing(0)
	this.LstClean = widgets.NewQListWidget(Dialog)
	this.LstClean.SetObjectName("LstClean")
	this.LstClean.SetMinimumSize(core.NewQSize2(200, 0))
	this.GridLayout.AddWidget3(this.LstClean, 1, 0, 1, 1, 0)
	this.CbCleanXml = widgets.NewQCheckBox(Dialog)
	this.CbCleanXml.SetObjectName("CbCleanXml")
	this.CbCleanXml.SetChecked(true)
	this.CbCleanXml.SetTristate(false)
	this.GridLayout.AddWidget3(this.CbCleanXml, 0, 0, 1, 1, 0)
	this.GridLayout3.AddLayout2(this.GridLayout, 1, 0, 1, 1, 0)
	this.GridLayout2 = widgets.NewQGridLayout(Dialog)
	this.GridLayout2.SetObjectName("gridLayout_2")
	this.GridLayout2.SetContentsMargins(0, 0, 8, 0)
	this.GridLayout2.SetSpacing(0)
	this.LstAdd = widgets.NewQListWidget(Dialog)
	this.LstAdd.SetObjectName("LstAdd")
	this.LstAdd.SetMinimumSize(core.NewQSize2(200, 0))
	this.GridLayout2.AddWidget3(this.LstAdd, 1, 0, 1, 1, 0)
	this.CbAddXml = widgets.NewQCheckBox(Dialog)
	this.CbAddXml.SetObjectName("CbAddXml")
	this.CbAddXml.SetChecked(true)
	this.CbAddXml.SetTristate(false)
	this.GridLayout2.AddWidget3(this.CbAddXml, 0, 0, 1, 1, 0)
	this.GridLayout3.AddLayout2(this.GridLayout2, 1, 1, 1, 2, 0)
	this.GridLayout4 = widgets.NewQGridLayout(Dialog)
	this.GridLayout4.SetObjectName("gridLayout_4")
	this.GridLayout4.SetContentsMargins(0, 0, 8, 0)
	this.GridLayout4.SetSpacing(0)
	this.LstCleanMedia = widgets.NewQListWidget(Dialog)
	this.LstCleanMedia.SetObjectName("LstCleanMedia")
	this.LstCleanMedia.SetMinimumSize(core.NewQSize2(200, 0))
	this.GridLayout4.AddWidget3(this.LstCleanMedia, 1, 0, 1, 1, 0)
	this.CbCleanMedia = widgets.NewQCheckBox(Dialog)
	this.CbCleanMedia.SetObjectName("CbCleanMedia")
	this.CbCleanMedia.SetChecked(true)
	this.CbCleanMedia.SetTristate(false)
	this.GridLayout4.AddWidget3(this.CbCleanMedia, 0, 0, 1, 1, 0)
	this.GridLayout3.AddLayout2(this.GridLayout4, 1, 3, 1, 2, 0)
	this.GridLayout9 = widgets.NewQGridLayout(Dialog)
	this.GridLayout9.SetObjectName("gridLayout_9")
	this.GridLayout9.SetContentsMargins(0, 0, 8, 0)
	this.GridLayout9.SetSpacing(0)
	this.LstCleanRom = widgets.NewQListWidget(Dialog)
	this.LstCleanRom.SetObjectName("LstCleanRom")
	this.LstCleanRom.SetMinimumSize(core.NewQSize2(200, 0))
	this.GridLayout9.AddWidget3(this.LstCleanRom, 1, 0, 1, 1, 0)
	this.CbCleanRom = widgets.NewQCheckBox(Dialog)
	this.CbCleanRom.SetObjectName("CbCleanRom")
	this.CbCleanRom.SetChecked(false)
	this.CbCleanRom.SetTristate(false)
	this.GridLayout9.AddWidget3(this.CbCleanRom, 0, 0, 1, 1, 0)
	this.GridLayout3.AddLayout2(this.GridLayout9, 1, 5, 1, 1, 0)
	this.GridLayout5 = widgets.NewQGridLayout(Dialog)
	this.GridLayout5.SetObjectName("gridLayout_5")
	this.GridLayout5.SetContentsMargins(8, 8, 8, 8)
	this.GridLayout5.SetSpacing(0)
	this.BtnCancel = widgets.NewQPushButton(Dialog)
	this.BtnCancel.SetObjectName("BtnCancel")
	this.GridLayout5.AddWidget3(this.BtnCancel, 0, 0, 1, 1, 0)
	this.BtnScan = widgets.NewQPushButton(Dialog)
	this.BtnScan.SetObjectName("BtnScan")
	this.GridLayout5.AddWidget3(this.BtnScan, 0, 1, 1, 1, 0)
	this.BtnExec = widgets.NewQPushButton(Dialog)
	this.BtnExec.SetObjectName("BtnExec")
	this.GridLayout5.AddWidget3(this.BtnExec, 0, 2, 1, 1, 0)
	this.GridLayout3.AddLayout2(this.GridLayout5, 2, 0, 1, 6, 0)


    this.RetranslateUi(Dialog)

}

func (this *UIGamecleanDialog) RetranslateUi(Dialog *widgets.QDialog) {
    _translate := core.QCoreApplication_Translate
	Dialog.SetWindowTitle(_translate("Dialog", "Gamelist.xml重建", "", -1))
	this.Label.SetText(_translate("Dialog", "rom扩展名:", "", -1))
	this.Label2.SetText(_translate("Dialog", "图片扩展名:", "", -1))
	this.Label3.SetText(_translate("Dialog", "视频扩展名:", "", -1))
	this.CbCleanXml.SetText(_translate("Dialog", "GameList.xml清理项", "", -1))
	this.CbAddXml.SetText(_translate("Dialog", "GameList.xml添加项", "", -1))
	this.CbCleanMedia.SetText(_translate("Dialog", "图片/视频清理项", "", -1))
	this.CbCleanRom.SetText(_translate("Dialog", "未配置Rom清理", "", -1))
	this.BtnCancel.SetText(_translate("Dialog", "取消(&C)", "", -1))
	this.BtnScan.SetText(_translate("Dialog", "扫描(&S)", "", -1))
	this.BtnExec.SetText(_translate("Dialog", "执行(&C)", "", -1))
}
