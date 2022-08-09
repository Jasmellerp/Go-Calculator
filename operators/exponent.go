package operators

import (
	"fmt"
	"strconv"
	"strings"
)

func Power(operation string) int {
	values := strings.Split(operation, "^")
	result := 0

	for i := 0; i < len(values); i++ {
		num, error := strconv.Atoi(values[i])

		if error != nil {
			fmt.Println("Error: Tiene que ingresar un numero entero")
			fmt.Println("Solo debes realizar con un solo operador")
		} else {
			result ^= num
		}

	}

	return result

}
