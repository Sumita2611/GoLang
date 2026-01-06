package models

type Movie struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Genre   string  `json:"genre"`
	Rating  float32 `json:"rating"`
	IsPrime bool    `json:"isPrime"`
}

