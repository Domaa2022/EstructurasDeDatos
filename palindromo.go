package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Espalindromo( b int ) bool {

	boleano := true
	c := strconv.Itoa(b)
	x := strings.Split(c, "")
	final := len(x) - 1
	v := make([]string , len(x))


	for i := 0 ; i < len(x) ; i ++ {

		v[i] = x[final-i];
	} 

	for i := 0 ; i < len(x) ; i++{

		if( v[i] == x[i]){

		}else{
			boleano = false
		}
	}

	return boleano

}






func main(){


	var Entero int 

	fmt.Print("Digite un numero")
	fmt.Scan(&Entero)


	n := Espalindromo(Entero)

	fmt.Print(" ¿ El numero : ", Entero , " Es un palindro? : ", n )
		


	

	

}