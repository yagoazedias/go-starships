package mock

import (
	"post-example/model"
	"time"
)

func GetMockStarshipModel(name string) *model.Starship {
	return model.NewStarship(
		name,
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		[]interface{}{},
		[]string{""},
		time.Date(1998, 12, 1, 1,1,1,1, time.Local),
		time.Date(1998, 12, 1, 1,1,1,1, time.Local),
		"",
	)
}

func NewStarshipsResponse(names []string) *model.StarshipsResponse {
	var starships []model.Starship
	for _, name := range names {
		starships = append(starships, *GetMockStarshipModel(name))
	}
	return &model.StarshipsResponse{
		Count: 1,
		Next: "",
		Previous: 1,
		Results: starships,
	}
}