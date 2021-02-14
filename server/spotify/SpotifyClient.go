package spotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	TokenExpiredError = "The access token expired"
	AuthURL           = "https://accounts.spotify.com/api/token"
)

type Client struct {
	accessToken  string
	refreshToken string
	client       *http.Client
}

type ErrorResponse struct {
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func Initialize(accessToken, refreshToken string) *Client {
	httpClient := &http.Client{}

	if !strings.HasPrefix(accessToken, "Bearer ") {
		accessToken = "Bearer " + accessToken
	}

	if !strings.HasPrefix(refreshToken, "Bearer ") {
		refreshToken = "Bearer " + refreshToken
	}

	return &Client{
		accessToken:  accessToken,
		refreshToken: refreshToken,
		client:       httpClient,
	}
}

func CombineTokens(access, refresh string) string {
	return access + "|" + refresh
}

func SplitTokens(combined string) (string, string, error) {
	tokens := strings.Split(combined, "|")
	if len(tokens) != 2 {
		return "", "", fmt.Errorf("token not in valid format")
	}

	return tokens[0], tokens[1], nil
}

func (c *Client) GetAccessToken() string {
	return c.accessToken
}

func (c *Client) GetCombinedToken() string {
	return CombineTokens(c.accessToken, c.refreshToken)
}

func (c *Client) Do(req *http.Request) (*http.Response, bool, error) {
	// Add bearer token to request header
	req.Header.Set("Authorization", c.accessToken)

	// Try request
	res, err := c.client.Do(req)
	if err != nil {
		return nil, false, err
	}

	// Any response that isn't a token expired error, return to user
	if res.StatusCode != http.StatusUnauthorized {
		return res, false, nil
	}

	defer res.Body.Close()
	errRes := ErrorResponse{}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&errRes); err != nil {
		return nil, false, err
	}

	if errRes.Error.Message != TokenExpiredError {
		return res, false, nil
	}

	// If token expired, try refresh
	refreshBody, err := json.Marshal(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": c.refreshToken,
	})
	refreshReq, err := http.NewRequest(http.MethodPost, AuthURL, bytes.NewBuffer(refreshBody))
	if err != nil {
		return nil, false, err
	}

	refreshRes, err := c.client.Do(refreshReq)
	if err != nil {
		return nil, false, err
	}

	defer refreshRes.Body.Close()
	refreshResponse := AuthResponse{}
	decoder = json.NewDecoder(refreshRes.Body)
	if err = decoder.Decode(&refreshResponse); err != nil {
		return nil, false, err
	}

	c.accessToken = refreshResponse.AccessToken

	if refreshResponse.RefreshToken != "" {
		c.refreshToken = refreshResponse.RefreshToken
	}

	// Try request again
	req.Header.Set("Authorization", c.accessToken)
	res, err = c.client.Do(req)
	if err != nil {
		return nil, true, err
	}

	return res, true, nil
}
