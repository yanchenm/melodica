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

	var recommendations Recommendations
	serialBody, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal([]byte(serialBody), &recommendations)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&recommendations); err != nil {
		panic(err)
	}
}
