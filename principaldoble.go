package main

import(
	"fmt"
	doble "ESCTRUCTURAS/Estructura"
)

func main(){

	lista := doble.NewLista()
	lista.Insertar(5)
	lista.InsertarInicio(10)
	lista.InsertarEntre(9,2)
	lista.Imprimir()
	fmt.Print()




}