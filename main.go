package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jasmellerp/calculadora/operators"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Calculator")
	scanner.Scan()
	operation := scanner.Text()

	result := 0
	if strings.Contains(operation, "+") {
		result = operators.Sum(operation)
	} else if strings.Contains(operation, "-") {
		result = operators.Subtraction(operation)
	} else if strings.Contains(operation, "*") {
		result = operators.Multiplication(operation)
	} else if strings.Contains(operation, "/") {
		result = operators.Division(operation)
	} else if strings.Contains(operation, "^") {
		result = operators.Power(operation)
	} else {
		fmt.Println("Error: El operador esta mal ingresado")
		fmt.Println("Este operador no esta implementado")
	}

	fmt.Println(result)
}
