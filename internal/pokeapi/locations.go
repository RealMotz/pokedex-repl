package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpRequest(c *Client, method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
    return []byte{}, err
	}
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func (c *Client) GetLocation(url string) PokeLocation {
	body, err := HttpRequest(c, "GET", url)
	if err != nil {
		fmt.Println(err)
		return PokeLocation{}
	}

	location := PokeLocation{}
	decodingErr := json.Unmarshal(body, &location)
	if decodingErr != nil {
		fmt.Println(decodingErr)
		return PokeLocation{}
	}

	return location
}

func (c *Client) GetArea(url string) (PokeArea, error) {
	body, err := HttpRequest(c, "GET", url)
	if err != nil {
		fmt.Println(err)
		return PokeArea{}, err
	}
	area := PokeArea{}
	decodingErr := json.Unmarshal(body, &area)
	if decodingErr != nil {
		fmt.Println(decodingErr)
		return PokeArea{}, decodingErr
	}

	return area, nil
}
