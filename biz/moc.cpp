

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDialog>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaMethod>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class MainWindow16b1d1: public QDialog
{
Q_OBJECT
public:
	MainWindow16b1d1(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QDialog(parent, ff) {qRegisterMetaType<quintptr>("quintptr");MainWindow16b1d1_MainWindow16b1d1_QRegisterMetaType();MainWindow16b1d1_MainWindow16b1d1_QRegisterMetaTypes();callbackMainWindow16b1d1_Constructor(this);};
	 ~MainWindow16b1d1() { callbackMainWindow16b1d1_DestroyMainWindow(this); };
	void accept() { callbackMainWindow16b1d1_Accept(this); };
	void Signal_Accepted() { callbackMainWindow16b1d1_Accepted(this); };
	void closeEvent(QCloseEvent * e) { callbackMainWindow16b1d1_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackMainWindow16b1d1_ContextMenuEvent(this, e); };
	void done(int r) { callbackMainWindow16b1d1_Done(this, r); };
	bool eventFilter(QObject * o, QEvent * e) { return callbackMainWindow16b1d1_EventFilter(this, o, e) != 0; };
	int exec() { return callbackMainWindow16b1d1_Exec(this); };
	void Signal_Finished(int result) { callbackMainWindow16b1d1_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackMainWindow16b1d1_KeyPressEvent(this, e); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackMainWindow16b1d1_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void open() { callbackMainWindow16b1d1_Open(this); };
	void reject() { callbackMainWindow16b1d1_Reject(this); };
	void Signal_Rejected() { callbackMainWindow16b1d1_Rejected(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackMainWindow16b1d1_ResizeEvent(this, vqr); };
	void setVisible(bool visible) { callbackMainWindow16b1d1_SetVisible(this, visible); };
	void showEvent(QShowEvent * event) { callbackMainWindow16b1d1_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackMainWindow16b1d1_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void actionEvent(QActionEvent * event) { callbackMainWindow16b1d1_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackMainWindow16b1d1_ChangeEvent(this, event); };
	bool close() { return callbackMainWindow16b1d1_Close(this) != 0; };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackMainWindow16b1d1_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackMainWindow16b1d1_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackMainWindow16b1d1_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackMainWindow16b1d1_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackMainWindow16b1d1_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackMainWindow16b1d1_EnterEvent(this, event); };
	bool event(QEvent * event) { return callbackMainWindow16b1d1_Event(this, event) != 0; };
	void focusInEvent(QFocusEvent * event) { callbackMainWindow16b1d1_FocusInEvent(this, event); };
	bool focusNextPrevChild(bool next) { return callbackMainWindow16b1d1_FocusNextPrevChild(this, next) != 0; };
	void focusOutEvent(QFocusEvent * event) { callbackMainWindow16b1d1_FocusOutEvent(this, event); };
	bool hasHeightForWidth() const { return callbackMainWindow16b1d1_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackMainWindow16b1d1_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackMainWindow16b1d1_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackMainWindow16b1d1_HideEvent(this, event); };
	void initPainter(QPainter * painter) const { callbackMainWindow16b1d1_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackMainWindow16b1d1_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackMainWindow16b1d1_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	void keyReleaseEvent(QKeyEvent * event) { callbackMainWindow16b1d1_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackMainWindow16b1d1_LeaveEvent(this, event); };
	void lower() { callbackMainWindow16b1d1_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackMainWindow16b1d1_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackMainWindow16b1d1_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackMainWindow16b1d1_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackMainWindow16b1d1_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackMainWindow16b1d1_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackMainWindow16b1d1_MoveEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackMainWindow16b1d1_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackMainWindow16b1d1_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	void paintEvent(QPaintEvent * event) { callbackMainWindow16b1d1_PaintEvent(this, event); };
	void raise() { callbackMainWindow16b1d1_Raise(this); };
	void repaint() { callbackMainWindow16b1d1_Repaint(this); };
	void setDisabled(bool disable) { callbackMainWindow16b1d1_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackMainWindow16b1d1_SetEnabled(this, vbo); };
	void setFocus() { callbackMainWindow16b1d1_SetFocus2(this); };
	void setHidden(bool hidden) { callbackMainWindow16b1d1_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackMainWindow16b1d1_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackMainWindow16b1d1_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackMainWindow16b1d1_SetWindowTitle(this, vqsPacked); };
	void show() { callbackMainWindow16b1d1_Show(this); };
	void showFullScreen() { callbackMainWindow16b1d1_ShowFullScreen(this); };
	void showMaximized() { callbackMainWindow16b1d1_ShowMaximized(this); };
	void showMinimized() { callbackMainWindow16b1d1_ShowMinimized(this); };
	void showNormal() { callbackMainWindow16b1d1_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackMainWindow16b1d1_TabletEvent(this, event); };
	void update() { callbackMainWindow16b1d1_Update(this); };
	void updateMicroFocus() { callbackMainWindow16b1d1_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackMainWindow16b1d1_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackMainWindow16b1d1_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray* t3c6de1 = new QByteArray(title.toUtf8()); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1->prepend("WHITESPACE").constData()+10), t3c6de1->size()-10, t3c6de1 };callbackMainWindow16b1d1_WindowTitleChanged(this, titlePacked); };
	void childEvent(QChildEvent * event) { callbackMainWindow16b1d1_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMainWindow16b1d1_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMainWindow16b1d1_CustomEvent(this, event); };
	void deleteLater() { callbackMainWindow16b1d1_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMainWindow16b1d1_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMainWindow16b1d1_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackMainWindow16b1d1_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackMainWindow16b1d1_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(MainWindow16b1d1*)


void MainWindow16b1d1_MainWindow16b1d1_QRegisterMetaTypes() {
}

int MainWindow16b1d1_MainWindow16b1d1_QRegisterMetaType()
{
	return qRegisterMetaType<MainWindow16b1d1*>();
}

int MainWindow16b1d1_MainWindow16b1d1_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<MainWindow16b1d1*>(const_cast<const char*>(typeName));
}

int MainWindow16b1d1_MainWindow16b1d1_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MainWindow16b1d1>();
#else
	return 0;
#endif
}

int MainWindow16b1d1_MainWindow16b1d1_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<MainWindow16b1d1>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int MainWindow16b1d1_MainWindow16b1d1_QmlRegisterUncreatableType(char* uri, int versionMajor, int versionMinor, char* qmlName, struct Moc_PackedString message)
{
#ifdef QT_QML_LIB
	return qmlRegisterUncreatableType<MainWindow16b1d1>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName), QString::fromUtf8(message.data, message.len));
#else
	return 0;
#endif
}

int MainWindow16b1d1_MainWindow16b1d1_QmlRegisterAnonymousType(char* uri, int versionMajor)
{
#ifdef QT_QML_LIB
	return qmlRegisterAnonymousType<MainWindow16b1d1>(const_cast<const char*>(uri), versionMajor);
#else
	return 0;
#endif
}

void* MainWindow16b1d1___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow16b1d1___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MainWindow16b1d1___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MainWindow16b1d1___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow16b1d1___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MainWindow16b1d1___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MainWindow16b1d1___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow16b1d1___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* MainWindow16b1d1___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* MainWindow16b1d1___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow16b1d1___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MainWindow16b1d1___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* MainWindow16b1d1___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void MainWindow16b1d1___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* MainWindow16b1d1___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* MainWindow16b1d1___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow16b1d1___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MainWindow16b1d1___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MainWindow16b1d1___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void MainWindow16b1d1___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* MainWindow16b1d1___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* MainWindow16b1d1_NewMainWindow(void* parent, long long ff)
{
		return new MainWindow16b1d1(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void MainWindow16b1d1_DestroyMainWindow(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->~MainWindow16b1d1();
}

void MainWindow16b1d1_DestroyMainWindowDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void MainWindow16b1d1_AcceptDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::accept();
}

void MainWindow16b1d1_CloseEventDefault(void* ptr, void* e)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::closeEvent(static_cast<QCloseEvent*>(e));
}

void MainWindow16b1d1_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void MainWindow16b1d1_DoneDefault(void* ptr, int r)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::done(r);
}

char MainWindow16b1d1_EventFilterDefault(void* ptr, void* o, void* e)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::eventFilter(static_cast<QObject*>(o), static_cast<QEvent*>(e));
}

int MainWindow16b1d1_ExecDefault(void* ptr)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::exec();
}

void MainWindow16b1d1_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void* MainWindow16b1d1_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MainWindow16b1d1*>(ptr)->QDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void MainWindow16b1d1_OpenDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::open();
}

void MainWindow16b1d1_RejectDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::reject();
}

void MainWindow16b1d1_ResizeEventDefault(void* ptr, void* vqr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void MainWindow16b1d1_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::setVisible(visible != 0);
}

void MainWindow16b1d1_ShowEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::showEvent(static_cast<QShowEvent*>(event));
}

void* MainWindow16b1d1_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<MainWindow16b1d1*>(ptr)->QDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void MainWindow16b1d1_ActionEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::actionEvent(static_cast<QActionEvent*>(event));
}

void MainWindow16b1d1_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::changeEvent(static_cast<QEvent*>(event));
}

char MainWindow16b1d1_CloseDefault(void* ptr)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::close();
}

void MainWindow16b1d1_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void MainWindow16b1d1_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void MainWindow16b1d1_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void MainWindow16b1d1_DropEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::dropEvent(static_cast<QDropEvent*>(event));
}

void MainWindow16b1d1_EnterEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::enterEvent(static_cast<QEvent*>(event));
}

char MainWindow16b1d1_EventDefault(void* ptr, void* event)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::event(static_cast<QEvent*>(event));
}

void MainWindow16b1d1_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::focusInEvent(static_cast<QFocusEvent*>(event));
}

char MainWindow16b1d1_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::focusNextPrevChild(next != 0);
}

void MainWindow16b1d1_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
}

char MainWindow16b1d1_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::hasHeightForWidth();
}

int MainWindow16b1d1_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::heightForWidth(w);
}

void MainWindow16b1d1_HideDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::hide();
}

void MainWindow16b1d1_HideEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::hideEvent(static_cast<QHideEvent*>(event));
}

void MainWindow16b1d1_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::initPainter(static_cast<QPainter*>(painter));
}

void MainWindow16b1d1_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* MainWindow16b1d1_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<MainWindow16b1d1*>(ptr)->QDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void MainWindow16b1d1_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void MainWindow16b1d1_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::leaveEvent(static_cast<QEvent*>(event));
}

void MainWindow16b1d1_LowerDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::lower();
}

int MainWindow16b1d1_MetricDefault(void* ptr, long long m)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void MainWindow16b1d1_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void MainWindow16b1d1_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void MainWindow16b1d1_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void MainWindow16b1d1_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void MainWindow16b1d1_MoveEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::moveEvent(static_cast<QMoveEvent*>(event));
}

char MainWindow16b1d1_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void* MainWindow16b1d1_PaintEngineDefault(void* ptr)
{
	return static_cast<MainWindow16b1d1*>(ptr)->QDialog::paintEngine();
}

void MainWindow16b1d1_PaintEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::paintEvent(static_cast<QPaintEvent*>(event));
}

void MainWindow16b1d1_RaiseDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::raise();
}

void MainWindow16b1d1_RepaintDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::repaint();
}

void MainWindow16b1d1_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::setDisabled(disable != 0);
}

void MainWindow16b1d1_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::setEnabled(vbo != 0);
}

void MainWindow16b1d1_SetFocus2Default(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::setFocus();
}

void MainWindow16b1d1_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::setHidden(hidden != 0);
}

void MainWindow16b1d1_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void MainWindow16b1d1_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::setWindowModified(vbo != 0);
}

void MainWindow16b1d1_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void MainWindow16b1d1_ShowDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::show();
}

void MainWindow16b1d1_ShowFullScreenDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::showFullScreen();
}

void MainWindow16b1d1_ShowMaximizedDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::showMaximized();
}

void MainWindow16b1d1_ShowMinimizedDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::showMinimized();
}

void MainWindow16b1d1_ShowNormalDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::showNormal();
}

void MainWindow16b1d1_TabletEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::tabletEvent(static_cast<QTabletEvent*>(event));
}

void MainWindow16b1d1_UpdateDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::update();
}

void MainWindow16b1d1_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::updateMicroFocus();
}

void MainWindow16b1d1_WheelEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::wheelEvent(static_cast<QWheelEvent*>(event));
}

void MainWindow16b1d1_ChildEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::childEvent(static_cast<QChildEvent*>(event));
}

void MainWindow16b1d1_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void MainWindow16b1d1_CustomEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::customEvent(static_cast<QEvent*>(event));
}

void MainWindow16b1d1_DeleteLaterDefault(void* ptr)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::deleteLater();
}

void MainWindow16b1d1_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}



void MainWindow16b1d1_TimerEventDefault(void* ptr, void* event)
{
	static_cast<MainWindow16b1d1*>(ptr)->QDialog::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
