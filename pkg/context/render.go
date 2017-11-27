package context

import (
	"bytes"
	"log"
	"strconv"
	"text/template"
)

func (c *Context) Render(name string, datas ...map[string]interface{}) string {
	data := c.Data
	if len(datas) > 0 {
		data = datas[0]
	}
	tpl := c.Tr(name)
	tmpl, err := template.New("iop").Funcs(map[string]interface{}{
		"fl": func(fl float64) string {
			return strconv.FormatFloat(fl, 'f', -1, 32)
		},
	}).Parse(tpl)
	if err != nil {
		log.Fatalln(err)
	}
	var buf = bytes.NewBuffer(nil)
	err = tmpl.Execute(buf, data)
	if err != nil {
		log.Fatalln(err)
	}
	return buf.String()
}
