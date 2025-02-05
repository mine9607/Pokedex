package pokecache

import (
	"fmt"
	"testing"
	"time"
)

// Note: because the Function Add doesn't return anything we need to test both Add and Get

func TestCacheAddGetFunction(t *testing.T) {
	const interval = 5 * time.Second
	tests := []struct {
		key string
		val []byte
	}{
		{
			key: "https://google.com",
			val: []byte("testData"),
		},
		{
			key: "https://boot.dev",
			val: []byte("moretestdata"),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(test.key, test.val)
			val, ok := cache.Get(test.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(test.val) {
				t.Errorf("expected to find %s. got=%s\n", string(test.val), string(val))
				return
			}
		})
	}
}

func TestReaperLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	cache := NewCache(baseTime)
	cache.Add("https://google.com", []byte("testdata"))

	_, ok := cache.Get("https://google.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://google.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
