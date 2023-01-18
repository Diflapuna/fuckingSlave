package models

type Response struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Id        int    `json:"id"`
}

type Joke struct {
	Joke string `json:"joke"`
}
