package biz

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
	std_widgets "github.com/therecipe/qt/widgets"
)

func cGoFreePacked(ptr unsafe.Pointer) { std_core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_Moc_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type MainWindow_ITF interface {
	std_widgets.QDialog_ITF
	MainWindow_PTR() *MainWindow
}

func (ptr *MainWindow) MainWindow_PTR() *MainWindow {
	return ptr
}

func (ptr *MainWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *MainWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromMainWindow(ptr MainWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MainWindow_PTR().Pointer()
	}
	return nil
}

func NewMainWindowFromPointer(ptr unsafe.Pointer) (n *MainWindow) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(MainWindow)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *MainWindow:
			n = deduced

		case *std_widgets.QDialog:
			n = &MainWindow{QDialog: *deduced}

		default:
			n = new(MainWindow)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *MainWindow) Init() { this.init() }

//export callbackMainWindow16b1d1_Constructor
func callbackMainWindow16b1d1_Constructor(ptr unsafe.Pointer) {
	this := NewMainWindowFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

func MainWindow_QRegisterMetaType() int {
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QRegisterMetaType()))
}

func (ptr *MainWindow) QRegisterMetaType() int {
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QRegisterMetaType()))
}

func MainWindow_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QRegisterMetaType2(typeNameC)))
}

func (ptr *MainWindow) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QRegisterMetaType2(typeNameC)))
}

func MainWindow_QmlRegisterType() int {
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QmlRegisterType()))
}

func (ptr *MainWindow) QmlRegisterType() int {
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QmlRegisterType()))
}

func MainWindow_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *MainWindow) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func MainWindow_QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *MainWindow) QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func MainWindow_QmlRegisterAnonymousType(uri string, versionMajor int) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QmlRegisterAnonymousType(uriC, C.int(int32(versionMajor)))))
}

func (ptr *MainWindow) QmlRegisterAnonymousType(uri string, versionMajor int) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	return int(int32(C.MainWindow16b1d1_MainWindow16b1d1_QmlRegisterAnonymousType(uriC, C.int(int32(versionMajor)))))
}

func (ptr *MainWindow) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MainWindow16b1d1___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MainWindow) __actions_newList() unsafe.Pointer {
	return C.MainWindow16b1d1___actions_newList(ptr.Pointer())
}

func (ptr *MainWindow) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MainWindow16b1d1___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MainWindow) __addActions_actions_newList() unsafe.Pointer {
	return C.MainWindow16b1d1___addActions_actions_newList(ptr.Pointer())
}

func (ptr *MainWindow) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.MainWindow16b1d1___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *MainWindow) __insertActions_actions_newList() unsafe.Pointer {
	return C.MainWindow16b1d1___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *MainWindow) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MainWindow16b1d1___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MainWindow) __children_newList() unsafe.Pointer {
	return C.MainWindow16b1d1___children_newList(ptr.Pointer())
}

func (ptr *MainWindow) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.MainWindow16b1d1___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *MainWindow) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.MainWindow16b1d1___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *MainWindow) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MainWindow16b1d1___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MainWindow) __findChildren_newList() unsafe.Pointer {
	return C.MainWindow16b1d1___findChildren_newList(ptr.Pointer())
}

func (ptr *MainWindow) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.MainWindow16b1d1___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *MainWindow) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *MainWindow) __findChildren_newList3() unsafe.Pointer {
	return C.MainWindow16b1d1___findChildren_newList3(ptr.Pointer())
}

func NewMainWindow(parent std_widgets.QWidget_ITF, ff std_core.Qt__WindowType) *MainWindow {
	MainWindow_QRegisterMetaType()
	tmpValue := NewMainWindowFromPointer(C.MainWindow16b1d1_NewMainWindow(std_widgets.PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackMainWindow16b1d1_DestroyMainWindow
func callbackMainWindow16b1d1_DestroyMainWindow(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~MainWindow"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).DestroyMainWindowDefault()
	}
}

func (ptr *MainWindow) ConnectDestroyMainWindow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~MainWindow"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~MainWindow", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~MainWindow", unsafe.Pointer(&f))
		}
	}
}

func (ptr *MainWindow) DisconnectDestroyMainWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~MainWindow")
	}
}

func (ptr *MainWindow) DestroyMainWindow() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.MainWindow16b1d1_DestroyMainWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *MainWindow) DestroyMainWindowDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.MainWindow16b1d1_DestroyMainWindowDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackMainWindow16b1d1_Accept
func callbackMainWindow16b1d1_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accept"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *MainWindow) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_AcceptDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_Accepted
func callbackMainWindow16b1d1_Accepted(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accepted"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackMainWindow16b1d1_CloseEvent
func callbackMainWindow16b1d1_CloseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(e))
	} else {
		NewMainWindowFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *MainWindow) CloseEventDefault(e std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(e))
	}
}

//export callbackMainWindow16b1d1_ContextMenuEvent
func callbackMainWindow16b1d1_ContextMenuEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewMainWindowFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *MainWindow) ContextMenuEventDefault(e std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(e))
	}
}

//export callbackMainWindow16b1d1_Done
func callbackMainWindow16b1d1_Done(ptr unsafe.Pointer, r C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		(*(*func(int))(signal))(int(int32(r)))
	} else {
		NewMainWindowFromPointer(ptr).DoneDefault(int(int32(r)))
	}
}

func (ptr *MainWindow) DoneDefault(r int) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_DoneDefault(ptr.Pointer(), C.int(int32(r)))
	}
}

//export callbackMainWindow16b1d1_EventFilter
func callbackMainWindow16b1d1_EventFilter(ptr unsafe.Pointer, o unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(o), std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(o), std_core.NewQEventFromPointer(e)))))
}

func (ptr *MainWindow) EventFilterDefault(o std_core.QObject_ITF, e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow16b1d1_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(o), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackMainWindow16b1d1_Exec
func callbackMainWindow16b1d1_Exec(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "exec"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewMainWindowFromPointer(ptr).ExecDefault()))
}

func (ptr *MainWindow) ExecDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.MainWindow16b1d1_ExecDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackMainWindow16b1d1_Finished
func callbackMainWindow16b1d1_Finished(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func(int))(signal))(int(int32(result)))
	}

}

//export callbackMainWindow16b1d1_KeyPressEvent
func callbackMainWindow16b1d1_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(e))
	} else {
		NewMainWindowFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *MainWindow) KeyPressEventDefault(e std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(e))
	}
}

//export callbackMainWindow16b1d1_MinimumSizeHint
func callbackMainWindow16b1d1_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMainWindowFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *MainWindow) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MainWindow16b1d1_MinimumSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMainWindow16b1d1_Open
func callbackMainWindow16b1d1_Open(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).OpenDefault()
	}
}

func (ptr *MainWindow) OpenDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_OpenDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_Reject
func callbackMainWindow16b1d1_Reject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reject"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).RejectDefault()
	}
}

func (ptr *MainWindow) RejectDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_RejectDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_Rejected
func callbackMainWindow16b1d1_Rejected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rejected"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackMainWindow16b1d1_ResizeEvent
func callbackMainWindow16b1d1_ResizeEvent(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(vqr))
	} else {
		NewMainWindowFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(vqr))
	}
}

func (ptr *MainWindow) ResizeEventDefault(vqr std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(vqr))
	}
}

//export callbackMainWindow16b1d1_SetVisible
func callbackMainWindow16b1d1_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *MainWindow) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackMainWindow16b1d1_ShowEvent
func callbackMainWindow16b1d1_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *MainWindow) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackMainWindow16b1d1_SizeHint
func callbackMainWindow16b1d1_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewMainWindowFromPointer(ptr).SizeHintDefault())
}

func (ptr *MainWindow) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.MainWindow16b1d1_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackMainWindow16b1d1_ActionEvent
func callbackMainWindow16b1d1_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *MainWindow) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackMainWindow16b1d1_ChangeEvent
func callbackMainWindow16b1d1_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MainWindow) ChangeEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMainWindow16b1d1_Close
func callbackMainWindow16b1d1_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).CloseDefault())))
}

func (ptr *MainWindow) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow16b1d1_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMainWindow16b1d1_CustomContextMenuRequested
func callbackMainWindow16b1d1_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackMainWindow16b1d1_DragEnterEvent
func callbackMainWindow16b1d1_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *MainWindow) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackMainWindow16b1d1_DragLeaveEvent
func callbackMainWindow16b1d1_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *MainWindow) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackMainWindow16b1d1_DragMoveEvent
func callbackMainWindow16b1d1_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *MainWindow) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackMainWindow16b1d1_DropEvent
func callbackMainWindow16b1d1_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *MainWindow) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackMainWindow16b1d1_EnterEvent
func callbackMainWindow16b1d1_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MainWindow) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMainWindow16b1d1_Event
func callbackMainWindow16b1d1_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *MainWindow) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow16b1d1_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackMainWindow16b1d1_FocusInEvent
func callbackMainWindow16b1d1_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MainWindow) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMainWindow16b1d1_FocusNextPrevChild
func callbackMainWindow16b1d1_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *MainWindow) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow16b1d1_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackMainWindow16b1d1_FocusOutEvent
func callbackMainWindow16b1d1_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *MainWindow) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackMainWindow16b1d1_HasHeightForWidth
func callbackMainWindow16b1d1_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *MainWindow) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.MainWindow16b1d1_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackMainWindow16b1d1_HeightForWidth
func callbackMainWindow16b1d1_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewMainWindowFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *MainWindow) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MainWindow16b1d1_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackMainWindow16b1d1_Hide
func callbackMainWindow16b1d1_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *MainWindow) HideDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_HideDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_HideEvent
func callbackMainWindow16b1d1_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *MainWindow) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackMainWindow16b1d1_InitPainter
func callbackMainWindow16b1d1_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewMainWindowFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *MainWindow) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackMainWindow16b1d1_InputMethodEvent
func callbackMainWindow16b1d1_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *MainWindow) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackMainWindow16b1d1_InputMethodQuery
func callbackMainWindow16b1d1_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewMainWindowFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *MainWindow) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.MainWindow16b1d1_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackMainWindow16b1d1_KeyReleaseEvent
func callbackMainWindow16b1d1_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *MainWindow) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackMainWindow16b1d1_LeaveEvent
func callbackMainWindow16b1d1_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MainWindow) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMainWindow16b1d1_Lower
func callbackMainWindow16b1d1_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *MainWindow) LowerDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_LowerDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_Metric
func callbackMainWindow16b1d1_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewMainWindowFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *MainWindow) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.MainWindow16b1d1_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackMainWindow16b1d1_MouseDoubleClickEvent
func callbackMainWindow16b1d1_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MainWindow) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMainWindow16b1d1_MouseMoveEvent
func callbackMainWindow16b1d1_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MainWindow) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMainWindow16b1d1_MousePressEvent
func callbackMainWindow16b1d1_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MainWindow) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMainWindow16b1d1_MouseReleaseEvent
func callbackMainWindow16b1d1_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *MainWindow) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackMainWindow16b1d1_MoveEvent
func callbackMainWindow16b1d1_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *MainWindow) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackMainWindow16b1d1_NativeEvent
func callbackMainWindow16b1d1_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMainWindowFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *MainWindow) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.MainWindow16b1d1_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackMainWindow16b1d1_PaintEngine
func callbackMainWindow16b1d1_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewMainWindowFromPointer(ptr).PaintEngineDefault())
}

func (ptr *MainWindow) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.MainWindow16b1d1_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackMainWindow16b1d1_PaintEvent
func callbackMainWindow16b1d1_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *MainWindow) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackMainWindow16b1d1_Raise
func callbackMainWindow16b1d1_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *MainWindow) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_RaiseDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_Repaint
func callbackMainWindow16b1d1_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *MainWindow) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_RepaintDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_SetDisabled
func callbackMainWindow16b1d1_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *MainWindow) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackMainWindow16b1d1_SetEnabled
func callbackMainWindow16b1d1_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *MainWindow) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMainWindow16b1d1_SetFocus2
func callbackMainWindow16b1d1_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *MainWindow) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_SetHidden
func callbackMainWindow16b1d1_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *MainWindow) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackMainWindow16b1d1_SetStyleSheet
func callbackMainWindow16b1d1_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewMainWindowFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *MainWindow) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.MainWindow16b1d1_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackMainWindow16b1d1_SetWindowModified
func callbackMainWindow16b1d1_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewMainWindowFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *MainWindow) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackMainWindow16b1d1_SetWindowTitle
func callbackMainWindow16b1d1_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewMainWindowFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *MainWindow) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.MainWindow16b1d1_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackMainWindow16b1d1_Show
func callbackMainWindow16b1d1_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *MainWindow) ShowDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ShowDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_ShowFullScreen
func callbackMainWindow16b1d1_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *MainWindow) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_ShowMaximized
func callbackMainWindow16b1d1_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *MainWindow) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_ShowMinimized
func callbackMainWindow16b1d1_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *MainWindow) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_ShowNormal
func callbackMainWindow16b1d1_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *MainWindow) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_TabletEvent
func callbackMainWindow16b1d1_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *MainWindow) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackMainWindow16b1d1_Update
func callbackMainWindow16b1d1_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *MainWindow) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_UpdateDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_UpdateMicroFocus
func callbackMainWindow16b1d1_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *MainWindow) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_WheelEvent
func callbackMainWindow16b1d1_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *MainWindow) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackMainWindow16b1d1_WindowIconChanged
func callbackMainWindow16b1d1_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackMainWindow16b1d1_WindowTitleChanged
func callbackMainWindow16b1d1_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackMainWindow16b1d1_ChildEvent
func callbackMainWindow16b1d1_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *MainWindow) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackMainWindow16b1d1_ConnectNotify
func callbackMainWindow16b1d1_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMainWindowFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MainWindow) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMainWindow16b1d1_CustomEvent
func callbackMainWindow16b1d1_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *MainWindow) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMainWindow16b1d1_DeleteLater
func callbackMainWindow16b1d1_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewMainWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *MainWindow) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.MainWindow16b1d1_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackMainWindow16b1d1_Destroyed
func callbackMainWindow16b1d1_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackMainWindow16b1d1_DisconnectNotify
func callbackMainWindow16b1d1_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMainWindowFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *MainWindow) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMainWindow16b1d1_ObjectNameChanged
func callbackMainWindow16b1d1_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackMainWindow16b1d1_TimerEvent
func callbackMainWindow16b1d1_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewMainWindowFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *MainWindow) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.MainWindow16b1d1_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

func init() {
}
