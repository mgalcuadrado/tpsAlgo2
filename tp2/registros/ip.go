package registros

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	_DELIMITADOR_IPv4      string = "."
	_BASE_NUMERICA         int    = 10
	_CANTIDAD_DE_BITS_IPv4 int    = 8
)

type IPv4 struct {
	partes [4]uint8
}

// IPParsear recibe una cadena y la devuelve como una dirección ipv4
// como precondicion la cadena debe ser una dirección ipv4
func IPParsear(cadena string) IPv4 { // O(4) = O(1)
	direccion := new(IPv4)
	partes := strings.Split(cadena, _DELIMITADOR_IPv4)
	for indice, _ := range partes {
		num, err := strconv.ParseUint(partes[indice], _BASE_NUMERICA, _CANTIDAD_DE_BITS_IPv4)
		if err != nil {
			panic(_MENSAJE_ERROR)
		}
		direccion.partes[indice] = uint8(num)
	}
	return *direccion
}

// IPCompare recibe dos direcciones ip en formato ipv4 y devuelve
// -1 si ip1 < ip2,
// 0 si ip1 = ip2,
// 1 si ip1 > ip2
func IPCompare(ip1 IPv4, ip2 IPv4) int {
	for i, _ := range ip1.partes {
		if ip1.partes[i] < ip2.partes[i] {
			return -1
		} else if ip1.partes[i] > ip2.partes[i] {
			return 1
		}
	}
	return 0
}

func IPCompareInverso(ip1 IPv4, ip2 IPv4) int {
	for i, _ := range ip1.partes {
		if ip1.partes[i] < ip2.partes[i] {
			return 1
		} else if ip1.partes[i] > ip2.partes[i] {
			return -1
		}
	}
	return 0
}

func IPDisplay(ip IPv4) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.partes[0], ip.partes[1], ip.partes[2], ip.partes[3])
}
