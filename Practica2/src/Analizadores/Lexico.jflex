//* ------------Seccion codigo-usuario -------- */
package Analizadores;

import java_cup.runtime.*;
import java.util.ArrayList;

%%
/*- Seccion de opciones y declaraciones -*/
/*-OPCIONES --*/
/* Nombre de la clase generada para el analizador lexico */
%class Lexico
/*Si se quiere que no se tome en cuenta la diferencia entre la mayuscula y la minuscula*/
//%ignorecase
%caseless 
/*Para que el lector del archivo tenga la cadena de 16 bits*/
%unicode
/*Para que las clases sean publicas*/
%public
/* Indicar funcionamiento autonomo, crea un metodo main para que solo sea llamando*/
//%standalone
/* Acceso a la columna y fila actual de analisis CUP */
%line
%column
/*El nombre de la clase que cup va a generar*/
%cupsym Simbolos
/* Habilitar la compatibilidad con el interfaz CUP para el generador sintactico*/
%cup


//%function next_token
//%type java_cup.runtime.Symbol


%eofval{
	//System.out.println("Fin eofval");
        return new Symbol(Simbolos.EOF);
        //return null;
%eofval}
%eof{
    //System.out.println("Fin eof");
%eof}
%eofclose

/*-- DECLARACIONES --*/
%{

    private Symbol SimboloACup(int Tipo, String Lexema) {
        Symbol Simbolo=new Symbol(Tipo, yyline,yycolumn,Lexema);
        return Simbolo;
    } 
	
	
%}

%init{ 
    
%init}

/*-- MACRO DECLARACIONES --*/
Espacio=[ \t]
Salto=\n
Bin=[0-1]

%%
/*-------- Secci?n de reglas lexicas ------ */
<YYINITIAL> {
	{Espacio} 	
		{
			//System.out.println("Vino espacio");
		}
	{Salto}		
		{
			//System.out.println("Vino Salto");
		}
        "."		
		{
                        int temp=Simbolos.Punto;
			//System.out.println("Vino "+Simbolos.terminalNames[temp]+" "+yytext());
                        return SimboloACup(temp,yytext());
		}
        {Bin}
                {
                        int temp=Simbolos.Bin;
			//System.out.println("Vino "+Simbolos.terminalNames[temp]+" "+yytext());
                        return SimboloACup(temp,yytext()); 
                }
	. 	{
			System.out.println("token ilegal < " + yytext()+ " > linea: " + (yyline+1) + "? columna:? " + (yycolumn+1));
		}
}


