package sql2struct

import (
	"fmt"
	"github.com/foxboolean/go_tour/go_tools/internal/word"
	"os"
	"text/template"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

// StructTemplateDB 渲染模板的 data
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	columns := make([]*StructColumn, 0, len(tbColumns))
	for _, col := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", col.ColumnName)
		column := &StructColumn{
			Name:    col.ColumnName,
			Type:    DBTypeToStructType[col.DataType],
			Tag:     tag,
			Comment: col.ColumnComment,
		}
		columns = append(columns, column)
	}
	return columns
}

// Generate 渲染模板，注意引入 text/template 否则会转义
func (t *StructTemplate) Generate(tbName string, cols []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))
	tplDB := StructTemplateDB{
		TableName: tbName,
		Columns:   cols,
	}
	return tpl.Execute(os.Stdout, tplDB)
}
