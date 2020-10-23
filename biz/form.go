package biz

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
	"path/filepath"
	"sumery.dev/common/utils/file"
)
type MainWindow struct {
	widgets.QDialog
	_ func() `constructor:"init"`
	//_ func() `signal:"layout,auto"`

	txtServer, txtPort, txtDb, txtUser, txtPwd, txtPkg,txtTable *widgets.QLineEdit
	cbxGo, cbxJs *widgets.QCheckBox
	btnOk, btnCancel *widgets.QPushButton

	cfg *ini.File
}

func (me *MainWindow) init()  {
	me.layout()
	me.load()
}

func (me *MainWindow) layout()  {
	var (
		hb *widgets.QHBoxLayout
	)
	vb := widgets.NewQVBoxLayout()
	me.txtServer, hb = me.newEdit("地址", "")
	vb.AddLayout(hb, 0)
	me.txtServer, hb = me.newEdit("端口", "")
	vb.AddLayout(hb, 0)
	me.txtServer, hb = me.newEdit("库名", "")
	vb.AddLayout(hb, 0)
	me.txtServer, hb = me.newEdit("用户", "")
	vb.AddLayout(hb, 0)
	me.txtServer, hb = me.newEdit("密码", "")
	vb.AddLayout(hb, 0)

	gp := widgets.NewQGroupBox(me)
	gp.SetTitle("数据库信息")
	gp.SetLayout(vb)

	mvl := widgets.NewQVBoxLayout()
	mvl.AddWidget(gp, 0, 0)

	me.txtPkg, hb = me.newEdit("包名", "")
	mvl.AddLayout(hb, 0)
	me.txtTable, hb = me.newEdit("指定表名", "")
	mvl.AddLayout(hb, 0)

	me.cbxGo = widgets.NewQCheckBox2("生成Go文件", me)
	me.cbxJs = widgets.NewQCheckBox2("生成js文件", me)
	hb = widgets.NewQHBoxLayout()
	hb.AddWidget(me.cbxGo, 0, 0)
	hb.AddWidget(me.cbxJs, 0, 0)
	mvl.AddLayout(hb, 0)

	// main horizontal layout
	ml := widgets.NewQHBoxLayout()
	// the left of main
	ml.AddLayout(mvl, 0)

	// button of right
	me.btnOk = widgets.NewQPushButton2("生成", me)
	me.btnCancel = widgets.NewQPushButton2("关闭", me)
	mvr := widgets.NewQVBoxLayout()
	mvr.AddWidget(me.btnOk, 0, 0)
	mvr.AddWidget(me.btnCancel, 0, 0)
	mvr.AddStretch(0)

	// the right of main
	ml.AddLayout(mvr, 0)

	me.SetLayout(ml)
	me.SetWindowTitle("模型文件生成器")

	me.btnOk.ConnectClicked(me.onOk)
	me.btnCancel.ConnectClicked(func(checked bool) {
		me.save()
		me.Done(0)
	})
}

func (me *MainWindow) newEdit(label, text string) (*widgets.QLineEdit, *widgets.QHBoxLayout) {
	lbl := widgets.NewQLabel2(label, me, 0)
	txt := widgets.NewQLineEdit2(text, me)
	lbl.SetBuddy(txt)

	lyt := widgets.NewQHBoxLayout()
	lyt.AddWidget(lbl, 0, 0)
	lyt.AddWidget(txt, 0, 0)

	return txt, lyt
}

func (me *MainWindow) load() {
	mp := "app.ini"
	mp = getPath(mp)
	if !file.FileExists(mp) {
		fl, _ := file.Open(mp, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		fl.Close()
	}

	me.cfg, _ = ini.Load(mp)
	me.txtServer.SetText(me.cfg.Section("db").Key("server").String())
	me.txtPort.SetText(me.cfg.Section("db").Key("port").String())
	me.txtDb.SetText(me.cfg.Section("db").Key("db").String())
	me.txtUser.SetText(me.cfg.Section("db").Key("user").String())
	me.txtPwd.SetText(me.cfg.Section("db").Key("pwd").String())
	me.txtPkg.SetText(me.cfg.Section("model").Key("pkg").String())
	me.txtTable.SetText(me.cfg.Section("model").Key("table").String())
	if me.cfg.Section("model").Key("go").String() == "1" {
		me.cbxGo.SetChecked(true)
	}
	if me.cfg.Section("model").Key("js").String() == "1" {
		me.cbxJs.SetChecked(true)
	}
}

func (me *MainWindow) save() {
	me.cfg.Section("db").Key("server").SetValue(me.txtServer.Text())
	me.cfg.Section("db").Key("port").SetValue(me.txtPort.Text())
	me.cfg.Section("db").Key("db").SetValue(me.txtDb.Text())
	me.cfg.Section("db").Key("user").SetValue(me.txtUser.Text())
	me.cfg.Section("db").Key("pwd").SetValue(me.txtPwd.Text())
	me.cfg.Section("model").Key("pkg").SetValue(me.txtPkg.Text())
	me.cfg.Section("model").Key("table").SetValue(me.txtTable.Text())
	chk := "0"
	if me.cbxGo.IsChecked(){
		chk = "1"
	}
	me.cfg.Section("model").Key("go").SetValue(chk)
	chk = "0"
	if me.cbxJs.IsChecked(){
		chk = "1"
	}
	me.cfg.Section("model").Key("js").SetValue(chk)
}

func (me *MainWindow) onOk(checked bool) {
	dbName := me.txtDb.Text()
	user := me.txtUser.Text()
	password := me.txtPwd.Text()
	host := me.txtServer.Text()
	port := me.txtPort.Text()

	importName := me.txtPkg.Text()
	tableName := me.txtTable.Text()

	err := DBInit(dbName, user, password, host, port)
	if err != nil {

		dlg := widgets.NewQMessageBox2(widgets.QMessageBox__Warning, "出错了", fmt.Sprintf("%s", err), widgets.QMessageBox__Ok, me, core.Qt__Dialog)
		dlg.Show()
	}
	if me.cbxGo.IsChecked() {
		err = ModelGenerate(importName, tableName, getPath("./tpl/model_go.tmpl"))
	}
	if me.cbxJs.IsChecked() {
		err = ModelGenerate(importName, tableName, getPath("./tpl/model_js.tmpl"))
	}
	if err != nil {
		dlg := widgets.NewQMessageBox2(widgets.QMessageBox__Information, "执行成功", "文件生成成功！", widgets.QMessageBox__Ok, me, core.Qt__Dialog)
		dlg.Show()
	}
}

func getPath(name string) string {
	mp, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	mp = filepath.Join(mp, name)
	return mp
}