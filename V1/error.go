package V1

import (
	"AlgebraCalculator/V1/log"
	"reflect"
	"strings"
)

const (
	errorTypExternal      = "External Error"
	errorTypParsing       = "Parsing Error"
	errorTypSolving       = "Solving Error"
	errorTypSimplifying   = "Simplifying Error"
	errorTypErrorhandling = "Errorhandling Error"

	errorCriticalLevelNon     = "non critical"
	errorCriticalLevelPartial = "partial critical"
	errorCriticalLevelFatal   = "fatal"
)

type calculatorError struct {
	text     string
	typ      string
	critical string
}

func newError(typ string, critical string, text string) *calculatorError {
	return &calculatorError{
		text:     text,
		typ:      typ,
		critical: critical,
	}
}

func (e calculatorError) Error() string {
	builder := strings.Builder{}
	builder.WriteString(e.typ)
	builder.WriteString(": \" ")
	builder.WriteString(e.text)
	builder.WriteString(" \" ")
	return builder.String()
}

func handelError(err error) bool {

	if err == nil {
		return false
	}

	if reflect.TypeOf(err) == reflect.TypeOf(&calculatorError{}) {
		switch err.(*calculatorError).critical {
		case errorCriticalLevelNon:
			log.Print(err.Error())
			return false
		case errorCriticalLevelPartial:
			log.Print(err.Error())
			return true
		case errorCriticalLevelFatal:
			panic(err)
		}
		handelError(newError(errorTypErrorhandling, errorCriticalLevelFatal, "calculatorError \""+err.(*calculatorError).text+"\" has not valid critiacl level!"))
	} else {
		handelError(newError(errorTypExternal, errorCriticalLevelFatal, err.Error()))
	}
	return false
}
