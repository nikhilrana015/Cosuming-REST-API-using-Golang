package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type JokeResponse struct {
	Success bool `json:"success"`
	Body    []struct {
		ID        string `json:"_id"`
		Setup     string `json:"setup"`
		Punchline string `json:"punchline"`
		Type      string `json:"type"`
		Likes     []any  `json:"likes"`
		Author    struct {
			Name string `json:"name"`
			ID   any    `json:"id"`
		} `json:"author"`
		Approved      bool   `json:"approved"`
		Date          int    `json:"date"`
		Nsfw          bool   `json:"NSFW"`
		ShareableLink string `json:"shareableLink"`
	} `json:"body"`
}

func getKeys() (string, string) {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	apiKey, _ := viper.Get("RapidAPI_Key").(string)
	apiKeyHost, _ := viper.Get("RapidAPI_Host").(string)

	//fmt.Println(apiKey, apiKeyHost)

	return apiKey, apiKeyHost
}

func main() {

	client := &http.Client{}

	url := "https://dad-jokes.p.rapidapi.com/random/joke"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	apiKey, apiKeyHost := getKeys()

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", apiKeyHost)

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var joke JokeResponse

	//fmt.Println(response.Body)

	bytesData, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(bytesData))

	err = json.Unmarshal(bytesData, &joke)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Api response as: %v", joke.Body)
	fmt.Printf("Setup: %v\nPunchline: %v", joke.Body[0].Setup, joke.Body[0].Punchline)

}
