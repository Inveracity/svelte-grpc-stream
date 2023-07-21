// main.go
package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Email    string
	Password string
	BaseURL  string
}

type AuthToken struct {
	Token string `json:"token"`
	Admin struct {
		ID      string `json:"id"`
		Created string `json:"created"`
		Updated string `json:"updated"`
		Email   string `json:"email"`
		Avatar  int    `json:"avatar"`
	} `json:"admin"`
}

type Users struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
	Items      []struct {
		Avatar          string `json:"avatar"`
		CollectionID    string `json:"collectionId"`
		CollectionName  string `json:"collectionName"`
		Created         string `json:"created"`
		Email           string `json:"email"`
		EmailVisibility bool   `json:"emailVisibility"`
		ID              string `json:"id"`
		Name            string `json:"name"`
		Updated         string `json:"updated"`
		Username        string `json:"username"`
		Verified        bool   `json:"verified"`
	} `json:"items"`
}

func New(BaseURL, email, password string) *Auth {
	return &Auth{
		Email:    email,
		Password: password,
		BaseURL:  BaseURL,
	}
}

// VerifyUserToken verifies a user token from the PocketBase API
func (a *Auth) VerifyUserToken(usertoken string) (bool, error) {

	claims, err := a.decodeUserToken(usertoken)
	if err != nil {
		return false, err
	}

	if a.tokenExpired(claims) {
		return false, fmt.Errorf("token expired")
	}

	admintoken, err := a.getAdminToken()
	if err != nil {
		return false, err
	}

	users, err := a.lookupUser(claims["id"].(string), admintoken)
	if err != nil {
		return false, err
	}

	if len(users.Items) == 0 {
		log.Println("No user found")
		return false, nil
	}

	if users.Items[0].ID != claims["id"].(string) {
		log.Println("User ID mismatch")
		return false, nil
	}

	return true, nil
}

// GetAdminToken gets an admin token from the PocketBase API
func (a *Auth) getAdminToken() (string, error) {
	postBody, _ := json.Marshal(map[string]string{
		"identity": a.Email,
		"password": a.Password,
	})

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(a.BaseURL+"/api/admins/auth-with-password", "application/json", responseBody)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("pocketbase returned status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	authToken := AuthToken{}

	if err := json.Unmarshal(body, &authToken); err != nil {
		return "", err
	}

	return authToken.Token, nil
}

// DecodeUserToken decodes a PocketBase user token
func (a *Auth) decodeUserToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	p := jwt.Parser{}
	_, _, err := p.ParseUnverified(tokenString, claims)

	if err != nil {
		return nil, fmt.Errorf("unable to parse token: %s", err)
	}

	return claims, nil
}

// LookupUser looks up a user from the PocketBase API
func (a *Auth) lookupUser(userID, adminToken string) (Users, error) {
	client := &http.Client{}
	URL, err := url.Parse(a.BaseURL + "/api/collections/users/records")
	if err != nil {
		return Users{}, err
	}

	q := URL.Query()
	q.Add("filter", fmt.Sprintf(`id="%s"`, userID))
	URL.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return Users{}, err
	}

	req.Header.Add("Authorization", adminToken)

	resp, err := client.Do(req)
	if err != nil {
		return Users{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Users{}, err
	}

	users := Users{}
	if err := json.Unmarshal(body, &users); err != nil {
		return Users{}, err
	}

	return users, nil
}

// tokenExpired checks if a token is expired
func (a *Auth) tokenExpired(claims jwt.MapClaims) bool {
	return claims["exp"].(float64) < float64(time.Now().Unix())
}
