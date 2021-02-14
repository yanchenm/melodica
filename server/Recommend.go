package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yanchenm/musemoods/spotify"
)

type Recommendations struct {
	Data []SingleRec `json:"tracks"`
}

type SingleRec struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Album   Album  `json:"album"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	URL struct {
		SpotifyURL string `json:"spotify"`
	} `json:"external_urls"`
}

type Album struct {
	Images []Image `json:"images"`
}

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

func GetRecommended(w http.ResponseWriter, r *http.Request) {
	// Preflight CORS
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Origin", "https://melodica.tech")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, ")
	w.Header().Set("Access-Control-Allow-Origin", "https://melodica.tech")

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

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/recommendations", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emotion, ok := r.URL.Query()["emotion"]
	if !ok {
		http.Error(w, "No emotion query parameter", http.StatusInternalServerError)
		return
	}
	queryParams := req.URL.Query()
	queryParams.Add("seed_genres", strings.Join(emotion[:], ","))
	queryParams.Add("market", "CA")
	req.URL.RawQuery = queryParams.Encode()

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

	var recommendations Recommendations
	serialBody, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal([]byte(serialBody), &recommendations)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&recommendations); err != nil {
		panic(err)
	}
}
