package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GptClient struct {
	apiKey string
	messages []map[string]string
}

// CreateClient returns a new GptClient instance with a given API key.
//
// The client is initialized with a default set of messages that tell the
// AI to respond with short answers.
func CreateClient(apiKey string) *GptClient {
	// Create a new client instance
	client := &GptClient{
		apiKey: apiKey,
		messages: []map[string]string{
		{
			"role":    "user",
			"content": "Give as short answers as possible",
		},
		{
			"role":    "system",
			"content": "Okay!",
		},
		},
	}
	return client
}

// GenerateResponse sends a request to the OpenAI API to generate a response to a given prompt, using the given model.
//
// The method takes a GptClient instance, a prompt to generate a response for, and a model to use for response generation.
// The method returns a string containing the generated response.
//
// The method sends a POST request to the OpenAI API with the given prompt and model.
// The method uses the client's messages slice to store the history of the conversation.
// The method uses the client's apiKey property to set the Authorization header of the request.
func GenerateResponse(client *GptClient, prompt string, model string) (string, error) {
	messages := append(client.messages, map[string]string{
		"role":    "user",
		"content": prompt,
	})

	jsonData, err := json.Marshal(map[string]interface{}{
		"model":    model,
		"messages": messages,
	})
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	req, err := http.NewRequest(
		"POST", 
		"https://api.proxyapi.ru/openai/v1/chat/completions", 
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.apiKey)

	client1 := &http.Client{}
	resp, err := client1.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	response := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	fmt.Println(response)
	client.messages = append(client.messages, map[string]string{
		"role":    "assistant",
		"content": response,
	})
	return response, nil
}


type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

// GetBalance gets the current balance of the given API key.
//
// The request is sent to the ProxyAPI /balance endpoint. The response
// is expected to be a JSON object containing the "balance" key.
//
// If the request is successful, the function returns a pointer to a
// BalanceResponse, which contains the balance of the given API key.
//
// If the request fails, the function returns an error.
func GetBalance(apiKey string) (*BalanceResponse, error) {
	url := "https://api.proxyapi.ru/proxyapi/balance"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var balanceResponse BalanceResponse
	err = json.Unmarshal(body, &balanceResponse)
	if err != nil {
		return nil, err
	}

	return &balanceResponse, nil
}