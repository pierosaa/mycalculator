package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
		return operador1 - operador2,nil
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2,nil
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2,nil
	default:
		fmt.Println(operador, "No está soportado")
		return 0,nil
	}

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
