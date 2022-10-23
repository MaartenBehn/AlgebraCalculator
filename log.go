package AlgebraCalculator

import (
	"bytes"
	"fmt"
)

var logBuffer *bytes.Buffer

func InitLog() {
	logBuffer = bytes.NewBuffer(nil)
}

func GetLog() string {
	text := logBuffer.String()
	logBuffer.Reset()
	return text
}

func PrintLog() {
	text := logBuffer.String()
	logBuffer.Reset()
	fmt.Print(text)
}

func Print(string string) {
	logBuffer.WriteString(string)
}

func Println(string string) {
	Print(string + "\n")
}

func Printf(format string, a ...interface{}) {
	Print(fmt.Sprintf(format, a...))
}
