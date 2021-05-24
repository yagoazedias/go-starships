package service

import (
	"github.com/golang/mock/gomock"
	"post-example/mock"
	"post-example/model"
	"post-example/repository"
	"testing"
)

func TestStarshipsServiceSuccessResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := repository.NewMockIStarship(ctrl)
	starship := Starships{repoMock}
	defer ctrl.Finish()

	strings := []string{"test1", "test2"}

	// GetStarships should return the 'expectedBody' and no error
	repoMock.EXPECT().GetStarships().Return(GetMockStarshipResponseModel(strings), nil)
	response, err := starship.GetStarships(); if err != nil {
		t.Error("Error should not be nil")
	}

	if len(strings) != len(response) {
		t.Errorf("length: %d does not match with %d", len(strings), len(response))
	}

	for i, value := range response {
		if value.Name != strings[i] {
			t.Errorf("Name: %s does not match with %s", value.Name, strings[i])
		}
	}
}

func GetMockStarshipResponseModel(strings []string) *model.StarshipsResponse {
	return mock.NewStarshipsResponse(strings)
}