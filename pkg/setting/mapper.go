// code from https://github.com/go-xorm/core/blob/master/mapper.go

package setting

import (
	"strings"
)

var gm = GonicMapper{}

func mapper(name string) string {
	return gm.Obj2Table(name)
}

// GonicMapper implements IMapper. It will consider initialisms when mapping names.
// E.g. id -> ID, user -> User and to table names: UserID -> user_id, MyUID -> my_uid
type GonicMapper map[string]bool

func isASCIIUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func toASCIIUpper(r rune) rune {
	if 'a' <= r && r <= 'z' {
		r -= ('a' - 'A')
	}
	return r
}

func gonicCasedName(name string) string {
	newstr := make([]rune, 0, len(name)+3)
	for idx, chr := range name {
		if isASCIIUpper(chr) && idx > 0 {
			if !isASCIIUpper(newstr[len(newstr)-1]) {
				newstr = append(newstr, '_')
			}
		}

		if !isASCIIUpper(chr) && idx > 1 {
			l := len(newstr)
			if isASCIIUpper(newstr[l-1]) && isASCIIUpper(newstr[l-2]) {
				newstr = append(newstr, newstr[l-1])
				newstr[l-1] = '_'
			}
		}

		newstr = append(newstr, chr)
	}
	return strings.ToLower(string(newstr))
}

func (mapper GonicMapper) Obj2Table(name string) string {
	return gonicCasedName(name)
}

func (mapper GonicMapper) Table2Obj(name string) string {
	newstr := make([]rune, 0)

	name = strings.ToLower(name)
	parts := strings.Split(name, "_")

	for _, p := range parts {
		_, isInitialism := mapper[strings.ToUpper(p)]
		for i, r := range p {
			if i == 0 || isInitialism {
				r = toASCIIUpper(r)
			}
			newstr = append(newstr, r)
		}
	}

	return string(newstr)
}

// A GonicMapper that contains a list of common initialisms taken from golang/lint
var LintGonicMapper = GonicMapper{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,

	//my values
	"BTC": true,
	"RUB": true,
	"USD": true,
}
