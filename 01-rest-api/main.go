package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"
)

const baseURL = "https://swapi.co/api/people"
const appPort = "8080"

type Character struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
	URL       string `json:"url,omitempty"`
}

type CharacterList struct {
	Results []Character `json:"results"`
}

func GetStarWarsCharacters(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}

	resp, err := http.Get(baseURL)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, http.StatusText(resp.StatusCode), resp.StatusCode)
		return
	}

	var response CharacterList
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	var payload []Character

	for _, character := range response.Results {
		u, _ := url.Parse(character.URL)
		character.ID = path.Base(u.Path)
		character.URL = ""
		payload = append(payload, character)
	}

	respondWithJSON(w, payload, http.StatusOK)
}

func respondWithJSON(w http.ResponseWriter, payload interface{}, code int) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	log.Printf("Starting server on port %s", appPort)
	http.HandleFunc("/starwars/characters", GetStarWarsCharacters)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
