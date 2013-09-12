package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const DefaultValueStringMax = 128

var ValueStringMax int = DefaultValueStringMax

func ValueAsString(v interface{}) string {
	value := fmt.Sprintf("%#v", v)
	if ValueStringMax > 0 {
		if len(value) > ValueStringMax {
			value = value[:ValueStringMax-5] + "<...>" + value[len(value)-128+5:]
		}
	}
	return fmt.Sprintf("%s (%T)", value, v)
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

// Returns a stacktrace string to the invocation of GetStackTrace()
// Due to the inaccuracys of runetime.Callers, this line number may
// be off by one
func GetStackTrace(offset int) string {
	strstack := make([]string, 0)
	stack := make([]uintptr, 100)
	/*
		_, file, line, _ := runtime.Caller(offset + 1)
		fmt.Printf("\n\n%s:%d\n", file, line)
	*/
	count := runtime.Callers(offset+2, stack)
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
