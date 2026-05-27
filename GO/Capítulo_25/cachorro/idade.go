// Package criado como exercícion 1 do capítulo 26
package cachorro

// Constante de equivalência de ano humano para anos caninos
const UM_ANO_CANINO int = 7

// Função responsável por calcular a idade humana em idade canina.
func Idade(idade int) int {
	anos_caninos := idade * UM_ANO_CANINO
	return anos_caninos
}
