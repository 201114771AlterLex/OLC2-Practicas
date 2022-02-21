package listener

import (
	"strconv"

	"github.com/OLC2-Practicas/Evaluacion1/parser"
)

type Listener struct {
	*parser.BaseEvaluacion1Listener
	Resultado string
}

func (l *Listener) ExitDosNumero(ctx *parser.DosNumeroContext) {
	s1 := strconv.FormatFloat(ctx.H(0).GetHDecimal(), 'f', -1, 64)
	s2 := strconv.FormatFloat(ctx.H(1).GetHDecimal(), 'f', -1, 64)
	l.Resultado = s1 + "," + s2
}

func (l *Listener) ExitHTPT(ctx *parser.HTPTContext) {
	PEntera := float64(ctx.T(0).GetDhex().Valor)
	PDecimal := (float64(ctx.T(1).GetDhex().Valor) / float64(ctx.T(1).GetDhex().Posicion*16))
	ctx.HDecimal = PEntera + PDecimal
	//fmt.Println(ctx.HDecimal)
}

func (l *Listener) ExitHT(ctx *parser.HTContext) {
	ctx.HDecimal = float64(ctx.T().GetDhex().Valor)
}

func (l *Listener) ExitTHex(c *parser.THexContext) {
	TextHex := c.Hex().GetText()
	Num := 0
	switch TextHex {
	case "1":
		Num = 1
	case "2":
		Num = 2
	case "3":
		Num = 3
	case "4":
		Num = 4
	case "5":
		Num = 5
	case "6":
		Num = 6
	case "7":
		Num = 7
	case "8":
		Num = 8
	case "9":
		Num = 9
	case "A":
		Num = 10
	case "B":
		Num = 11
	case "C":
		Num = 12
	case "D":
		Num = 13
	case "E":
		Num = 14
	case "F":
		Num = 15
	default:
		Num = 0
	}
	//fmt.Println(Num)
	HeredarDhex := new(parser.DataHex)
	HeredarDhex.Valor = Num*c.Tp().GetDhex().Posicion + c.Tp().GetDhex().Valor
	HeredarDhex.Posicion = c.Tp().GetDhex().Posicion
	c.SetDhex(HeredarDhex)
}

func (l *Listener) ExitTpHex(ctx *parser.TpHexContext) {
	TextHex := ctx.Hex().GetText()
	Num := 0
	switch TextHex {
	case "1":
		Num = 1
	case "2":
		Num = 2
	case "3":
		Num = 3
	case "4":
		Num = 4
	case "5":
		Num = 5
	case "6":
		Num = 6
	case "7":
		Num = 7
	case "8":
		Num = 8
	case "9":
		Num = 9
	case "A":
		Num = 10
	case "B":
		Num = 11
	case "C":
		Num = 12
	case "D":
		Num = 13
	case "E":
		Num = 14
	case "F":
		Num = 15
	default:
		Num = 0
	}
	//fmt.Println(Num)
	HeredarDhex := new(parser.DataHex)
	HeredarDhex.Valor = Num*ctx.Tp().GetDhex().Posicion + ctx.Tp().GetDhex().Valor
	HeredarDhex.Posicion = ctx.Tp().GetDhex().Posicion * 16
	ctx.SetDhex(HeredarDhex)
}

func (l *Listener) ExitEpTp(ctx *parser.EpTpContext) {
	HeredarDhex := new(parser.DataHex)
	HeredarDhex.Posicion = 1
	ctx.SetDhex(HeredarDhex)
}
