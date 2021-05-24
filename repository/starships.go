package repository

import (
	"fmt"
	"github.com/imroc/req"
	"post-example/model"
)

type IStarship interface {
	GetStarships() (*model.StarshipsResponse, error)
}

type Starship struct {}

func (Starship) GetStarships() (*model.StarshipsResponse, error) {
	r, err := req.Get("https://swapi.dev/api/starships/")
	if err != nil {
		fmt.Sprintf("Error request to url: %s", err.Error())
		return nil, err
	}
	response := model.StarshipsResponse{}
	err = r.ToJSON(&response)
	if err != nil {
		fmt.Sprintf("Error on decoding json: %s", err.Error())
		return nil, err
	}
	return &response, nil
}

func NewStarships() Starship {
	return Starship{}
}