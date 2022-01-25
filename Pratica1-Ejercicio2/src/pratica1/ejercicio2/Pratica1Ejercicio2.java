/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package pratica1.ejercicio2;

import Analizadores.Lexico;
import Analizadores.Sintactico;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.io.StringReader;
import javax.swing.JFileChooser;
import javax.swing.JOptionPane;
import javax.swing.filechooser.FileFilter;
import javax.swing.filechooser.FileNameExtensionFilter;

/**
 *
 * @author alterlex
 */
public class Pratica1Ejercicio2 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        String Cadena="";
        JFileChooser Abrir =new JFileChooser(".");
        FileFilter Filtro= new FileNameExtensionFilter("File .txt", "txt");
        Abrir.setFileFilter(Filtro);
        int temp= Abrir.showOpenDialog(Abrir);
        if (temp == JFileChooser.APPROVE_OPTION){
            File Archivo = Abrir.getSelectedFile();
            String Bufer="";
          try {
                FileReader Lector;
                Lector = new FileReader(Archivo);
                BufferedReader LectorTexto = new BufferedReader(Lector);
                String LineaArchivo;
                while((LineaArchivo= LectorTexto.readLine())!=null){
                    Bufer=Bufer+LineaArchivo+"\n";
                }
                Lector.close();
                LectorTexto.close();
                Cadena=Bufer;
                
            } 
             catch (IOException ex) {
                 JOptionPane.showMessageDialog(null, "No se pudo abrir el archivo "+ex, "Abrir Archivo", JOptionPane.INFORMATION_MESSAGE);

            } 
            
        }
        else if (temp== JFileChooser.CANCEL_OPTION){
            JOptionPane.showMessageDialog(null, "No Selecciono ningun Archivo", "Abrir Archivo", JOptionPane.INFORMATION_MESSAGE);
        }
        
        Cadena=Cadena.trim();
        System.out.println("Entrada: \n"+Cadena);
        System.out.println("Salida:");
        StringReader Lector=new StringReader(Cadena);
        Lexico lex=new Lexico(Lector);
        Sintactico sin=new Sintactico(lex);
        try {
            sin.parse();
        } catch (Exception ex) {
            
        }
    }
    
}
