package listener

import (
	"fmt"

	"github.com/OLC2-Practicas/Practica3/parser"
)

type Listener struct {
	*parser.BasePractica3Listener
	Type  string
	Value string
}

func (l *Listener) ExitStart(ctx *parser.StartContext) {
	fmt.Println("Listener Resultado")
	l.Type = ctx.E().GetD1().DataType
	l.Value = ctx.E().GetD1().DataValue
}
