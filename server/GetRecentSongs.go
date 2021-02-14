package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"

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
	Name      string
	ID        string
	Artists   []string
	Attribute Attribute
}

type AttributeList struct {
	Attributes []Attribute `json:"audio_features"`
}

type Attribute struct {
	Danceability     float32 `json:"danceability"`
	Energy           float32 `json:"energy"`
	Key              int     `json:"key"`
	Loudness         float32 `json:"loudness"`
	Mode             int     `json:"mode"`
	Speechiness      float32 `json:"speechiness"`
	Acousticness     float32 `json:"acousticness"`
	Instramentalness float32 `json:"instrumentalness"`
	Liveness         float32 `json:"liveness"`
	Valence          float32 `json:"valence"`
	Tempo            float32 `json:"tempo"`
	Type             string  `json:"type"`
	ID               string  `json:"id"`
	Href             string  `json:"track_href"`
	Duration         int     `json:"duration_ms"`
	TimeSignature    int     `json:"4"`
}

func getIDParam(songList *SongList) string {
	idList := []string{}

	for i := 0; i < len(songList.Songs); i++ {
		idList = append(idList, songList.Songs[i].ID)
	}

	return strings.Join(idList[:], ",")
}

func appendAttributes(attributeList AttributeList, songList *SongList) *SongList {

	for i := 0; i < len(songList.Songs); i++ {
		songID := songList.Songs[i].ID
		for j := 0; j < len(attributeList.Attributes); j++ {
			if songID == attributeList.Attributes[j].ID {
				songList.Songs[i].Attribute = attributeList.Attributes[j]
				break
			}
		}
	}
	return songList
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

const (
	RecentsCollection = "recents"
)

func init() {
	// Pre-declared to avoid shadowing client
	var err error

	fsClient, err = firestore.NewClient(context.Background(), os.Getenv("GCP_PROJECT"))
	if err != nil {
		log.Fatalf("firestore new client: %v\n", err)
	}
}

func GetRecentlyPlayed(w http.ResponseWriter, r *http.Request) {

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

	var parsed Tracks
	serialBody, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal([]byte(serialBody), &parsed)

	// Extract song info and build new song list struct
	trackList := convertJSON(&parsed)
	idList := getIDParam(trackList)

	// Make another get request to
	req, err = http.NewRequest("GET", "https://api.spotify.com/v1/audio-features", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	attrParams := req.URL.Query()
	attrParams.Add("ids", idList)
	req.URL.RawQuery = attrParams.Encode()

	res, refreshed, err = client.Do(req)
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

	serialAttributes, _ := ioutil.ReadAll(res.Body)

	var attributes AttributeList
	json.Unmarshal([]byte(serialAttributes), &attributes)

	songList := appendAttributes(attributes, trackList)

	// Write results to database
	userReq, _ := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/me", nil)
	res, refreshed, err = client.Do(userReq)
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

	// If Spotify user has no email on account, skip this step
	if spotifyUser.Email != "" {
		currTime := time.Now().Unix()
		_, err = fsClient.Collection(UsersCollection).Doc(spotifyUser.Email).
			Collection(RecentsCollection).Doc(strconv.FormatInt(currTime, 10)).
			Set(r.Context(), songList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&songList); err != nil {
		panic(err)
	}
}
