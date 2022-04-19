package pokemon

import "github.com/rafaelpardor/pokego/pkg/database"

func ETL() {
	pokemonIdChannel := make(chan uint)
	pokemonJsonChannel := make(chan *PokemonJson)
	pokemonChannel := make(chan *Pokemon)

	db := database.Get()

	go processPokemon(pokemonIdChannel, pokemonJsonChannel)
	go transformPokemon(pokemonJsonChannel, pokemonChannel)
	go savePokemon(db, pokemonChannel)

	for i := 1; i <= 10; i++ {
		pokemonIdChannel <- uint(i)
	}

	close(pokemonIdChannel)
}
