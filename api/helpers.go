package api

import (
	"fmt"
	"strings"
)

func writeString(input string) string {
	return fmt.Sprintf("%s", input)
}

func writeNumeric(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

func writeKeyValue(key string, value interface{}) string {
	return fmt.Sprintf("%s=%v", key, escape(value))
}

func escape(value interface{}) interface{} {
	switch v := value.(type) {
	case string:
		tmp := strings.Replace(v, "|", "&#124;", 1)
		tmp = strings.Replace(v, "=", "&#61;", 1)
		return tmp
	}
	return value
}
