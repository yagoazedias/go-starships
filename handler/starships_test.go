package handler

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"post-example/mock"
	"post-example/model"
	"post-example/service"
	"reflect"
	"testing"
)

func TestStarshipsHandlerSuccessResponse(t *testing.T) {

	// Setting up writer, router, mock controller and starshipMock
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	ctrl := gomock.NewController(t)
	serviceMock := service.NewMockIStarships(ctrl)
	starship := Starship{serviceMock}
	defer ctrl.Finish()

	// It is necessary to create the var expectedBody to compare it with the body later
	expectedBody := []model.Starship{
		*mock.GetMockStarshipModel("Death Start"),
		*mock.GetMockStarshipModel("Millenium Falcon"),
	}

	// GetStarships should return the 'expectedBody' and no error
	serviceMock.EXPECT().GetStarships().Return(
		expectedBody, nil,
	)

	r.HandleFunc("/starships", starship.StarshipsHandler)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/starships", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	var response []model.Starship
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&response); err != nil {
		t.Error("Could not parse result from payload", err)
	}

	if !reflect.DeepEqual(response, expectedBody) {
		t.Error("Response is not equals to expectedBody")
	}
}

func TestStarshipsHandlerErrorResponse(t *testing.T) {

	// Setting up writer, router, mock controller and starshipMock
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	ctrl := gomock.NewController(t)
	serviceMock := service.NewMockIStarships(ctrl)
	starship := Starship{serviceMock}
	defer ctrl.Finish()

	// GetStarships should return the 'expectedBody' and no error
	serviceMock.EXPECT().GetStarships().Return(
		nil, errors.New("A generic error"),
	)

	r.HandleFunc("/starships", starship.StarshipsHandler)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/starships", nil))

	if w.Code != http.StatusInternalServerError {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		return
	}
}
