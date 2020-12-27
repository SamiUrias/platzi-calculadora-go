package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {}

func (calc) operate (entrada string, operador string) int {

	valores := strings.Split(entrada, operador)
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])

	// Cast valores from text to number
	operador1 := parsear(valores[0])
	operador2 := parsear(valores[1])


	respuesta := 0
	switch operador {
	// Suma de valores
	case "+":
		respuesta = operador1 + operador2
	case "-":
		respuesta = operador1 - operador2
	case "*":
		respuesta = operador1 * operador2
	case "/":
		respuesta = operador1 / operador2
	default:
		fmt.Println("El operador no esta soportado")
	}
	fmt.Println("Respuesta: ", respuesta)
	return respuesta
}


func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Print("Ingrese la entrada: ")
	entrada := leerEntrada()

	fmt.Println("Ingrese el operador:")
	operador := leerEntrada()

	fmt.Println("Entrada: ", entrada)
	fmt.Println("Operador: ", operador)

	c := calc{}
	fmt.Printf(string(c.operate(entrada, operador)))
}