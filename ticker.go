package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmpSite := "https://financialmodelingprep.com/api/v3"
	apiKey := "?apikey=3633c381ffbe8e07e7fe8d8e3196f386 "
	response, err := http.Get(fmpSite + "/quote-short/GME" + apiKey)
	fmt.Println(fmpSite + "/quote-short/GME" + apiKey)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	type Message struct {
		Symbol string
		Price  float64
	}

	var m []Message
	responseBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseBody))
	json.Unmarshal(responseBody, &m)
	fmt.Printf("%s: %f\n", m[0].Symbol, m[0].Price)
}
