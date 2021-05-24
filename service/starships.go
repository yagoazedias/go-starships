package service

import (
	"fmt"
	"post-example/model"
	"post-example/repository"
)

type IStarships interface {
	GetStarships() ([]model.Starship, error)
}

type Starships struct {
	repository repository.IStarship
}

func (s Starships) GetStarships() ([]model.Starship, error) {
	apiResponse, err := s.repository.GetStarships()
	if err != nil {
		fmt.Sprintf("Error request to api: %s", err.Error())
		return nil, err
	}
	return apiResponse.Results, nil
}

func NewStarships() IStarships {
	return Starships{
		repository: repository.NewStarships(),
	}
}