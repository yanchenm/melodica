package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yanchenm/musemoods/spotify"
)

type Tracks struct {
	Tracks []struct {
		Data Track `json:"track"`
	} `json:"items"`
}

type Track struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
}

type SongList struct {
	Songs []Song `json:"songs"`
}

type Song struct {
	Name    string
	ID      string
	Artists []string
}

func convertJSON(trackList *Tracks) *SongList {

	// Empty song list
	songList := SongList{}

	for i := 0; i < len(trackList.Tracks); i++ {
		singleTrack := trackList.Tracks[i].Data
		parsedList := []string{}

		// Parse artist list since it's in weird format
		for i := 0; i < len(singleTrack.Artists); i++ {
			artist := (strings.Replace(singleTrack.Artists[i].Name, "{", "", -1))
			parsedList = append(parsedList, artist)
		}

		// Create new song
		newSong := Song{
			Name:    singleTrack.Name,
			ID:      singleTrack.ID,
			Artists: parsedList,
		}
		// Add song to song list
		songList.Songs = append(songList.Songs, newSong)
	}

	return &songList
}

func GetRecentlyPlayed(w http.ResponseWriter, r *http.Request) {

	// Preflight CORS
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, *")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, *")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

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

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/recently-played?", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	queryParams := req.URL.Query()
	queryParams.Add("limit", "50")
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

	serialBody, _ := ioutil.ReadAll(res.Body)

	var parsed Tracks
	json.Unmarshal([]byte(serialBody), &parsed)

	// Extract song info and build new song list struct
	trackList := convertJSON(&parsed)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&trackList); err != nil {
		panic(err)
	}
}
