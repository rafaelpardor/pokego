package pokemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type PokemonJson struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
}

func (r *PokemonJson) ToPokemonModel() *Pokemon {
	return &Pokemon{
		Id:     r.Id,
		Name:   r.Name,
		Height: r.Height,
	}
}

func (p *PokemonJson) printReturn() string {
	return fmt.Sprintf("id: %d, name: %s, height: %d", p.Id, p.Name, p.Height)
}

var urlFetch = "https://pokeapi.co/api/v2/pokemon/%d"

func FetchPokemon(pokemonId uint) (*PokemonJson, error) {
	url := fmt.Sprintf(urlFetch, pokemonId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if !strings.Contains(res.Status, "200") {
		return nil, errors.New("Request failed")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var pokemon PokemonJson
	json.Unmarshal([]byte(body), &pokemon)

	return &pokemon, err
}

func processPokemon(idChannel chan uint, pokemonChannel chan *PokemonJson) {
	for id := range idChannel {
		res, err := FetchPokemon(id)
		if err != nil {
			log.Fatalf("Error to fetch: %s", err)
		} else {
			log.Println(res.printReturn())
			pokemonChannel <- res
		}
	}
}

func transformPokemon(pokemonJson chan *PokemonJson, pokemonChannel chan *Pokemon) {
	for pokemon := range pokemonJson {
		toPokemon := pokemon.ToPokemonModel()
		pokemonChannel <- toPokemon
	}
	close(pokemonChannel)
}

func savePokemon(db *gorm.DB, pokemonChannel chan *Pokemon) {
	for pokemon := range pokemonChannel {
		err := db.Create(pokemon).Error
		if err != nil {
			log.Fatalln("Couldn't save record: ", err)
		}
	}
}
