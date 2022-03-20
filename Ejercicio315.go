package main


import (

	"fmt"
)

func NumerosDuplicado(Arreglo []int){

	var aux int 
	var entradas_modificadas int

	for i := 0 ; i < len(Arreglo) ; i++ {

		for j := i + 1 ; j<len(Arreglo); j++{
			aux = Arreglo[i]

			if(aux == -5 ){

			}else{

				if(aux == Arreglo[j]){
					Arreglo[j] = -5 
					entradas_modificadas++
				}

				



			}

			

		}
		
	}

	
	fmt.Print("[")

	for i := 0 ; i < len(Arreglo) ; i++ {

		fmt.Print(Arreglo[i], " ")
		
	}

	fmt.Println("]")


	fmt.Print("Numero de entradas modificadas :", entradas_modificadas)
	



}

func main(){

	var x,y int
	fmt.Print(" Digite el tamaño del arreglo : ")
	fmt.Scan(&x)

	Arreglo := make([] int, x)

	for i := 0 ; i < len(Arreglo) ; i++ {

		fmt.Print(" Digite un numero : ")
	    fmt.Scan(&y)
		Arreglo[i] = y 

	}
	fmt.Println("Arreglo Original")

	fmt.Print("[")

	for i := 0 ; i < len(Arreglo) ; i++ {

		fmt.Print(Arreglo[i], " ")
		
	}

	fmt.Println("]")


	fmt.Println("Arreglo Pedido")


	NumerosDuplicado(Arreglo)

	





}