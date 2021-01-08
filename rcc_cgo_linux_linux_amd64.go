// +build !ubports

package main

/*
#cgo CFLAGS: -pipe -O2 -Wall -Wextra -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -O2 -Wall -Wextra -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../myopen -I. -I/usr/include/qt -I/usr/include/qt/QtGui -I/usr/include/qt/QtCore -I. -I/usr/lib/qt/mkspecs/linux-g++
#cgo LDFLAGS: -O1
#cgo LDFLAGS:  /usr/lib/libQt5Gui.so /usr/lib/libQt5Core.so -lGL -lpthread
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
