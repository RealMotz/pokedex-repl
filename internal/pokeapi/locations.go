package pokeapi

import (
  "encoding/json"
  "io"
  "net/http"
  "fmt"
)

type PokeLocation struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func (c *Client) GetLocation(url string) PokeLocation {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
    return PokeLocation{}
	}

  res, err := c.httpClient.Do(req)
  if err != nil {
    fmt.Println(err)
    return PokeLocation{}
  }

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
    return PokeLocation{}
	}

	location := PokeLocation{}
	decodingErr := json.Unmarshal(body, &location)
	if decodingErr != nil {
		fmt.Println(err)
    return PokeLocation{}
	}

	return location
}
