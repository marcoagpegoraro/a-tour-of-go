package main

import "fmt"

var globalVarAvaliableInPackage = "Olá"

func main() {
	//A fmt.Println, alem de printar na tela, ela retorna o
	// numero de bytes e também se houve algum erro
	numeroDeBytes, erros := fmt.Println("Hello world", "caraioooo")

	//toda variavel no go tem que ser utilizada se não da erro,
	// caso não queria utilziar, coloque ela como _
	fmt.Println(numeroDeBytes, erros)

	x := 16
	y := "texto"
	z := true

	// fmt.Println(x, y, z)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)

	x = 20
	fmt.Printf("x: %v, %T\n", x, x)

	x, a := 30, false //Operador de atribuição pode ser usado se pelo menos uma das variaveis é nova
	fmt.Printf("a: %v, %T\n", a, a)

}
