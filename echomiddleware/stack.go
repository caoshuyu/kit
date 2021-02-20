package echomiddleware

import (
	"bytes"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func callers(skip int) []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip+1, pcs[:])
	return pcs[:n]
}

func stackString(stack []uintptr) string {
	if len(stack) == 0 {
		return ""
	}
	frames := runtime.CallersFrames(stack)

	var (
		frame    runtime.Frame
		more     bool
		funcName string
		fileName string
		buf      bytes.Buffer
	)
	for {
		frame, more = frames.Next()

		switch frame.Function {
		case "runtime.main":
			goto ERRORLOOP
		case "runtime.goexit":
			goto ERRORLOOP
		}

		if !strings.EqualFold(frame.Function, "") {
			funcName = trimFuncName(frame.Function)
		} else {
			funcName = "unknown_function"
		}
		if !strings.EqualFold(frame.File, "") {
			fileName = trimFileName(frame.File)
		} else {
			fileName = "unknown_file"
		}

		if buf.Len() > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(funcName)
		buf.WriteString("\n\t")
		buf.WriteString(fileName)
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(frame.Line))

		if !more {
			goto ERRORLOOP
		}
	}

ERRORLOOP:

	return buf.String()
}

func trimFileName(name string) string {
	i := strings.Index(name, "/src/")
	if i < 0 {
		return name
	}
	name = name[i+len("/src/"):]
	i = strings.Index(name, "/vendor/")
	if i < 0 {
		return name
	}
	return name[i+len("/vendor/"):]
}

func trimFuncName(name string) string {
	return path.Base(name)
}
