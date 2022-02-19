package listener

import (
	"math"
	"strconv"

	"github.com/OLC2-Practicas/Practica4/parser"
)

type Listener struct {
	*parser.BasePractica4Listener
	Distance    float64
	Coordinate1 [2]float64
	Coordinate2 [2]float64
}

func (l *Listener) ExitCordenada(c *parser.CordenadaContext) {
	x, err := strconv.Atoi(c.Numbers(0).GetText())
	if err != nil {
		panic(err.Error())
	}
	y, err := strconv.Atoi(c.Numbers(1).GetText())
	if err != nil {
		panic(err.Error())
	}
	if l.Coordinate1[0] == 0 && l.Coordinate1[1] == 0 && l.Coordinate2[0] == 0 && l.Coordinate2[1] == 0 {
		l.Coordinate1[0] = float64(x)
		l.Coordinate1[1] = float64(y)
		l.Distance = 0
	} else if l.Coordinate1[0] != 0 && l.Coordinate1[1] != 0 && l.Coordinate2[0] == 0 && l.Coordinate2[1] == 0 {
		l.Coordinate2[0] = float64(x)
		l.Coordinate2[1] = float64(y)
		l.Distance = l.Distance + math.Sqrt(math.Pow(l.Coordinate2[0]-l.Coordinate1[0], 2)+math.Pow(l.Coordinate2[1]-l.Coordinate1[1], 2))
	} else {
		l.Coordinate1[0] = l.Coordinate2[0]
		l.Coordinate1[1] = l.Coordinate2[1]
		l.Coordinate2[0] = float64(x)
		l.Coordinate2[1] = float64(y)
		l.Distance = l.Distance + math.Sqrt(math.Pow(l.Coordinate2[0]-l.Coordinate1[0], 2)+math.Pow(l.Coordinate2[1]-l.Coordinate1[1], 2))
	}

}
