package main

import (
	"github.com/oxiao/gentity/biz"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	ico := gui.NewQIcon5(":/assets/icon.ico")
	dlg := biz.NewMainWindow(nil, 0)
	dlg.SetWindowIcon(ico)
	dlg.Show()

	//mw := widgets.NewQDialog(nil, core.Qt__Dialog)
	//au := ui.UIAppDialog{}
	//au.SetupUI(mw)
	//mw.Show()
	app.Exec()
}