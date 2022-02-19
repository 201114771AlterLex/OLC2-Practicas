
//
grammar Practica4;

PRight: '(';
PLeft: ')';
Minus: '-';
Colon: ',';
Numbers:Minus?[0-9]+;
Whitespace: [ \r\n\t]+ -> skip;

start: coordinate coordinates EOF #Inicio;

coordinates
    :coordinate coor #C1
    ;

coor
    :coordinate coor #C2
    | #C3
    ;

coordinate    : PRight Numbers Colon Numbers PLeft #Cordenada;