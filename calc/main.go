package main // Nome do arquivo

import ("fmt") //Importando o formatador

func sum(x, y int) int{ //Função para somar
	return x + y
}
func sub(x, y int) int{ //Função para subtrair
	return x - y
}
func mult(x, y int) int{ //Função para multiplicar
	return x * y
}
func div(x, y int) (int) { //Função para Dividir
	return x / y
}

func main() { //Função de entrada
	var tipoCalc string //Váriaveis 
	var input int
	var input2 int

    fmt.Print("Digite o método do calculo: (OBS: Somar, Subtrair, Multiplicar , Dividir)") //Mensagem de escolha
	fmt.Scanln(&tipoCalc) // Campo para digitar
	switch tipoCalc { //Condição para calcular de acordo como escolhido
	case "Somar":
		fmt.Print("Primeiro valor: ")
		fmt.Scanln(&input)
		fmt.Print("Segundo Valor: ")
		fmt.Scanln(&input2)
		fmt.Println("Resultado da Soma:", sum(input, input2)) //Mostrando o resultado
	case "Subtrair":
		fmt.Print("Primeiro valor: ")
		fmt.Scanln(&input)
		fmt.Print("Segundo Valor: ")
		fmt.Scanln(&input2)
		fmt.Println("Resultado da Subtração:", sub(input, input2))
	case "Multiplicar":
		fmt.Print("Primeiro valor: ")
		fmt.Scanln(&input)
		fmt.Print("Segundo Valor: ")
		fmt.Scanln(&input2)
		fmt.Println("Resultado da Multiplicação:", mult(input, input2))
	case "Dividir":
		fmt.Print("Primeiro valor: ")
		fmt.Scanln(&input)
		fmt.Print("Segundo Valor: ")
		fmt.Scanln(&input2)
		
		if input <= 0 || input2 <= 0 { // Condição para impedir divisão por zero 
			fmt.Println("Não e possível divisão por zero")
		}else {
			fmt.Println("Resultado da Divisão:", div(input, input2, ))
		}

	}
    
    
}

