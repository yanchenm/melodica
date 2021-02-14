package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/firestore"

	"github.com/yanchenm/musemoods/spotify"
)

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	New   bool   `json:"new"`
}

type SpotifyUser struct {
	Email string `json:"email"`
	Name  string `json:"display_name"`
}

const (
	UsersCollection = "users"
)

var fsClient *firestore.Client

func init() {
	// Pre-declared to avoid shadowing client
	var err error

	fsClient, err = firestore.NewClient(context.Background(), os.Getenv("GCP_PROJECT"))
	if err != nil {
		log.Fatalf("firestore new client: %v\n", err)
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	// Preflight CORS
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, *")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Origin", "https://melodica.tech")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, *")
	w.Header().Set("Access-Control-Allow-Origin", "https://melodica.tech")

	for _, cookie := range r.Cookies() {
		fmt.Println("Found a cookie named: ", cookie.Name)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	refreshCookie, err := r.Cookie("refresh")
	if err != nil {
		s := fmt.Sprintf("No refresh cookie: %v", err)
		http.Error(w, s, http.StatusUnauthorized)
		return
	}

	accessCookie, err := r.Cookie("access")
	if err != nil {
		s := fmt.Sprintf("No access cookie: %v", err)
		http.Error(w, s, http.StatusUnauthorized)
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

	// Check if user has used service before. Can't be done if user has no email on Spotify account
	userExists := false
	if spotifyUser.Email != "" {
		_, err = fsClient.Collection(UsersCollection).Doc(spotifyUser.Email).Get(r.Context())
		if err != nil {
			// Find out if error is because user doesn't exist
			s := err.Error()
			if strings.Contains(s, "code = NotFound") {
				userExists = false
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			userExists = true
		}

		// Create user if not exists
		if !userExists {
			_, err = fsClient.Collection(UsersCollection).Doc(spotifyUser.Email).Set(r.Context(), spotifyUser)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	userResponse := UserResponse{
		Email: spotifyUser.Email,
		Name:  spotifyUser.Name,
		New:   !userExists,
	}
	if err = json.NewEncoder(w).Encode(&userResponse); err != nil {
		panic(err)
	}
}
