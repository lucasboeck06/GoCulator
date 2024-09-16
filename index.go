package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func execute(input string) {
	parts := strings.Split(input, " ")

	if len(parts) != 2 {
		fmt.Println("Erro, digite apenas dois números válidos!")
		return
	}

	a, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("Problema ao converter A:", err)
		return
	}

	b, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Problema ao converter B:", err)
		return
	}

	if a%b == 0 || b%a == 0 {
		fmt.Println("São Múltiplos")
	} else {
		fmt.Println("Não são Múltiplos")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Digite dois números inteiros, ou 'sair':")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Erro detectado!")
			return
		}

		input = strings.TrimSpace(input)
		if input == "sair" {
			fmt.Println("Encerrando programa...")
			break
		}

		execute(input)
	}
}
