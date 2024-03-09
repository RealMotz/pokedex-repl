package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocation(url string) (PokeLocation, error) {
	body, err := HttpRequest(c, "GET", url)
	if err != nil {
		return PokeLocation{}, err
	}

	location := PokeLocation{}
	decodingErr := json.Unmarshal(body, &location)
	if decodingErr != nil {
		return PokeLocation{}, decodingErr
	}

	return location, nil
}

func (c *Client) GetArea(url string) (PokeArea, error) {
	body, err := HttpRequest(c, "GET", url)
	if err != nil {
		return PokeArea{}, err
	}
	area := PokeArea{}
	decodingErr := json.Unmarshal(body, &area)
	if decodingErr != nil {
		return PokeArea{}, decodingErr
	}

	return area, nil
}
