package pokeapi

import (
	"io"
	"net/http"
	"encoding/json"
)

func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// if val, exists := c.cache.Get(url); exists {
	// 	pokemon := Pokemon{}
	// 	err := json.Unmarshal(val, &pokemon)
	// 	if err != nil {
	// 		return Pokemon{}, nil
	// 	}

	// 	return pokemon, nil
	// }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// c.cache.Add(url, dat)
	
	newPokemon := Pokemon{}
	err = json.Unmarshal(dat, &newPokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return newPokemon, nil
} // End GetLocations() func