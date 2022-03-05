package err

import (
	"fmt"
)

type Error struct {
	Mensaje string
}

var lista []*Error

func NewError(msg string) {
	lista = append(lista, &Error{Mensaje: msg})
}

func Mostrar() {
	for _, err := range lista {
		fmt.Println("Mensaje: " + err.Mensaje)
	}
}
