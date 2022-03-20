package main

import(
	"fmt" 
	"math/rand"
	"time"
)

func OrdenarArreglo( Arreglo [20]int ){
	 
var Aux int 

	for i := 0 ; i< len(Arreglo); i++ {
		for j := i+1 ; j < len(Arreglo); j++{


			if (Arreglo[i] < Arreglo[j]){
				Aux = Arreglo[i]
				Arreglo[i] = Arreglo[j]
				Arreglo[j] = Aux
			}
		}
	}

	for i := 0 ; i < len(Arreglo) ; i++ {

		fmt.Print(Arreglo[i], " ")
		
	}

}





func main(){

	semilla := time.Now().UnixMilli()
	rand.Seed(semilla)

	var arreglo [20]int

	for i := 0 ; i < len(arreglo) ; i++ {

		arreglo[i] = rand.Intn(100)

	}

	fmt.Println("Arreglo Original")

	fmt.Print("[")

	for i := 0 ; i < len(arreglo) ; i++ {

		fmt.Print(arreglo[i], " ")
		
	}

	fmt.Println("]")


	fmt.Println("Arreglo Pedido")

	fmt.Print("[")

	OrdenarArreglo(arreglo)

	fmt.Print("]")



// Big O = 1 porque depende de el tamaño del arayy




}