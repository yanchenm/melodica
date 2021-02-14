package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yanchenm/musemoods/spotify"
)

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type SpotifyUser struct {
	Email string `json:"email"`
	Name  string `json:"display_name"`
}

func UserHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	refreshCookie, err := r.Cookie("refresh")
	if err != nil {
		http.Error(w, "No refresh cookie", http.StatusUnauthorized)
		return
	}

	accessCookie, err := r.Cookie("access")
	if err != nil {
		http.Error(w, "No access cookie", http.StatusUnauthorized)
		return
	}

	refreshToken := refreshCookie.Value
	accessToken := accessCookie.Value

	client := spotify.Initialize(accessToken, refreshToken)
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)

	res, refreshed, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if refreshed {
		http.SetCookie(w, &http.Cookie{
			Name:     "access",
			Value:    client.GetAccessToken(),
			HttpOnly: true,
		})
	}

	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		w.WriteHeader(res.StatusCode)
		_, _ = io.Copy(w, res.Body)
		return
	}

	defer res.Body.Close()
	spotifyUser := SpotifyUser{}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&spotifyUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	userResponse := UserResponse{spotifyUser.Email, spotifyUser.Name}
	if err = json.NewEncoder(w).Encode(&userResponse); err != nil {
		panic(err)
	}
}
