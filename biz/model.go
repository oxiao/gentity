/*
 *
 *
 *  * Copyright 2019 koalaone@163.com
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *       http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package biz

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instance *gorm.DB
var instanceMysql *gorm.DB

var DbName string

func DBInit(dbName, user, password, host, port string) (err error) {
	conn_str := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True"
	instance, err = gorm.Open("mysql", conn_str)
	if err != nil {
		return err
	}

	mysql_conn_str := user + ":" + password + "@tcp(" + host + ":" + port + ")/information_schema?charset=utf8mb4&parseTime=True"
	instanceMysql, err = gorm.Open("mysql", mysql_conn_str)
	if err != nil {
		return err
	}

	DbName = dbName

	return nil
}

type TableInfo struct {
	TableName    string `json:"table_name"`
	TableComment string `json:"table_comment"`
}

func ModelGenerate(importName, tableName string, tplFile string) error {
	var tablaNames []TableInfo
	err := instanceMysql.Raw(`SELECT table_name as table_name, TABLE_COMMENT as table_comment from tables where table_schema = ? `, DbName).Find(&tablaNames).Error
	if err != nil {
		return err
	}

	modelData, err := ioutil.ReadFile(tplFile)
	if err != nil {
		fmt.Println("read tplFile err:", err)
		return err
	}

	render := template.Must(template.New("model").
		Funcs(template.FuncMap{
			"FirstCharUpper":       FirstCharUpper,
			"TypeConvert":          TypeConvert,
			"Tags":                 Tags,
			"ExportColumn":         ExportColumn,
			"Join":                 Join,
			"MakeQuestionMarkList": MakeQuestionMarkList,
			"ColumnAndType":        ColumnAndType,
			"ColumnWithPostfix":    ColumnWithPostfix,
			"Tags3":                Tags3,
			"ExportLabel":          ExportLabel,
		}).Parse(string(modelData)))

	subs := strings.Split(tplFile, "/")
	fileType := subs[len(subs)-1]
	fileType = strings.Split(fileType, ".")[0]
	fileType = strings.Split(fileType, "_")[1]

	nameStr := ""
	for _, table := range tablaNames {
		nameStr = nameStr + "&" + HumpStructName(table.TableName) + "{}, "
		if (tableName == "") || (tableName != "" && tableName == table.TableName) {
			err := genModelFile(render, importName, &table, fileType)
			if err != nil {
				fmt.Println("genModelFile err:", err)
				return err
			}
		}
	}

	// write all model names to a single file
	err = ioutil.WriteFile("model_list.txt", []byte(nameStr), 0666)
	if err != nil {
		fmt.Println("WriteFile model_list.txt err:", err)
	}

	return nil
}

type ModelInfo struct {
	BDName          string
	DBConnection    string
	TableName       string
	ExportModelName string
	ImportName      string
	PackageName     string
	ModelName       string
	TableSchema     *[]TableSchema
	TableComment    string
}

type TableSchema struct {
	ColumnName    string `db:"column_name" json:"column_name"`
	DataType      string `db:"data_type" json:"data_type"`
	ColumnType    string `db:"column_type" json:"column_type"`
	ColumnKey     string `db:"column_key" json:"column_key"`
	ColumnComment string `db:"column_comment" json:"column_comment"`
}

func genModelFile(render *template.Template, importName string, table *TableInfo, fileType string) error {
	if table == nil {
		return nil
	}
	tableName := table.TableName

	var tableSchema []TableSchema
	err := instanceMysql.Raw(`SELECT column_name as column_name, column_type as column_type, data_type as data_type,column_key as column_key,`+
		`column_comment as column_comment from COLUMNS `+
		` where TABLE_NAME= ? and table_schema = ? and column_name not in ('created_at', 'updated_at', 'deleted_at')`, tableName, DbName).Find(&tableSchema).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(tableSchema) <= 0 {
		fmt.Println(tableName, "tableSchema ["+tableName+"] is null")
		return errors.New("tableSchema [" + tableName + "] is null")
	}

	packageName := "local"
	if importName != "" {
		packageName = path.Base(importName)
	}

	dirPath := path.Base("") + string(filepath.Separator) + packageName + string(filepath.Separator) +
		fileType + string(filepath.Separator)
	if !IsExist(dirPath) {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	fileName := dirPath + strings.ToLower(tableName) + "_auto." + fileType
	if IsExist(fileName) {
		err = os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("os.Create error:", err.Error())
		return err
	}
	defer f.Close()

	model := &ModelInfo{
		ImportName:      importName,
		PackageName:     packageName,
		BDName:          DbName,
		ExportModelName: HumpStructName(tableName),
		TableName:       tableName,
		ModelName:       tableName,
		TableSchema:     &tableSchema,
		TableComment:    table.TableComment,
	}

	if err := render.Execute(f, model); err != nil {
		log.Fatal(err)
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
