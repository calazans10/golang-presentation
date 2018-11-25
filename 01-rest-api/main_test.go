package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestGetStarWarsCharacters(t *testing.T) {
	tests := []struct {
		name                 string
		mockedHTPPMethod     string
		mockedResponse       *http.Response
		mockedError          error
		expectedResponseCode int
		expectedResponseBody string
	}{
		{
			name:                 "returns 405 on receiving method not allowed",
			mockedHTPPMethod:     http.MethodPost,
			mockedResponse:       nil,
			mockedError:          nil,
			expectedResponseCode: http.StatusMethodNotAllowed,
			expectedResponseBody: "Method Not Allowed\n",
		},
		{
			name:                 "returns 500 on connection refused",
			mockedHTPPMethod:     http.MethodGet,
			mockedResponse:       nil,
			mockedError:          errors.New("An error"),
			expectedResponseCode: http.StatusInternalServerError,
			expectedResponseBody: "Internal Server Error\n",
		},
		{
			name:                 "returns 404 on missing resource",
			mockedHTPPMethod:     http.MethodGet,
			mockedResponse:       httpmock.NewBytesResponse(http.StatusNotFound, []byte(`{"detail": "Not found"}`)),
			mockedError:          nil,
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: "Not Found\n",
		},
		{
			name:                 "returns 400 on decoding invalid resource",
			mockedHTPPMethod:     http.MethodGet,
			mockedResponse:       httpmock.NewStringResponse(http.StatusOK, ""),
			mockedError:          nil,
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: "Bad Request\n",
		},
		{
			name:             "returns 200 on getting resource",
			mockedHTPPMethod: http.MethodGet,
			mockedResponse: httpmock.NewBytesResponse(http.StatusOK, []byte(`{
				"count": 87,
				"next": "https://swapi.co/api/people/?page=2",
				"previous": null,
				"results": [
					{
						"name": "Luke Skywalker",
						"height": "172",
						"mass": "77",
						"hair_color": "blond",
						"skin_color": "fair",
						"eye_color": "blue",
						"birth_year": "19BBY",
						"gender": "male",
						"homeworld": "https://swapi.co/api/planets/1/",
						"films": [
							"https://swapi.co/api/films/2/",
							"https://swapi.co/api/films/6/",
							"https://swapi.co/api/films/3/",
							"https://swapi.co/api/films/1/",
							"https://swapi.co/api/films/7/"
						],
						"species": [
							"https://swapi.co/api/species/1/"
						],
						"vehicles": [
							"https://swapi.co/api/vehicles/14/",
							"https://swapi.co/api/vehicles/30/"
						],
						"starships": [
							"https://swapi.co/api/starships/12/",
							"https://swapi.co/api/starships/22/"
						],
						"created": "2014-12-09T13:50:51.644000Z",
						"edited": "2014-12-20T21:17:56.891000Z",
						"url": "https://swapi.co/api/people/1/"
				    	}
				]
			}`)),
			mockedError:          nil,
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: "[{\"id\":\"1\",\"name\":\"Luke Skywalker\",\"height\":\"172\",\"mass\":\"77\",\"hair_color\":\"blond\",\"skin_color\":\"fair\",\"eye_color\":\"blue\",\"birth_year\":\"19BBY\",\"gender\":\"male\"}]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder(
				http.MethodGet,
				"https://swapi.co/api/people",
				func(req *http.Request) (*http.Response, error) {
					return test.mockedResponse, test.mockedError
				},
			)

			request, _ := http.NewRequest(test.mockedHTPPMethod, "/starwars/characters", nil)
			response := httptest.NewRecorder()

			GetStarWarsCharacters(response, request)

			assert.Equal(t, test.expectedResponseCode, response.Code)
			assert.Equal(t, test.expectedResponseBody, response.Body.String())
		})
	}
}
