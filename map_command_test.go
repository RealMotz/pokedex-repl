package main

import (
	"testing"
)

func TestGenerateUrl(t *testing.T) {
	const baseUrl = "http://test.com/"
	const endpoint = "test"
  const pagination = "?limit=20"

	t.Run("Test case # 1", func(t *testing.T) {
		const expected = baseUrl + pagination
		actual := Concatenate(baseUrl, "", pagination)
		if actual != expected {
			t.Errorf("urls do not match")
		}
	})
	t.Run("Test case # 2", func(t *testing.T) {
		const expected = baseUrl + endpoint + pagination
		actual := Concatenate(baseUrl, endpoint, pagination)
		if actual != expected {
			t.Errorf("urls do not match")
		}
	})
}
