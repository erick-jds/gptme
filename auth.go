package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthRequest struct {
	APIKey string `json:"apiKey"`
}

type ChatInput struct {
	Input string `json:"input"`
}

type ChatResponse struct {
	Output string `json:"output"`
}

func authenticate(apiKey string) error {
	url := "https://api.chatgpt.com/v1/authenticate"
	authReq := AuthRequest{APIKey: apiKey}
	authReqBytes, err := json.Marshal(authReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(authReqBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication failed with status code %d", resp.StatusCode)
	}

	return nil
}

func main() {
	key := "sk-aaaaaaaaaaaaaaaaaaaaaaaaa"
	err := authenticate(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Autenticação bem-sucedida")
	}

	// Define a mensagem de prompt a ser enviada para o ChatGPT
	promptMessage := "Qual é a minha sorte hoje?"

	// Define a URL da API do ChatGPT
	url := "https://api.chatgpt.com/v1/chat"

	// Cria a estrutura de dados para a mensagem de entrada
	inputData := &ChatInput{
		Input: promptMessage,
	}

	// Converte a estrutura de dados em JSON
	inputJson, _ := json.Marshal(inputData)

	// Cria uma nova solicitação HTTP POST para a API do ChatGPT
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(inputJson))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Envia a solicitação HTTP POST para a API do ChatGPT e trata qualquer erro
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar solicitação para o ChatGPT:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro do bodyBytes: ", err)
		return
	}
	fmt.Println(string(bodyBytes))

	// Lê a resposta da API do ChatGPT e exibe a mensagem de saída
	var chatResponse ChatResponse
	fmt.Println(chatResponse)
	err = json.NewDecoder(resp.Body).Decode(&chatResponse)
	if err != nil {
		fmt.Println("Erro ao ler resposta da API do ChatGPT:", err)
		return
	}
	fmt.Println("Resposta do ChatGPT:", chatResponse.Output)

}
