package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func areaCirc(reader *bufio.Reader) {
	fmt.Println("Digite o Raio do Círculo:")

	raioCirc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler Raio:", err)
		return
	}

	raioCircFloat, _ := strconv.ParseFloat(strings.TrimSpace(raioCirc), 64)

	result := math.Pi * (raioCircFloat * raioCircFloat)

	fmt.Println("A Área do seu Círculo é:", result)
	reader.ReadString('\n')
}

func areaQuad(reader *bufio.Reader) {
	fmt.Println("Digite o Lado do Quadrado:")

	ladoQuad, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler Lado:", err)
		return
	}

	ladoQuadFloat, _ := strconv.ParseFloat(strings.TrimSpace(ladoQuad), 64)

	result := ladoQuadFloat * ladoQuadFloat

	fmt.Println("A Área do seu Quadrado é:", result)
	reader.ReadString('\n')
}

func volumeCub(reader *bufio.Reader) {
	fmt.Println("Digite a Aresta do Cubo:")

	arestaCub, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler Aresta:", err)
		return
	}

	arestaCubFloat, _ := strconv.ParseFloat(strings.TrimSpace(arestaCub), 64)
	result := arestaCubFloat * arestaCubFloat * arestaCubFloat

	fmt.Println("O Volume do seu Cubo é:", result)
	reader.ReadString('\n')
}

func volumeEsf(reader *bufio.Reader) {
	fmt.Println("Digite o Raio da Esfera:")

	raioEsf, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler Raio:", err)
		return
	}

	raioEsfFloat, _ := strconv.ParseFloat(strings.TrimSpace(raioEsf), 64)
	result := (4.0 / 3.0) * math.Pi * (raioEsfFloat * raioEsfFloat * raioEsfFloat)

	fmt.Println("O Volume da sua Esfera é:", result)
	reader.ReadString('\n')

}

func volumeCon(reader *bufio.Reader) {
	fmt.Println("Digite o Raio e Altura do Cone:")

	valoresCon, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler valores:", err)
		return
	}

	parts := strings.Split(valoresCon, " ")

	if len(parts) != 2 {
		fmt.Println("Erro, digite apenas dois números válidos!")
		return
	}

	raio, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Erro ao converter raio para float:", err)
		return
	}
	altura, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		fmt.Println("Erro ao converter altura para float:", err)
	}

	result := (1.0 / 3.0) * math.Pi * (raio * raio) * altura

	fmt.Println("O Volume do seu Cone é:", result)
	reader.ReadString('\n')
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()

		fmt.Println("========== Bem-vindo ao GoCulator! =============\n" +
			"Digite o número referente a equação que deseja realizar:\n" +
			"1 - Área\n" +
			"2 - Volume\n" +
			"3 - Perímetro\n" +
			"4 - Sair")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler Input:", err)
			return
		}

		switch strings.TrimSpace(input) {
		case "1":
			clearScreen()

			fmt.Println("=========== Cálculo de Área ===========\n" +
				"Digite o tipo da operação de área:\n" +
				"1 - Círculo\n" +
				"2 - Quadrado")

			areaInput, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Erro ao ler forma:", err)
				return
			}

			switch strings.TrimSpace(areaInput) {
			case "1":
				clearScreen()
				areaCirc(reader)
			case "2":
				clearScreen()
				areaQuad(reader)
			default:
				clearScreen()
				fmt.Println("Digite um valor válido!")
			}

		case "2":
			clearScreen()

			fmt.Println("========= Cálculo de Volume ========\n" +
				"Digite o tipo de operação de volume:\n" +
				"1 - Cubo\n" +
				"2 - Esfera\n" +
				"3 - Cone\n" +
				"4 - Cilindro")

			volumeInput, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Erro ao ler forma:", err)
				return
			}

			switch strings.TrimSpace(volumeInput) {
			case "1":
				clearScreen()
				volumeCub(reader)
			case "2":
				clearScreen()
				volumeEsf(reader)
			case "3":
				clearScreen()
				volumeCon(reader)
			// case "4":
			// 	clearScreen()
			// 	volumeCil(reader)
			default:
				clearScreen()
				fmt.Println("Digite um valor válido!")
			}

		case "4":
			clearScreen()
			fmt.Println("Saindo do programa...")
			return

		default:
			fmt.Println("Digite um valor válido!")
		}
	}
}
