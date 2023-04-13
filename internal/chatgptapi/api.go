package chatgptapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type API struct {
	key string
}

func New(apiKey string) *API {
	return &API{
		key: apiKey,
	}
}

func (a *API) GetResponse(prompt string) string {
	gptURL := "https://api.openai.com/v1/engines/davinci-codex/completions"
	client := &http.Client{}

	requestBody, _ := json.Marshal(gptRequest{
		Prompt:    prompt,
		MaxTokens: 10000,
	})

	req, _ := http.NewRequest("POST", gptURL, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.key))

	resp, err := client.Do(req)
	if err != nil {
		return "Error fetching response from GPT API"
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var gptResponse gptResponse
	json.Unmarshal(bodyBytes, &gptResponse)

	if len(gptResponse.Choices) == 0 {
		return "No response from GPT API"
	}

	return gptResponse.Choices[0].Text
}

type gptRequest struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type gptResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}
