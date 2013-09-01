package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func ValueAsString(v interface{}) string {
	return fmt.Sprintf("%#v (%T)", v, v)
}

type nilValueType interface{}

var nilValue *nilValueType = nil

func AppendValueFor(array []reflect.Value, obj interface{}) []reflect.Value {
	var value reflect.Value
	if reflect.TypeOf(obj) == nil {
		value = reflect.ValueOf(nilValue)
	} else {
		value = reflect.ValueOf(obj)
	}
	return append(array, value)
}

func GetStackTrace(offset int) string {
	strstack := make([]string, 0)
	stack := make([]uintptr, 100)
	count := runtime.Callers(offset, stack)
	stack = stack[0:count]
	for i, pc := range stack {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		if i == 0 {
			line--
		}
		strstack = append(strstack, fmt.Sprintf("%s:%d\n    inside %s", file, line, fn.Name()))
	}
	return strings.Join(strstack, "\n")
}

func Tabulate(prefix, content, sep string) string {
	lines := strings.Split(content, sep)
	buffer := strings.Repeat(" ", len(prefix))
	for i, line := range lines {
		if i != 0 {
			lines[i] = buffer + line
		}
	}
	return prefix + strings.Join(lines, sep)
}
