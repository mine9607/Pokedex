package main

type Config struct {
	Next_URL     string
	Previous_URL string
}

func NewConfig() *Config {
	// initialize a pointer to the config struct with an empty previous_URL and the input url as the Next_URL
	c := &Config{Previous_URL: "", Next_URL: ""}
	return c
}

func UpdateConfigURLs(prev_url string, next_url string, c *Config) {
	// update the previousURL to Next_URL and update Next_URL to input url then return that url
	c.Previous_URL = prev_url
	c.Next_URL = next_url
}
