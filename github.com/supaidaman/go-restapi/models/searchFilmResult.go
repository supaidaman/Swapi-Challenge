package models

import (
	"github.com/peterhellberg/swapi"
)

// SearchFilmResult struct for searchs
type SearchFilmResult struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []swapi.Planet `json:"results"`
}
