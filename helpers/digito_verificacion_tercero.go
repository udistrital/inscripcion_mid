package helpers

import (
	// "fmt"
	"regexp"
	"strconv"
)

func CalcularDigitoVerificacion(nit string) int {
	re := regexp.MustCompile(`[^0-9]`)
	nit = re.ReplaceAllString(nit, "")

	factores := []int{3, 7, 13, 17, 19, 23, 29, 37, 41, 43, 47, 53, 59, 67, 71}
	if len(nit) > len(factores) {
		nit = nit[len(nit)-len(factores):]
	}

	suma := 0
	for i := 0; i < len(nit); i++ {
		posicion := len(nit) - 1 - i
		digito, _ := strconv.Atoi(string(nit[posicion]))
		suma += digito * factores[i]
	}

	modulo := suma % 11

	if modulo == 0 || modulo == 1 {
		return (modulo)
	}

	return (11 - modulo)
}
