package main

import ( 
	
	"fmt"
)

func unir_arreglos (arr1 [ ]int , arr2 [ ]int) {


   arr3 := make([]int, len(arr1)+len(arr2))
   var contador int



   //introducir arr 1 en arr 3 

   for i := 0 ; i < len(arr1) ; i ++ {
	  arr3[i] = arr1[i]
   }

   //introducir arre 2 en arr 3 

   for i := 0 ; i < len(arr2) ; i++ {
	   arr3[len(arr1)+i] = arr2[i]
   }
   
   // ordenar arreglo 

   for i := 0 ; i< len(arr3) ; i++{

	for j := i + 1 ; j < len(arr3) ; j ++ {

	   if( arr3[i] > arr3[j]){
		   aux := arr3[i]
		   arr3[i] = arr3[j]
		   arr3[j] = aux
	   }
	}
  }
   // eliminar numeros 

   for i := 0 ; i < len(arr3) ; i++ {
	for j := i + 1 ; j<len(arr3); j++{
		aux := arr3[i]
		if(aux == -5){
			
		}else{

			if(aux == arr3[j]){
				arr3[j] = -5
				contador++
			
				
			}



		}
	}
  }
  for i := 0 ; i< len(arr3) ; i++{

	for j := i + 1 ; j < len(arr3) ; j ++ {

	   if( arr3[i] > arr3[j]){
		   aux := arr3[i]
		   arr3[i] = arr3[j]
		   arr3[j] = aux
	   }
	}
  }

  arre5 := arr3[contador:len(arr3)]


 

  //imprimir 

  

	fmt.Print(arre5)


	
   
	
}



func main(){

	arr1 := []int{1, 1, 2, 1, 9, 11, 12}
	arr2 := []int{1, 2, 3, 4, 5, 10, 12}

	unir_arreglos(arr1,arr2)





}