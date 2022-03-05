package gen

import (
	"fmt"
)

// funcion para generar cadena
func Gen(out string) {
	fmt.Print(out)
}

func Genln(out string) {
	fmt.Println(out)
}

// Genera operacion unaria
func GenOperacionUnaria(dir1 string, op string, dir2 string) {
	Genln(dir1 + " = " + op + " " + dir2)
}

// Genera operacion binaria
func GenOperacion(dir1 string, dir2 string, op string, dir3 string) {
	Genln(dir1 + " = " + dir2 + " " + op + " " + dir3)
}

// Genera instruccion condicional
// retorna la operacion
func GenIf(dir1 string, op string, dir2 string, lv string) string {
	cad := dir1 + " " + op + " " + dir2
	Genln("if " + cad + " then goto " + lv)
	return cad
}

func GenIfCad(cad string, lv string) {
	Genln("if " + cad + " then goto " + lv)
}

// Genera instruccion incondicional
func GenGoto(eti string) {
	Genln("goto " + eti)
}

// Genera etiqueta destino
func GenDestino(eti string) {
	// L1:
	Genln(eti + ":")
}

// Genera asignacion
func GenAsignacion(dir1 string, dir2 string) {
	Genln(dir1 + " = " + dir2)
}

func GenSetArreglo(dir1 string, pos string, dir2 string) {
	Genln(dir1 + "[" + pos + "] = " + dir2)
}

func GenGetArreglo(dir1 string, dir2 string, pos string) {
	Genln(dir1 + " = " + dir2 + "[" + pos + "]")
}
