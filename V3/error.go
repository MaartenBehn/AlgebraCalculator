package V3

import (
	log "github.com/sirupsen/logrus"
	"os"
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

	if reflect.TypeOf(err) == reflect.TypeOf(&calculatorError{}) {
		switch err.(*calculatorError).critical {
		case errorCriticalLevelNon:
			log.Print(err)
			log.Println(" -> Programm execution was continued.")
			return false
		case errorCriticalLevelPartial:
			log.Print(err)
			log.Println(" -> Some parts of the programm where not executed.")
			return true
		case errorCriticalLevelFatal:
			log.Print(err)
			log.Println(" -> Programm was stopped!")
			os.Exit(1)
		}
		handelError(newError(errorTypErrorhandling, errorCriticalLevelFatal, "calculatorError "+err.(*calculatorError).text+" has not valid critiacl level!"))
	} else {
		handelError(newError(errorTypExternal, errorCriticalLevelFatal, err.Error()))
	}
	return false
}
