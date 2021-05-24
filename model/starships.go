package model

import (
	"time"
)

type Starship struct {
	Name                 string        `json:"name"`
	Model                string        `json:"model"`
	Manufacturer         string        `json:"manufacturer"`
	CostInCredits        string        `json:"cost_in_credits"`
	Length               string        `json:"length"`
	MaxAtmosphereSpeed   string          `json:"max_atmosphering_speed"`
	Crew                 string        `json:"crew"`
	Passengers           string        `json:"passengers"`
	CargoCapacity        string        `json:"cargo_capacity"`
	Consumables          string        `json:"consumables"`
	HyperdriveRating     string        `json:"hyperdrive_rating"`
	Mglt                 string        `json:"MGLT"`
	StarshipClass        string        `json:"starship_class"`
	Pilots               []interface{} `json:"pilots"`
	Films                []string      `json:"films"`
	Created              time.Time     `json:"created"`
	Edited               time.Time     `json:"edited"`
	URL                  string        `json:"url"`
}

func NewStarship(
	name string,
	model string,
	manufacturer string,
	costInCredits string,
	length string,
	maxAtmosphereSpeed string,
	crew string,
	passengers string,
	cargoCapacity string,
	consumables string,
	hyperdriveRating string,
	mglt string,
	starshipClass string,
	pilots []interface{},
	films []string,
	created time.Time,
	edited time.Time,
	URL string,
	) *Starship {

	return &Starship{
		Name: name,
		Model: model,
		Manufacturer: manufacturer,
		CostInCredits: costInCredits,
		Length: length,
		MaxAtmosphereSpeed: maxAtmosphereSpeed,
		Crew: crew,
		Passengers: passengers,
		CargoCapacity: cargoCapacity,
		Consumables: consumables,
		HyperdriveRating: hyperdriveRating,
		Mglt: mglt,
		StarshipClass: starshipClass,
		Pilots: pilots,
		Films: films,
		Created: created,
		Edited: edited,
		URL: URL}
}

type StarshipsResponse struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Starship  `json:"results"`
}
