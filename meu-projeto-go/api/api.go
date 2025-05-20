package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Definindo um tipo personalizado
type Estudo struct {
	Tema    string `json:"tema"`    //campo de Tema
	Horas   int    `json:"horas"`   //campo de horas
	Materia string `json:"materia"` //campo de matéria
}

// Função para tratar requisições feitas para a URL /estudos
// w http.ResponseWriter -> permite a resposta HTTP
// r *http.Request -> contém os dados requisição
func estudosHandler(w http.ResponseWriter, r *http.Request) {

	//Criando uma lista chamada Estudo
	estudos := []Estudo{
		{Tema: "Algoritmos de Escalonamento", Horas: 2, Materia: "Sistemas Operacionais"},
		{Tema: "Regex", Horas: 3, Materia: "Linguagem de Programação e Compiladores"},
	}

	//Diz para o navegador que o tipo de resposta será JSON
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	//Converte a lista de Estudo para JSON e envia como resposta da API
	json.NewEncoder(w).Encode(estudos)
}

func main() {
	//Associando a URL /estudos à função estudosHandler
	http.HandleFunc("/estudos", estudosHandler)

	//Mostrar uma mensagem no terminal
	fmt.Println("Servidor rodando em http://localhost:8080")

	//Inicia o servidor HTTP na porta 8080
	//nil -> roteador padrão da Go
	http.ListenAndServe(":8080", nil)
}
