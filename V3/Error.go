package V3

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

const (
	ErrorTypExternal      = "External Error"
	ErrorTypParsing       = "Parsing Error"
	ErrorTypSolving       = "Solving Error"
	ErrorTypSimplifying   = "Simplifying Error"
	ErrorTypErrorhandling = "Errorhandling Error"

	ErrorCriticalLevelNot     = "non critical"
	ErrorCriticalLevelPartial = "partial critical"
	ErrorCriticalLevelFatal   = "fatal"
)

type Error struct {
	text     string
	typ      string
	critical string
}

func NewError(typ string, critical string, text string) *Error {
	return &Error{
		text:     text,
		typ:      typ,
		critical: critical,
	}
}

func (e Error) Error() string {
	builder := strings.Builder{}
	builder.WriteString(e.typ)
	builder.WriteString(": \"")
	builder.WriteString(e.text)
	builder.WriteString("\" with critical Level: ")
	builder.WriteString(e.critical)
	builder.WriteString(" accured!")
	return builder.String()
}

func handelError(err error) bool {

	if err == nil {
		return false
	}

	if reflect.TypeOf(err) == reflect.TypeOf(&Error{}) {
		switch err.(*Error).critical {
		case ErrorCriticalLevelNot:
			fmt.Print(err)
			fmt.Println(" -> Programm execution was continued.")
			return false
		case ErrorCriticalLevelPartial:
			fmt.Print(err)
			fmt.Println(" -> Some parts of the programm where not executed.")
			return true
		case ErrorCriticalLevelFatal:
			fmt.Print(err)
			fmt.Println(" -> Programm was stopped!")
			os.Exit(1)
		}
		handelError(NewError(ErrorTypErrorhandling, ErrorCriticalLevelFatal, "Error "+err.(*Error).text+" has not valid critiacl level!"))
	} else {
		handelError(NewError(ErrorTypExternal, ErrorCriticalLevelFatal, err.Error()))
	}
	return false
}
