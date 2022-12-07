package biz

import (
	"cqccteg.hlw/up/common/logging"
	"github.com/oxiao/pongo2"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"io/ioutil"
	"os"

	//"os/exec"
	"path"
	"path/filepath"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *xorm.Engine
	DbName string
	tbls   []*core.Table
)

func DBInit(dbms, dbName, user, password, host, port string) (err error) {
	conn_str := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True"
	if dbms == "mssql" {
		conn_str = "sqlserver://" + user + ":" + password + "@" + host + ":" + port + "?database=" + dbName + ""
	}
	db, err = xorm.NewEngine(dbms, conn_str)
	if err != nil {
		logging.Err(err)
		return err
	}

	DbName = dbName

	return nil
}

func ModelGenerate(importName, tableName string, tplFile string) error {
	ls, err := db.DBMetas()
	if nil == err {
		tbls = make([]*core.Table, 0, len(ls))
		for _, tb := range ls {
			tbls = append(tbls, tb)
		}
	} else {
		logging.Err(err)
	}

	//modelData, err := ioutil.ReadFile(tplFile)
	//if err != nil {
	//	fmt.Println("read tplFile err:", err)
	//	return err
	//}

	render, err := pongo2.FromFile(tplFile)
	if err != nil {
		logging.Error("read tplFile err:", err)
		return err
	}

	//render := template.Must(template.New("model").
	//	Funcs(template.FuncMap{
	//		"FirstCharUpper":       FirstCharUpper,
	//		"TypeConvert":          TypeConvert,
	//		"ExportColumn":         ExportColumn,
	//		"Join":                 Join,
	//		"MakeQuestionMarkList": MakeQuestionMarkList,
	//		"ColumnAndType":        ColumnAndType,
	//		"ColumnWithPostfix":    ColumnWithPostfix,
	//		"Tags3":                Tags3,
	//		"ExportLabel":          ExportLabel,
	//	}).Parse(string(modelData)))

	subs := strings.Split(tplFile, "/")
	fileType := subs[len(subs)-1]
	fileType = strings.Split(fileType, ".")[0]
	fileType = strings.Split(fileType, "_")[1]

	nameStr := ""
	for _, table := range tbls {
		nameStr = nameStr + "&" + HumpStructName(table.Name) + "{}, "
		if (tableName == "") || (tableName != "" && tableName == table.Name) {
			//go func() {
			//	genModelFile(render, importName, &table, fileType)
			//}()
			err := genModelFile(render, importName, table, fileType)
			if err != nil {
				logging.Error("genModelFile err:", err)
				continue
			}
		}
	}
	//db.DumpTablesToFile(ls, getPath("all_tables.txt"))
	// write all model names to a single file
	err = ioutil.WriteFile("model_list.txt", []byte(nameStr), 0666)
	if err != nil {
		logging.Error("WriteFile model_list.txt err:", err)
	}

	return nil
}

type ModelInfo struct {
	*core.Table
	BDName          string
	DBConnection    string
	TableName       string
	ExportModelName string
	ImportName      string
	PackageName     string
	ModelName       string
	Columns         []*core.Column
	TableComment    string
}

func genModelFile(render *pongo2.Template, importName string, table *core.Table, fileType string) error {
	packageName := "local"
	if importName != "" {
		packageName = path.Base(importName)
	}

	dirPath := path.Base("") + string(filepath.Separator) + packageName + string(filepath.Separator) +
		fileType + string(filepath.Separator)
	if !IsExist(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			logging.Err(err)
			return err
		}
	}

	fileName := dirPath + strings.ToLower(table.Name) + "_auto." + fileType
	if IsExist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			logging.Err(err)
			return err
		}
	}

	//f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Println("os.Create error:", err.Error())
	//	return err
	//}
	//defer f.Close()

	model := &ModelInfo{
		Table:           table,
		ImportName:      importName,
		PackageName:     packageName,
		BDName:          DbName,
		ExportModelName: HumpStructName(table.Name),
		TableName:       table.Name,
		ModelName:       table.Name,
		Columns:         table.Columns(),
		TableComment:    table.Comment,
	}

	ctx := pongo2.Context{
		"FirstCharUpper":       FirstCharUpper,
		"TypeConvert":          TypeConvert,
		"ExportColumn":         ExportColumn,
		"Join":                 Join,
		"MakeQuestionMarkList": MakeQuestionMarkList,
		"ColumnAndType":        ColumnAndType,
		"ColumnWithPostfix":    ColumnWithPostfix,
		"Tags3":                Tags3,
		"ExportLabel":          ExportLabel,
		"md":                   model,
	}
	render.Options.TrimBlocks = true
	render.Options.LStripBlocks = true

	if str, err := render.Execute(ctx); err != nil {
		logging.Err(err)
		return err
	} else {
		if err = ioutil.WriteFile(fileName, []byte(str), 0666); nil != err {
			logging.Err(err)
			return err
		}
	}

	//if "go" == fileType {
	//	cmd := exec.Command("goimports", "-w", fileName)
	//	err = cmd.Run()
	//	if err != nil {
	//		fmt.Println("format go code error:", err.Error())
	//		return err
	//	}
	//}

	return nil
}

//判断是否存在文件或者目录
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
