// WARNING! All changes made in this file will be lost!
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIAppDialog struct {
	VerticalLayout *widgets.QVBoxLayout
	Gbox *widgets.QGroupBox
	FormLayout2 *widgets.QFormLayout
	FormLayout *widgets.QFormLayout
	LineEdit2 *widgets.QLineEdit
	Label *widgets.QLabel
	Label4 *widgets.QLabel
	LineEdit5 *widgets.QLineEdit
	Label2 *widgets.QLabel
	LineEdit3 *widgets.QLineEdit
	Label3 *widgets.QLabel
	LineEdit4 *widgets.QLineEdit
	Label5 *widgets.QLabel
	LineEdit6 *widgets.QLineEdit
	HorizontalLayout *widgets.QHBoxLayout
	Label6 *widgets.QLabel
	LineEdit *widgets.QLineEdit
	HorizontalLayout2 *widgets.QHBoxLayout
	CheckBox2 *widgets.QCheckBox
	CheckBox *widgets.QCheckBox
	HorizontalLayout3 *widgets.QHBoxLayout
	PushButton *widgets.QPushButton
	PushButton2 *widgets.QPushButton
}

func (this *UIAppDialog) SetupUI(Dialog *widgets.QDialog) {
	Dialog.SetObjectName("Dialog")
	Dialog.SetGeometry(core.NewQRect4(0, 0, 313, 407))
	this.VerticalLayout = widgets.NewQVBoxLayout2(Dialog)
	this.VerticalLayout.SetObjectName("verticalLayout")
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout.SetSpacing(0)
	this.Gbox = widgets.NewQGroupBox(Dialog)
	this.Gbox.SetObjectName("Gbox")
	this.FormLayout2 = widgets.NewQFormLayout(this.Gbox)
	this.FormLayout2.SetObjectName("formLayout_2")
	this.FormLayout2.SetContentsMargins(0, 0, 0, 0)
	this.FormLayout2.SetSpacing(0)
	this.FormLayout = widgets.NewQFormLayout(this.Gbox)
	this.FormLayout.SetObjectName("formLayout")
	this.FormLayout.SetContentsMargins(0, 0, 0, 0)
	this.FormLayout.SetSpacing(0)
	this.LineEdit2 = widgets.NewQLineEdit(this.Gbox)
	this.LineEdit2.SetObjectName("LineEdit2")
	this.FormLayout.SetWidget(0, widgets.QFormLayout__FieldRole, this.LineEdit2)
	this.Label = widgets.NewQLabel(this.Gbox, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.FormLayout.SetWidget(0, widgets.QFormLayout__LabelRole, this.Label)
	this.Label4 = widgets.NewQLabel(this.Gbox, core.Qt__Widget)
	this.Label4.SetObjectName("Label4")
	this.FormLayout.SetWidget(1, widgets.QFormLayout__LabelRole, this.Label4)
	this.LineEdit5 = widgets.NewQLineEdit(this.Gbox)
	this.LineEdit5.SetObjectName("LineEdit5")
	this.FormLayout.SetWidget(1, widgets.QFormLayout__FieldRole, this.LineEdit5)
	this.Label2 = widgets.NewQLabel(this.Gbox, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.FormLayout.SetWidget(2, widgets.QFormLayout__LabelRole, this.Label2)
	this.LineEdit3 = widgets.NewQLineEdit(this.Gbox)
	this.LineEdit3.SetObjectName("LineEdit3")
	this.FormLayout.SetWidget(2, widgets.QFormLayout__FieldRole, this.LineEdit3)
	this.Label3 = widgets.NewQLabel(this.Gbox, core.Qt__Widget)
	this.Label3.SetObjectName("Label3")
	this.FormLayout.SetWidget(3, widgets.QFormLayout__LabelRole, this.Label3)
	this.LineEdit4 = widgets.NewQLineEdit(this.Gbox)
	this.LineEdit4.SetObjectName("LineEdit4")
	this.FormLayout.SetWidget(3, widgets.QFormLayout__FieldRole, this.LineEdit4)
	this.Label5 = widgets.NewQLabel(this.Gbox, core.Qt__Widget)
	this.Label5.SetObjectName("Label5")
	this.FormLayout.SetWidget(4, widgets.QFormLayout__LabelRole, this.Label5)
	this.LineEdit6 = widgets.NewQLineEdit(this.Gbox)
	this.LineEdit6.SetObjectName("LineEdit6")
	this.FormLayout.SetWidget(4, widgets.QFormLayout__FieldRole, this.LineEdit6)
	this.FormLayout2.SetLayout(0, widgets.QFormLayout__LabelRole, this.FormLayout)
	this.VerticalLayout.AddWidget(this.Gbox, 0, 0)
	this.HorizontalLayout = widgets.NewQHBoxLayout2(Dialog)
	this.HorizontalLayout.SetObjectName("horizontalLayout")
	this.HorizontalLayout.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout.SetSpacing(0)
	this.Label6 = widgets.NewQLabel(Dialog, core.Qt__Widget)
	this.Label6.SetObjectName("Label6")
	this.HorizontalLayout.AddWidget(this.Label6, 0, 0)
	this.LineEdit = widgets.NewQLineEdit(Dialog)
	this.LineEdit.SetObjectName("LineEdit")
	this.HorizontalLayout.AddWidget(this.LineEdit, 0, 0)
	this.VerticalLayout.AddLayout(this.HorizontalLayout, 0)
	this.HorizontalLayout2 = widgets.NewQHBoxLayout2(Dialog)
	this.HorizontalLayout2.SetObjectName("horizontalLayout_2")
	this.HorizontalLayout2.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout2.SetSpacing(0)
	this.CheckBox2 = widgets.NewQCheckBox(Dialog)
	this.CheckBox2.SetObjectName("CheckBox2")
	this.HorizontalLayout2.AddWidget(this.CheckBox2, 0, 0)
	this.CheckBox = widgets.NewQCheckBox(Dialog)
	this.CheckBox.SetObjectName("CheckBox")
	this.HorizontalLayout2.AddWidget(this.CheckBox, 0, 0)
	this.VerticalLayout.AddLayout(this.HorizontalLayout2, 0)
	this.HorizontalLayout3 = widgets.NewQHBoxLayout2(Dialog)
	this.HorizontalLayout3.SetObjectName("horizontalLayout_3")
	this.HorizontalLayout3.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout3.SetSpacing(0)
	this.PushButton = widgets.NewQPushButton(Dialog)
	this.PushButton.SetObjectName("PushButton")
	this.PushButton.SetDefault(true)
	this.PushButton.SetFlat(false)
	this.HorizontalLayout3.AddWidget(this.PushButton, 0, 0)
	this.PushButton2 = widgets.NewQPushButton(Dialog)
	this.PushButton2.SetObjectName("PushButton2")
	this.HorizontalLayout3.AddWidget(this.PushButton2, 0, 0)
	this.VerticalLayout.AddLayout(this.HorizontalLayout3, 0)


    this.RetranslateUi(Dialog)

}

func (this *UIAppDialog) RetranslateUi(Dialog *widgets.QDialog) {
    _translate := core.QCoreApplication_Translate
	Dialog.SetWindowTitle(_translate("Dialog", "模型文件生成器", "", -1))
	this.Gbox.SetTitle(_translate("Dialog", "数据库信息", "", -1))
	this.Label.SetText(_translate("Dialog", "地址", "", -1))
	this.Label4.SetText(_translate("Dialog", "端口", "", -1))
	this.Label2.SetText(_translate("Dialog", "数据库", "", -1))
	this.Label3.SetText(_translate("Dialog", "用户", "", -1))
	this.Label5.SetText(_translate("Dialog", "密码", "", -1))
	this.Label6.SetText(_translate("Dialog", "包名", "", -1))
	this.CheckBox2.SetText(_translate("Dialog", "生成js文件", "", -1))
	this.CheckBox.SetText(_translate("Dialog", "生成Go文件", "", -1))
	this.PushButton.SetText(_translate("Dialog", "生成", "", -1))
	this.PushButton2.SetText(_translate("Dialog", "关闭", "", -1))
}
