package ui

import (
	"cqccteg.hlw/up/common/logging"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

const CFG_NAME = "gentity_cfg"
const VERSION = "V1.0.0-20221207-YaoGL"

var (
	application fyne.App
	win         fyne.Window
	cfg         *FormConfig
)

func Run() {
	logging.InfoF("Version: %s", VERSION)
	application = app.NewWithID("gentity.sumery.dev")
	//rs := fyne.NewStaticResource(CFG_NAME, []byte(ICO_DATA))
	if rs, err := fyne.LoadResourceFromPath("./assets/icon.ico"); err == nil {
		application.SetIcon(rs)
	}
	win = application.NewWindow("模型文件生成器 " + VERSION)
	win.Resize(fyne.NewSize(400, 500))
	cfg := GetCfg()
	fm := NewFormUI()

	win.SetContent(container.NewGridWithColumns(1, fm.Build(cfg)))
	win.ShowAndRun()
	_ = os.Unsetenv("FYNE_FONT")

}
func SetCfg(fc *FormConfig) {
	if fc == nil {
		fc = &FormConfig{}
	}
	if ss, err := application.Storage().Save(CFG_NAME); err == nil {
		if n, err := ss.Write(fc.Byte()); err != nil {
			logging.Err(err)
		} else {
			logging.InfoF("Write config %n bytes of %s", n, fc.Byte())
		}
		_ = ss.Close()
	}

}
func GetCfg() *FormConfig {
	rd, err := application.Storage().Open(CFG_NAME)

	fc := FormConfig{}
	if err != nil {
		logging.Err(err)
		//if wt, err := application.Storage().Create(CFG_NAME); err == nil {
		//	if n, err := wt.Write(fc.Byte()); err != nil {
		//		logging.Err(err)
		//	} else {
		//		logging.InfoF("Write config %n bytes of %s", n, fc.Byte())
		//	}
		//	_ = wt.Close()
		//}
	} else {
		p := rd.URI().Path()
		_ = rd.Close()
		bts, err := os.ReadFile(p)
		if err != nil || len(bts) == 0 {
			logging.Err(err)
		}
		FormConfigFromByte(&fc, bts)
	}
	return &fc
}

func init() {
	fontPaths := findfont.List()
	//logging.InfoF("Fonts found %n ", len(fontPaths))
	for _, path := range fontPaths {
		//fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
	//fmt.Println("=============")
}
