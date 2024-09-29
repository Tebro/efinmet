package utils

import "strings"


func FieldIcaoHasPrefix(prefix string, fieldIcao string) bool {
	return strings.HasPrefix(fieldIcao, prefix)
}
