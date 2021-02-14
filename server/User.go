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

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	combinedCookie, err := r.Cookie("__session")
	if err != nil {
		s := fmt.Sprintf("No session cookie: %v", err)
		http.Error(w, s, http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := spotify.SplitTokens(combinedCookie.Value)
	if err != nil {
		s := fmt.Sprintf("Invalid cookie: %v", err)
		http.Error(w, s, http.StatusUnauthorized)
		return
	}

	client := spotify.Initialize(accessToken, refreshToken)
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)

	res, refreshed, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if refreshed {
		newCookie := client.GetCombinedToken()
		http.SetCookie(w, &http.Cookie{
			Name:     "__session",
			Value:    newCookie,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Secure:   true,
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
