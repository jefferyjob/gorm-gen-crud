package crud_tmp

const CrudTmp = `
package main

import (
	"errors"

	"gorm.io/gorm"
)

type {{.Name}} struct {
{{range .Fields}}	{{.Name}} {{.Type}} ` + "`gorm:\"column:{{lower .Name}};type:{{.Type}};{{.GormOption}}\" json:\"{{.JsonTag}}\"`" + `
{{end}}}

func Create{{.Name}}(db *gorm.DB, {{.Name}}Data *{{.Name}}) error {
	result := db.Create({{.Name}}Data)
	if result.Error != nil {
		return errors.New("failed to create {{.Name}}")
	}
	return nil
}

func Update{{.Name}}(db *gorm.DB, id int, updateData map[string]interface{}) error {
	result := db.Model(&{{.Name}}{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		return errors.New("failed to update {{.Name}}")
	}
	if result.RowsAffected == 0 {
		return errors.New("{{.Name}} not found")
	}
	return nil
}

func Delete{{.Name}}(db *gorm.DB, id int) error {
	result := db.Where("id = ?", id).Delete(&{{.Name}}{})
	if result.Error != nil {
		return errors.New("failed to delete {{.Name}}")
	}
	if result.RowsAffected == 0 {
		return errors.New("{{.Name}} not found")
	}
	return nil
}

func Get{{.Name}}(db *gorm.DB, id int) (*{{.Name}}, error) {
	var {{.Name}}Data {{.Name}}
	result := db.Where("id = ?", id).First(&{{.Name}}Data)
	if result.Error != nil {
		return nil, errors.New("failed to get {{.Name}}")
	}
	return &{{.Name}}Data, nil
}
`
