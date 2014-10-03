package api

import (
	"fmt"
)

func writeString(input string) string {
	return fmt.Sprintf("%s", input)
}

func writeNumeric(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

func writeKeyValue(key string, value interface{}) string {
	return fmt.Sprintf("%s=%v", key, value)
}
