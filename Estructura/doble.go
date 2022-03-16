package estructura

import (
	"fmt"
)

//Lugar donde almacenaremos la informacion 
type Nodo struct{
	Anterior *Nodo
	Siguiente *Nodo
	Dato int 
}

// Estructura para almacenar nodos de informacion

type Lista struct{
	Inicio *Nodo
	Ultimo *Nodo
	Tamaño int 
}

//Crear una nueva lista 
func NewLista() *Lista {
	return &Lista{nil,nil,0}
}

//Insertar un nodo
//Func - Recibir  - Nombre(Mandar)
func (n *Lista) Insertar(valor int){
	nuevo := &Nodo{nil,nil,valor}
	
	if n.Inicio == nil{
		n.Inicio= nuevo
		n.Ultimo= nuevo
	}else{
		//Es el enlace siguiente que apunta al nuevo modo 
		n.Ultimo.Siguiente = nuevo
		//es el enlace anterior del nuevo nodo que apunta al ultimo nodo existente 
		nuevo.Anterior = n.Ultimo
		//Se le asigna al nuevo como el ultimo nodo de la lista 
		n.Ultimo = nuevo 
	}
	n.Tamaño ++ 
}

func (n *Lista) InsertarInicio(valor int){
	primero := &Nodo{nil,nil,valor}
	if(n.Inicio==nil){
		n.Inicio = primero
		n.Ultimo = primero
	}else{
		n.Inicio.Anterior= primero
		primero.Siguiente= n.Inicio
		n.Inicio = primero

	}
	n.Tamaño ++
}

func (n *Lista) InsertarEntre( valor int, Posicion int){
	aux := n.Inicio

	for  i := 2 ; i < Posicion ; i++ {
		aux=aux.Siguiente
	}
	//Aux = 5  Y Aux2= 9 
	aux2 := aux.Siguiente

	nuevo := &Nodo{nil,nil,valor}

	aux.Siguiente= nuevo
	nuevo.Anterior= aux
	nuevo.Siguiente= aux2
	aux2.Anterior= nuevo
	n.Tamaño++

}


//Imprimi la lista
func (n *Lista) Imprimir(){
	x := n.Inicio
	for x != nil {
		fmt.Print("<-[",x.Dato,"]->")
		x = x.Siguiente
	}
	fmt.Println()
	fmt.Print("Tamaño lista: " , n.Tamaño)




}


