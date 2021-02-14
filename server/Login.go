package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/yanchenm/musemoods/spotify"
)

type LoginRequest struct {
	AccessCode  string `json:"access_code"`
	RedirectURI string `json:"redirect_uri"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	loginRequest := LoginRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if loginRequest.AccessCode == "" || loginRequest.RedirectURI == "" {
		http.Error(w, "Invalid access code", http.StatusBadRequest)
		return
	}

	authData := url.Values{}
	authData.Set("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
	authData.Set("client_secret", os.Getenv("SPOTIFY_CLIENT_SECRET"))
	authData.Set("grant_type", "authorization_code")
	authData.Set("code", loginRequest.AccessCode)
	authData.Set("redirect_uri", loginRequest.RedirectURI)

	res, err := http.Post(spotify.AuthURL, "application/x-www-form-urlencoded", strings.NewReader(authData.Encode()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("request failed: %v", res)
		http.Error(w, "Failed to obtain authorization from Spotify", http.StatusUnauthorized)
		return
	}

	defer res.Body.Close()
	authResponse := AuthResponse{}
	decoder = json.NewDecoder(res.Body)
	if err = decoder.Decode(&authResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access",
		Value:    authResponse.AccessToken,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh",
		Value:    authResponse.RefreshToken,
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
	return
}
