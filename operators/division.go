package operators

import (
	"fmt"
	"strconv"
	"strings"
)

func Division(operation string) int {
	values := strings.Split(operation, "/")
	result := 0

	for i := 0; i < len(values); i++ {
		num, error := strconv.Atoi(values[i])

		if error != nil {
			fmt.Println("Error: Tiene que ingresar un numero entero")
			fmt.Println("Solo debes realizar con un solo operador")
		} else {
			if result == 0 {
				result = num
			} else {
				result /= num
			}

		}

	}

	return result

}
