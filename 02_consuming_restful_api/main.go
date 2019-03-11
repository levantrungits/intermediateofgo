package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"log"
)

// A response struct to map the Entire response
type Response struct {
	Name string 		`json:"name"`
	Pokemon []Pokemon 	`json:"pokemon_entries"`
}

// A Pokemon struct to map every pokemon to.
type Pokemon struct {
	EntryNo int				`json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
	Url string	`json:"url"`
}

func main()  {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for i:=0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
		fmt.Println(responseObject.Pokemon[i].Species.Url)
	}
}