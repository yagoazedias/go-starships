package respository

import (
	"fmt"
	"post-example/model"
	"post-example/respository"
)

type IStarships interface {
	GetStarships() ([]model.Starship, error)
}

type Starships struct {
	repository *respository.Starship
}

func (s *Starships) GetStarships() ([]model.Starship, error) {
	s.repository = respository.NewStarship()
	apiResponse, err := s.repository.GetStarships()
	if err != nil {
		fmt.Sprintf("Error request to api: %s", err.Error())
		return nil, err
	}
	return apiResponse.Results, nil
}