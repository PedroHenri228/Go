package main // Nome do Arquivo

import (
	"fmt" //Importardo o formatador
	"net/http" // Importando a biblioteca requisições http entre outros, 
)

func main() { // Função de entrada

	fs := http.FileServer(http.Dir("static")) // Definindo o diretório na pasta 'static'

	http.Handle("/", fs) //Mapeando a raiz do servidor para o FIleServer

	fmt.Println("Servidor rodando na porta 8080...") // Mensagem informando a porta

	err := http.ListenAndServe(":8080", nil) //Porta principal 8080
	if err != nil { // Condição para verificação de erros
		fmt.Println("Erro ao iniciar o servidor:", err)
	}



}