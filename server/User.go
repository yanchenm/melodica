package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserResponse struct {
	Email string `json:"email"`
	Name string `json:"name"`
}

type SpotifyUser struct {
	Email string `json:"email"`
	Name string `json:"display_name"`
}

func UserHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tokens, ok := r.Header["Authorization"]
	if !ok {
		http.Error(w, "No authorization token", http.StatusUnauthorized)
		return
	}

	token := tokens[0]
	if token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	req.Header.Set("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	spotifyUser := SpotifyUser{}
	err = json.Unmarshal(body, &spotifyUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	userResponse := UserResponse{spotifyUser.Email, spotifyUser.Name}
	if err = json.NewEncoder(w).Encode(&userResponse); err != nil {
		panic(err)
	}
}