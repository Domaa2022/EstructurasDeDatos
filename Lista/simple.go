package lista

import "fmt"

type Nodo struct {
	Siguiente *Nodo
	Dato int 
}

type Lista struct{
	inicio *Nodo
	Ultimo *Nodo
	Tamaño int 
}

func NewLista() (*Lista){
	return &Lista{nil,nil,0}

}

func (m *Lista) InsertarAlFinal(dato int){
	nuevo := &Nodo{nil,dato}

	if(m.inicio == nil){
		m.inicio = nuevo
		m.Ultimo = nuevo
	}else{
		m.Ultimo.Siguiente = nuevo
		m.Ultimo= nuevo
	}
	m.Tamaño ++

}

func(m * Lista) EliminarPosicion( Posicion int){



	if(m.Tamaño >= Posicion){
		aux := m.inicio
	 if(Posicion == 1){
		m.inicio = m.inicio.Siguiente

		}else {
		
		

		for i := 2 ; i <= Posicion ; i++ {
			aux = aux.Siguiente
		}
	
		aux.Siguiente = m.Ultimo.Siguiente
		m.Ultimo = aux 





	   }
	
	}else{
		fmt.Println("La posicion no exite en la lista")
	}
	
}




func (m *Lista) Imprimir() {
	x := m.inicio
	for x != nil{
		fmt.Print("[", x.Dato ,"]->")
		x = x.Siguiente

	}
	fmt.Println()
	
}