package utils

import (
	"OperationAndMonitoring/utils/convert"
	"fmt"
	"reflect"
	"strings"
)

func IsNull(value interface{}) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Ptr:
		return reflect.ValueOf(value).IsNil()
	case reflect.String:
		return len(strings.TrimSpace(convert.ToString(value))) == 0
	case reflect.Uint64:
		return convert.ToUint64(value) != 0
	default:
		fmt.Println("default")
		return false
	}
}
