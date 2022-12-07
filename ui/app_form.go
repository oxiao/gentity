package ui

import (
	"cqccteg.hlw/up/common/logging"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/oxiao/gentity/biz"
	"os"
	"path/filepath"
)

type FormConfig struct {
	Dbms   string `json:"dbms"`
	Server string `json:"server"`
	Port   string `json:"port"`
	Db     string `json:"db"`
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	Pkg    string `json:"pkg"`
	Table  string `json:"table"`
	Go     bool   `json:"go"`
	Js     bool   `json:"js"`
}

type FormUI struct {
	dbms   *widget.SelectEntry
	server *widget.Entry
	port   *widget.Entry
	db     *widget.Entry
	user   *widget.Entry
	pwd    *widget.Entry
	pkg    *widget.Entry
	table  *widget.Entry
	goo    *widget.Check
	js     *widget.Check
}

func NewFormUI() *FormUI {
	fui := FormUI{
		dbms:   widget.NewSelectEntry([]string{"mysql", "mssql", "oracle", "postgres"}),
		server: widget.NewEntry(),
		port:   widget.NewEntry(),
		db:     widget.NewEntry(),
		user:   widget.NewEntry(),
		pwd:    widget.NewPasswordEntry(),
		pkg:    widget.NewEntry(),
		table:  widget.NewEntry(),
		goo: widget.NewCheck("", func(b bool) {

		}),
		js: widget.NewCheck("", func(b bool) {

		}),
	}

	fui.dbms.SetPlaceHolder("选择数据库系统名称")
	fui.server.SetPlaceHolder("数据服务IP地址")
	fui.port.SetPlaceHolder("数据库服务端口")
	fui.db.SetPlaceHolder("数据库名称")
	fui.user.SetPlaceHolder("数据库访问账号")
	fui.pwd.SetPlaceHolder("数据库账号密码")
	fui.pkg.SetPlaceHolder("要生成的源代码包的名称")

	return &fui
}

func (me *FormUI) Build(fc *FormConfig) fyne.CanvasObject {
	fun := func() {
		fc.Dbms = me.dbms.Text
		fc.Server = me.server.Text
		fc.Port = me.port.Text
		fc.Db = me.db.Text
		fc.User = me.user.Text
		fc.Pwd = me.pwd.Text
		fc.Pkg = me.pkg.Text
		fc.Table = me.table.Text
		fc.Go = me.goo.Checked
		fc.Js = me.js.Checked
	}

	fm := widget.NewForm(
		&widget.FormItem{Text: "数据系统:", Widget: me.dbms},
		&widget.FormItem{Text: "地址:", Widget: me.server},
		&widget.FormItem{Text: "端口:", Widget: me.port},
		&widget.FormItem{Text: "数据库:", Widget: me.db},
		&widget.FormItem{Text: "用户:", Widget: me.user},
		&widget.FormItem{Text: "密码:", Widget: me.pwd},
		&widget.FormItem{Text: "包名:", Widget: me.pkg},
		&widget.FormItem{Text: "生成js文件:", Widget: me.js},
		&widget.FormItem{Text: "生成Go文件:", Widget: me.goo},
	)
	fm.CancelText = "退出"
	fm.SubmitText = "生成"
	fm.OnSubmit = func() {
		fun()
		SetCfg(fc)
		me.onOk()
	}
	fm.OnCancel = func() {
		fun()
		SetCfg(fc)
		win.Close()
		application.Quit()
	}

	me.dbms.SetText(fc.Dbms)
	me.server.SetText(fc.Server)
	me.port.SetText(fc.Port)
	me.db.SetText(fc.Db)
	me.user.SetText(fc.User)
	me.pwd.SetText(fc.Pwd)
	me.pkg.SetText(fc.Pkg)
	me.table.SetText(fc.Table)
	me.goo.SetChecked(fc.Go)
	me.js.SetChecked(fc.Js)

	card := widget.NewCard("数据库信息", "", fm)
	return container.NewVScroll(card)
}
func (me *FormUI) onOk() {

	var (
		err error
		dlg dialog.Dialog
	)
	err = biz.DBInit(me.dbms.Text,
		me.db.Text,
		me.user.Text,
		me.pwd.Text,
		me.server.Text,
		me.port.Text)
	if err != nil {
		logging.Err(err)
		dlg = dialog.NewError(err, win)
		dlg.Show()
		return
	}
	if me.goo.Checked {
		err = biz.ModelGenerate(me.pkg.Text,
			me.table.Text,
			getPath("./tpl/model_go.tmpl"))
		logging.Err(err)
	}
	if me.js.Checked {
		err = biz.ModelGenerate(me.pkg.Text,
			me.table.Text,
			getPath("./tpl/model_js.tmpl"))
		logging.Err(err)
	}

	if err != nil {
		logging.Err(err)
		dlg = dialog.NewError(err, win)
	} else {
		dlg = dialog.NewInformation("执行成功", "文件生成成功！", win)
	}
	dlg.Show()
}
func getPath(name string) string {
	mp, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	mp = filepath.Join(mp, name)
	return mp
}
func (me *FormConfig) String() string {
	return string(me.Byte())
}
func (me *FormConfig) Byte() []byte {
	bts, err := json.Marshal(me)
	if err != nil {
		logging.Err(err)
	}
	return bts
}
func FormConfigFromByte(fc *FormConfig, bts []byte) {

	err := json.Unmarshal(bts, fc)
	if err != nil {
		logging.Err(err)
	}
}
