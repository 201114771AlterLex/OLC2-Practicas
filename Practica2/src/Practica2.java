
import Analizadores.Lexico;
import Analizadores.Sintactico;
import java.io.StringReader;
import javax.swing.JOptionPane;

/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

/**
 *
 * @author alterlex
 */
public class Practica2 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        String Cadena="";
        Cadena=JOptionPane.showInputDialog("Ingrese Cadena Binaria con o sin punto decipal");
        
        Cadena=Cadena.trim();
        System.out.println("Entrada: \n"+Cadena);
        System.out.println("Salida:");
        StringReader Lector=new StringReader(Cadena);
        Lexico lex=new Lexico(Lector);
        Sintactico sin=new Sintactico(lex);
        try {
            sin.parse();
        } catch (Exception ex) {
            System.out.println("Error "+ex.toString());
        }
    }
    
}
