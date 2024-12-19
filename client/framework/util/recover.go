package util

import (
	"bytes"
	"fmt"
	"framework/log"
	"go.uber.org/zap"
	"runtime"
)

func GetStackInfo() string {
	kb := 4

	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")

	return string(stack)
}

func PanicErrStack() {
	if err := recover(); err != nil {
		errorStr := GetStackInfo()
		fmt.Println(errorStr)
		log.Error("PanicErrStack", zap.Reflect("err", err))
		//return errors.New(errorStr)
	}
}
