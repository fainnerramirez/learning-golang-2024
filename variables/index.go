package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Variables en Go");

	//strings
	var filename string = "nombre_archivo.go"; //tipado fuerte
	var filename2 = "otro_nombre_archivo.go"; // infiere el tipo
	fmt.Println("string: ", filename, filename2);

	//ints
	var bytes int = 100; // tipado fuerte
	var bytes2 = 200; // infiere el tipo
	
	fmt.Println("Bytes file: ", bytes, bytes2);
	
	//bool
	var sw bool = true;

	fmt.Println("bool: ", sw);
	fmt.Println(reflect.TypeOf(sw));

	// obtener tipos de datos
	var dataType = reflect.TypeOf(filename);
	var dataType2 = reflect.TypeOf(bytes);

	fmt.Println("dataType: ", dataType, dataType2)

	//==================================================

	//forma mas corta de declarar variables (:=)
	state := true;
	fmt.Println("el estado es: ", state);

	//constantes
	const API_KEY = "NKNSDJSDJKD7E2K3KJ23K2";
	fmt.Println("API KEY: ", API_KEY);
}