grammar Gramatica;
/*
E		E + E
	|	E – E
	|	E * E
	|	E / E
	|	( E )
	|	NUM

*/

options {
    language='Go';
}

@parser::header {
    import "github.com/OLC2-Practicas/Practica5/generador"
    import "github.com/OLC2-Practicas/Practica5/entorno"
}

@parser::members{
    type Nodo struct {
        dir string
        tipo string
    }

    var desp int = 0
    var sup *entorno.Entorno

    func gen(out string) {
        fmt.Println(out);
    }
}

start
    : marcador dec  { entorno.Mostrar(sup) }
    | exp EOF
    ;

marcador
    : { sup = entorno.NewEntorno(nil); desp = 0; }
    ;

dec
    : t=tipo id=ID                  {
                                        sim := entorno.NewSimbolo($id.text, $t.tipoDato, desp)
                                        sup.Put($id.text, sim)
                                        desp = desp + 1
                                    }
        ';' dec
    | 
    ;

tipo returns[ string tipoDato ]
    : t='int'                       {
                                        $tipoDato = $t.text
                                    }
    ;

exp returns [ *Nodo nodo ]
    : e1=exp op=('*'|'/') e2=exp    {
                                        $nodo = &Nodo { }
                                        $nodo.dir = generador.NewTemp()
                                        cad := $nodo.dir + " = " + $e1.nodo.dir + " " + $op.text + " "
                                        cad += $e2.nodo.dir
                                        gen(cad)
                                    }
    | e1=exp op=('+'|'-') e2=exp    {
                                        $nodo = &Nodo { }
                                        $nodo.dir = generador.NewTemp()
                                        cad := $nodo.dir + " = " + $e1.nodo.dir + " " + $op.text + " "
                                        cad += $e2.nodo.dir
                                        gen(cad)
                                    }
    | '(' e=exp ')'                 {
                                        $nodo = &Nodo { }
                                        $nodo.dir = $e.nodo.dir
                                    }
    | n=NUM                         {
                                        $nodo = &Nodo { }
                                        $nodo.dir = $n.text
                                    }
    ;


// Tokens
MAS: '+';
MEN: '-';
POR: '*';
DIV: '/';
PARI: '(';
PARD: ')';
NUM: [0-9]+;
ID: [_A-Za-z][_A-Za-z0-9]*;
WHITESPACE: [ \r\n\t]+ -> skip;