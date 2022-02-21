
//
grammar Evaluacion1;

options{
    language='Go';
}

@parser::members{
    type DataHex struct {
        Valor    int
        Posicion int
    }

    type Decimal struct {
        Decimal1 float64
        Decimal2 float64
    }


}

Point: '.';
Colon: ',';
Hex: Number | Letter;
Number:[0-9];
Letter:[A-F];
Whitespace: [ \r\n\t]+ -> skip;

start: list EOF #Inicio;

list: h Colon h #DosNumero
;

h  returns [ float64 HDecimal ]
    : t Point t #HTPT
    |t #HT
;

t returns [ *DataHex Dhex]
    : Hex tp #THex
;

tp returns [*DataHex Dhex]
    : Hex tp  #TpHex
    |#EpTp
;