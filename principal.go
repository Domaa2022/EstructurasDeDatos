package main

import (
	simple "ESCTRUCTURAS/Lista"
	"fmt"
	"math/rand"
	"time"
)

func main(){

	semilla := time.Now().UnixMilli()
	rand.Seed(semilla)
	valor  := rand.Intn(10)
	lis := simple.NewLista()

	var x int 

	for i := 0 ; i <= valor ; i++ {

		lis.InsertarAlFinal(rand.Intn(50))

	}

	fmt.Print()
	lis.Imprimir()

	fmt.Print("En que posicion quiere que termine la lista: ")
	fmt.Scan(&x)


	lis.EliminarPosicion(x)
	lis.Imprimir()

	
}