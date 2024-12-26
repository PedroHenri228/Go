package main // Nome do Arquivo

import (
	"encoding/json"
	"fmt"      //Importardo o formatador
	"net/http" // Importando a biblioteca requisições http entre outros,
)

func GetData(w http.ResponseWriter, r *http.Request)  {

	 data := "Olá Mundo!!!!"

	 w.Header().Set("Content-Type", "application/json")

	 err := json.NewEncoder(w).Encode(data)

	 if err != nil {
		http.Error(w, "Falha ao codificar resposta", http.StatusInternalServerError)
		
	 }
}
func main() {

 
	fs := http.FileServer(http.Dir("static")) // Definindo o diretório na pasta 'static'

	http.Handle("/", fs) //Mapeando a raiz do servidor para o FIleServer
	http.HandleFunc("/data", GetData)
	fmt.Println("Servidor rodando na porta 8080...") // Mensagem informando a porta

	err := http.ListenAndServe(":8080", nil) //Porta principal 8080
	if err != nil { // Condição para verificação de erros
		fmt.Println("Erro ao iniciar o servidor:", err)
	}



}