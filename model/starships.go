package model

import "time"

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

type StarshipsResponse struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Starship  `json:"results"`
}