package main //Nome do arquivo
import ( //Importando as Bibliotecas
	"fmt" 
	"io"
	"net/http"
	"encoding/json"
)

type Usuario struct { //Estrutura do JSON
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Causa string `json:"causa"`
	Data  string `json:"data"`
}

func main() { // Função de entrada

	url := "https://raw.githubusercontent.com/PedroHenri228/Python/main/JSON/output/dados.json" //url para fazer a requisição

	response, err := http.Get(url) //Fazendo um requisição GET
	if err != nil { //Condição para evitar erros
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer response.Body.Close() //Garantindo que a requisição será fechada no final da função

	if response.StatusCode != http.StatusOK { //Debug caso aconteça erros na requisição
		fmt.Printf("Erro: status da resposta %d %s\n", response.StatusCode, response.Status)
		return
	}

	body, err := io.ReadAll(response.Body) //Lendo tudo do arquivo recebido pela requisição
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return
	}

	var usuarios []Usuario //Array para armazenar os dados JSON recebidos

	err = json.Unmarshal(body, &usuarios) //Decodifica o JSON
	if err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	for _, usuario := range usuarios { //Loop para imprimir os dados JSON recebidos
		fmt.Println("Nome do usuário:", usuario.Nome, "\nIdade:", usuario.Idade, "\nCausa:", usuario.Causa, "\nData:", usuario.Data)
	}
}
