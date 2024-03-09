package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemon(url string) (Pokemon, error) {
	body, err := HttpRequest(c, "GET", url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	decodingErr := json.Unmarshal(body, &pokemon)
	if decodingErr != nil {
		return Pokemon{}, decodingErr
	}

	return pokemon, nil
}
