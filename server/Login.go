package server

import (
	"encoding/json"
	"io"
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

func Login(w http.ResponseWriter, r *http.Request) {
	// Preflight CORS
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Origin", "https://melodica.tech")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "https://melodica.tech")

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
		defer res.Body.Close()
		w.WriteHeader(res.StatusCode)
		_, _ = io.Copy(w, res.Body)
		return
	}

	defer res.Body.Close()
	authResponse := spotify.AuthResponse{}
	decoder = json.NewDecoder(res.Body)
	if err = decoder.Decode(&authResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Concat tokens and add all to __session token
	cookieVal := spotify.CombineTokens(authResponse.AccessToken, authResponse.RefreshToken)
	http.SetCookie(w, &http.Cookie{
		Name:     "__session",
		Value:    cookieVal,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
	w.WriteHeader(http.StatusOK)
	return
}
