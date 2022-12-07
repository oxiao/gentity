package test

import (
	"fmt"
	"github.com/oxiao/gentity/biz"
	"os"
	"path/filepath"
	"testing"
)

func TestGen(t *testing.T) {
	dbms := "mssql"
	dbName := "KJ190A"
	user := "kj"
	password := "1"
	host := "192.168.20.31"
	port := "1433"
	importName := "abc"
	tableName := ""
	gen(dbms, dbName, user, password, host, port, importName, tableName)
}

func gen(dbms, dbName, user, password, host, port, importName, tableName string) {
	var (
		err error
	)
	err = biz.DBInit(dbms, dbName, user, password, host, port)
	fmt.Println(err)
	err = biz.ModelGenerate(importName, tableName, "/usr/local/gopath/src/github.com/oxiao/gentity/tpl/model_go.tmpl")
	fmt.Println(err)
	err = biz.ModelGenerate(importName, tableName, "/usr/local/gopath/src/github.com/oxiao/gentity/tpl/model_js.tmpl")
	fmt.Println(err)
}
func getPath(name string) string {
	mp, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	mp = filepath.Join(mp, name)
	return mp
}
