package main

import "fmt"

// O ponto de ebulição da água em Kelvin é 373.15K
const pontoEbulicaoKelvin float64 = 373.15

func main() {

	// Fórmula de conversão: C = K - 273.15
	pontoEbulicaoCelsius := pontoEbulicaoKelvin - 273.15

	// Exibe o resultado formatado com duas casas decimais
	fmt.Printf("O ponto de ebulição da água é:\n")
	fmt.Printf("%.2f K (Kelvin)\n", pontoEbulicaoKelvin)
	fmt.Printf("%.2f °C (Celsius)\n", pontoEbulicaoCelsius)
}
