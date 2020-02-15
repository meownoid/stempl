package stempl

import (
	"strings"

	"github.com/pkg/errors"
)

func Format(format string, values map[string]string) (string, error) {
	var result strings.Builder

begin:
	for i, a := range format {
		if a == '{' {
			key := ""
			k := -1
		search:
			for j, b := range format[i+1:] {
				if b == '{' {
					if j == 0 {
						result.WriteRune('{')
						format = format[i+2:]
						goto begin
					} else {
						return "", errors.Errorf("unexpected { in field name (%d)", j)
					}
				}
				if b == '}' {
					key = format[i+1 : i+j+1]
					k = i + j + 2
					break search
				}
			}
			if k == -1 {
				return "", errors.Errorf("single { encountered in format string (%d)", i)
			}
			if key == "" {
				return "", errors.Errorf("empty field (%d)", i)
			}
			if value, ok := values[key]; ok {
				result.WriteString(value)
			} else {
				return "", errors.Errorf("no value for key %s (%d)", key, i)
			}
			format = format[k:]
			goto begin
		}
		if a == '}' {
			if i+1 < len(format) && format[i+1] == '}' {
				result.WriteRune('}')
				format = format[i+2:]
				goto begin
			} else {
				return "", errors.Errorf("single } encountered in format string (%d)", i)
			}
		}
		result.WriteRune(a)
	}

	return result.String(), nil
}
