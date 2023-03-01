package main

import (
	"fmt"
	crud_tmp "github.com/jefferyjob/gorm-gen-crud/crud-tmp"
	"os"
	"strings"
	"text/template"
)

type Field struct {
	Name       string
	Type       string
	GormOption string
	JsonTag    string
}

type StructData struct {
	Name   string
	Fields []Field
}

func main() {

	// 定义要生成的结构体
	structData := StructData{
		Name: "GoddessWithdrawal",
		Fields: []Field{
			{Name: "Id", Type: "int", GormOption: "primaryKey;AUTO_INCREMENT", JsonTag: "id"},
			{Name: "Uid", Type: "int", GormOption: "NOT NULL", JsonTag: "uid"},
			{Name: "Balance", Type: "int", GormOption: "comment:金额(分);NOT NULL", JsonTag: "balance"},
			{Name: "Date", Type: "string", GormOption: "type:date;NOT NULL", JsonTag: "date"},
			{Name: "CreatedAt", Type: "string", GormOption: "type:datetime;NOT NULL", JsonTag: "created_at"},
		},
	}

	// 定义要使用的模板
	tmpl := template.Must(template.New("").Parse(crud_tmp.CrudTmp))

	// 执行模板并生成代码文件
	f, err := os.Create("crud_" + structData.Name + ".go")
	if err != nil {
		fmt.Println("failed to create file: ", err)
		return
	}
	defer f.Close()

	funcMap := template.FuncMap{
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
	}

	err = tmpl.Funcs(funcMap).Execute(f, structData)
	if err != nil {
		fmt.Println("failed to generate code: ", err)
		return
	}

	fmt.Println("generated crud code for", structData.Name, "successfully")
}
