package gen

import (
	"fmt"
	"strconv"
)

var eti = 0

// Retorna nueva etiqueta
func NewEti() string {
	eti = eti + 1
	return "L" + strconv.Itoa(eti)
}

// imprime etiquetas destino
func Soltar(e []string) {
	for _, i := range e {
		Genln(i + ":")
	}
}

// une etiquetas
func Unir(e1 []string, e2 []string) []string {
	lista := e1
	for _, i := range e2 {
		lista = append(lista, i)
	}
	return lista
}

// mostrar etiquetas verdaderas y falsas
func Mostrar(lv []string, lf []string) {
	fmt.Println("\nEtiquetas verdaderas")
	Soltar(lv)
	fmt.Println("\nEtiquetas falsas")
	Soltar(lf)
}
