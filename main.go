package main

import (
	"github.com/oxiao/gentity/biz"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	dlg := biz.NewMainWindow(nil, 0)
	dlg.Show()

	//// left sider
	//splitterLeft := widgets.NewQSplitter2(core.Qt__Horizontal, nil)
	//textTop := widgets.NewQTextEdit2("左部文本", splitterLeft)
	//splitterLeft.AddWidget(textTop)
	//
	//// right sider
	//splitterRight := widgets.NewQSplitter2(core.Qt__Vertical, splitterLeft)
	//textRight := widgets.NewQTextEdit2("右部文本", splitterRight)
	//textbuttom := widgets.NewQTextEdit2("下部文本", splitterLeft)
	//splitterRight.AddWidget(textRight)
	//splitterRight.AddWidget(textbuttom)
	//
	//splitterLeft.SetWindowTitle("splitter")
	//splitterLeft.Show()


	//widgets.QApplication_Exec()
	app.Exec()
}