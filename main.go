package main

import (
	"cqccteg.hlw/up/common/conf/setting/serverinit"
	"cqccteg.hlw/up/common/logging"
	"github.com/oxiao/gentity/ui"
)

const VERSION = "1.1.0-20230513-YaoGL"

func main() {
	defer serverinit.PanicRecover(VERSION)

	logging.Setup()
	ui.Run()
}
