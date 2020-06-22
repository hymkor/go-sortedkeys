package sortedkeys

import (
	"reflect"
	"sort"
)

// Strings makes sorted strings' array from keys of the given map whose key's type is string.
func Strings(map1 interface{}) []string {
	values := reflect.ValueOf(map1).MapKeys()
	result := make([]string, len(values))
	for i, value1 := range values {
		result[i] = value1.String()
	}
	sort.Strings(result)
	return result
}
