package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://test.com",
			val: []byte("tesdata"),
		},
		{
			key: "https://test.com",
			val: []byte("tesdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
  const baseTime = 5*time.Millisecond
  const waitTime = baseTime + 5*time.Millisecond
  const url = "http://test.com"
  cache := NewCache(baseTime)
  cache.Add(url, []byte("testdata"))

  _, ok := cache.Get(url)
  if !ok {
    t.Errorf("expected to find key")
    return
  }

  time.Sleep(waitTime)

  _, ok = cache.Get(url)
  if ok {
    t.Errorf("expected not to find key")
  }
}


