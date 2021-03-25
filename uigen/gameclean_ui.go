// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIGamecleanDialog struct {
	GridLayoutWidget *widgets.QWidget
	GridLayout *widgets.QGridLayout
	LstClean *widgets.QListWidget
	CbCleanXml *widgets.QCheckBox
	GridLayoutWidget2 *widgets.QWidget
	GridLayout2 *widgets.QGridLayout
	LstAdd *widgets.QListWidget
	CbAddXml *widgets.QCheckBox
	GridLayoutWidget4 *widgets.QWidget
	GridLayout4 *widgets.QGridLayout
	LstCleanMedia *widgets.QListWidget
	CbCleanMedia *widgets.QCheckBox
	GridLayoutWidget5 *widgets.QWidget
	GridLayout5 *widgets.QGridLayout
	BtnCancel *widgets.QPushButton
	BtnScan *widgets.QPushButton
	BtnExec *widgets.QPushButton
	GridLayoutWidget6 *widgets.QWidget
	GridLayout6 *widgets.QGridLayout
	Label *widgets.QLabel
	LeRomExt *widgets.QLineEdit
	GridLayoutWidget7 *widgets.QWidget
	GridLayout7 *widgets.QGridLayout
	Label2 *widgets.QLabel
	LePicExt *widgets.QLineEdit
	GridLayoutWidget8 *widgets.QWidget
	GridLayout8 *widgets.QGridLayout
	Label3 *widgets.QLabel
	LeVideoExt *widgets.QLineEdit
}

func (this *UIGamecleanDialog) SetupUI(Dialog *widgets.QDialog) {
	Dialog.SetObjectName("Dialog")
	Dialog.SetGeometry(core.NewQRect4(0, 0, 958, 600))
	this.GridLayoutWidget = widgets.NewQWidget(Dialog, core.Qt__Widget)
	this.GridLayoutWidget.SetObjectName("GridLayoutWidget")
	this.GridLayoutWidget.SetGeometry(core.NewQRect4(20, 110, 202, 391))
	this.GridLayout = widgets.NewQGridLayout(this.GridLayoutWidget)
	this.GridLayout.SetObjectName("gridLayout")
	this.GridLayout.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout.SetSpacing(0)
	this.LstClean = widgets.NewQListWidget(this.GridLayoutWidget)
	this.LstClean.SetObjectName("LstClean")
	this.LstClean.SetMinimumSize(core.NewQSize2(200, 0))
	this.GridLayout.AddWidget3(this.LstClean, 1, 0, 1, 1, 0)
	this.CbCleanXml = widgets.NewQCheckBox(this.GridLayoutWidget)
	this.CbCleanXml.SetObjectName("CbCleanXml")
	this.GridLayout.AddWidget3(this.CbCleanXml, 0, 0, 1, 1, 0)
	this.GridLayoutWidget2 = widgets.NewQWidget(Dialog, core.Qt__Widget)
	this.GridLayoutWidget2.SetObjectName("GridLayoutWidget2")
	this.GridLayoutWidget2.SetGeometry(core.NewQRect4(240, 110, 202, 391))
	this.GridLayout2 = widgets.NewQGridLayout(this.GridLayoutWidget2)
	this.GridLayout2.SetObjectName("gridLayout_2")
	this.GridLayout2.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout2.SetSpacing(0)
	this.LstAdd = widgets.NewQListWidget(this.GridLayoutWidget2)
	this.LstAdd.SetObjectName("LstAdd")
	this.LstAdd.SetMinimumSize(core.NewQSize2(200, 0))
	this.GridLayout2.AddWidget3(this.LstAdd, 1, 0, 1, 1, 0)
	this.CbAddXml = widgets.NewQCheckBox(this.GridLayoutWidget2)
	this.CbAddXml.SetObjectName("CbAddXml")
	this.GridLayout2.AddWidget3(this.CbAddXml, 0, 0, 1, 1, 0)
	this.GridLayoutWidget4 = widgets.NewQWidget(Dialog, core.Qt__Widget)
	this.GridLayoutWidget4.SetObjectName("GridLayoutWidget4")
	this.GridLayoutWidget4.SetGeometry(core.NewQRect4(540, 110, 202, 391))
	this.GridLayout4 = widgets.NewQGridLayout(this.GridLayoutWidget4)
	this.GridLayout4.SetObjectName("gridLayout_4")
	this.GridLayout4.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout4.SetSpacing(0)
	this.LstCleanMedia = widgets.NewQListWidget(this.GridLayoutWidget4)
	this.LstCleanMedia.SetObjectName("LstCleanMedia")
	this.LstCleanMedia.SetMinimumSize(core.NewQSize2(200, 0))
	this.GridLayout4.AddWidget3(this.LstCleanMedia, 1, 0, 1, 1, 0)
	this.CbCleanMedia = widgets.NewQCheckBox(this.GridLayoutWidget4)
	this.CbCleanMedia.SetObjectName("CbCleanMedia")
	this.GridLayout4.AddWidget3(this.CbCleanMedia, 0, 0, 1, 1, 0)
	this.GridLayoutWidget5 = widgets.NewQWidget(Dialog, core.Qt__Widget)
	this.GridLayoutWidget5.SetObjectName("GridLayoutWidget5")
	this.GridLayoutWidget5.SetGeometry(core.NewQRect4(150, 520, 421, 80))
	this.GridLayout5 = widgets.NewQGridLayout(this.GridLayoutWidget5)
	this.GridLayout5.SetObjectName("gridLayout_5")
	this.GridLayout5.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout5.SetSpacing(0)
	this.BtnCancel = widgets.NewQPushButton(this.GridLayoutWidget5)
	this.BtnCancel.SetObjectName("BtnCancel")
	this.GridLayout5.AddWidget3(this.BtnCancel, 0, 0, 1, 1, 0)
	this.BtnScan = widgets.NewQPushButton(this.GridLayoutWidget5)
	this.BtnScan.SetObjectName("BtnScan")
	this.GridLayout5.AddWidget3(this.BtnScan, 0, 1, 1, 1, 0)
	this.BtnExec = widgets.NewQPushButton(this.GridLayoutWidget5)
	this.BtnExec.SetObjectName("BtnExec")
	this.GridLayout5.AddWidget3(this.BtnExec, 0, 2, 1, 1, 0)
	this.GridLayoutWidget6 = widgets.NewQWidget(Dialog, core.Qt__Widget)
	this.GridLayoutWidget6.SetObjectName("GridLayoutWidget6")
	this.GridLayoutWidget6.SetGeometry(core.NewQRect4(50, 30, 221, 51))
	this.GridLayout6 = widgets.NewQGridLayout(this.GridLayoutWidget6)
	this.GridLayout6.SetObjectName("gridLayout_6")
	this.GridLayout6.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout6.SetSpacing(0)
	this.Label = widgets.NewQLabel(this.GridLayoutWidget6, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.GridLayout6.AddWidget3(this.Label, 0, 0, 1, 1, 0)
	this.LeRomExt = widgets.NewQLineEdit(this.GridLayoutWidget6)
	this.LeRomExt.SetObjectName("LeRomExt")
	this.GridLayout6.AddWidget3(this.LeRomExt, 0, 1, 1, 1, 0)
	this.GridLayoutWidget7 = widgets.NewQWidget(Dialog, core.Qt__Widget)
	this.GridLayoutWidget7.SetObjectName("GridLayoutWidget7")
	this.GridLayoutWidget7.SetGeometry(core.NewQRect4(300, 30, 211, 51))
	this.GridLayout7 = widgets.NewQGridLayout(this.GridLayoutWidget7)
	this.GridLayout7.SetObjectName("gridLayout_7")
	this.GridLayout7.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout7.SetSpacing(0)
	this.Label2 = widgets.NewQLabel(this.GridLayoutWidget7, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.GridLayout7.AddWidget3(this.Label2, 0, 0, 1, 1, 0)
	this.LePicExt = widgets.NewQLineEdit(this.GridLayoutWidget7)
	this.LePicExt.SetObjectName("LePicExt")
	this.GridLayout7.AddWidget3(this.LePicExt, 0, 1, 1, 1, 0)
	this.GridLayoutWidget8 = widgets.NewQWidget(Dialog, core.Qt__Widget)
	this.GridLayoutWidget8.SetObjectName("GridLayoutWidget8")
	this.GridLayoutWidget8.SetGeometry(core.NewQRect4(530, 30, 231, 51))
	this.GridLayout8 = widgets.NewQGridLayout(this.GridLayoutWidget8)
	this.GridLayout8.SetObjectName("gridLayout_8")
	this.GridLayout8.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout8.SetSpacing(0)
	this.Label3 = widgets.NewQLabel(this.GridLayoutWidget8, core.Qt__Widget)
	this.Label3.SetObjectName("Label3")
	this.GridLayout8.AddWidget3(this.Label3, 0, 0, 1, 1, 0)
	this.LeVideoExt = widgets.NewQLineEdit(this.GridLayoutWidget8)
	this.LeVideoExt.SetObjectName("LeVideoExt")
	this.GridLayout8.AddWidget3(this.LeVideoExt, 0, 1, 1, 1, 0)


    this.RetranslateUi(Dialog)

}

func (this *UIGamecleanDialog) RetranslateUi(Dialog *widgets.QDialog) {
    _translate := core.QCoreApplication_Translate
	Dialog.SetWindowTitle(_translate("Dialog", "Gamelist.xml重建", "", -1))
	this.CbCleanXml.SetText(_translate("Dialog", "GameList.xml清理项", "", -1))
	this.CbAddXml.SetText(_translate("Dialog", "GameList.xml添加项", "", -1))
	this.CbCleanMedia.SetText(_translate("Dialog", "图片/视频清理项", "", -1))
	this.BtnCancel.SetText(_translate("Dialog", "取消(&C)", "", -1))
	this.BtnScan.SetText(_translate("Dialog", "扫描(&S)", "", -1))
	this.BtnExec.SetText(_translate("Dialog", "执行(&C)", "", -1))
	this.Label.SetText(_translate("Dialog", "rom扩展名", "", -1))
	this.Label2.SetText(_translate("Dialog", "图片扩展名", "", -1))
	this.Label3.SetText(_translate("Dialog", "视频扩展名", "", -1))
}
