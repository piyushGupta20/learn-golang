package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Difficulty string `json:"difficulty"`
}

func main() {
	url := "https://dummyjson.com/recipes"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request: ", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading the response body:", err)
		return
	}

	var data Data
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshelling the response", err)
		return
	}

	fmt.Print(string(body))

}
