package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type OAuthAccessToken struct {
	tokenType    string
	accessToken  string
	expiresAt    int64
	refreshToken string
	userId       string
}

type OAuthAccessTokenResponse struct {
	TokenType       string `json:"token_type"`
	ExpiresIn       int64  `json:"expires_in"`
	AccessTokenJWT  string `json:"access_token"`
	RefreshTokenJWT string `json:"refresh_token"`
}

type EventFarmAccessTokenResponse struct {
	OAuthAccessTokenResponse
	ExpiresAt int64  `json:"expires_at"`
	UserId    string `json:"user_id"`
}

func NewOAuthAccessTokenFromResponse(r *http.Response) (oAuthAccessToken *OAuthAccessToken, err error) {
	if r.StatusCode != 200 {
		err = errors.New(`Unable to load access token`)
		LogResponse(r)
		return
	}

	bodyBuff, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	eventFarmAccessTokenResponse := EventFarmAccessTokenResponse{}
	err = json.Unmarshal(bodyBuff, &eventFarmAccessTokenResponse)
	if err != nil {
		return
	}

	accessToken, err := getTokenFromJWT(eventFarmAccessTokenResponse.AccessTokenJWT, "")
	if err != nil {
		return
	}

	oAuthAccessToken = &OAuthAccessToken{
		tokenType:    eventFarmAccessTokenResponse.TokenType,
		expiresAt:    accessToken.expiresAt,
		userId:       accessToken.userId,
		accessToken:  accessToken.token,
		refreshToken: "", // wip decode the refresh token
	}

	return
}

type OAuthToken struct {
	token     string
	userId    string
	expiresAt int64
}

func getTokenFromJWT(token string, secret string) (oauthToken *OAuthToken, err error) {
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	jti := claims["jti"].(string)
	sub := claims["sub"].(string)
	exp := claims["exp"].(float64)

	return &OAuthToken{
		token:     jti,
		userId:    sub,
		expiresAt: int64(exp),
	}, nil
}

func (t OAuthAccessToken) getHeaderString() string {
	return t.tokenType + ` ` + t.accessToken
}

func (t *OAuthAccessToken) isExpired() bool {
	return time.Now().Unix() > t.expiresAt
}
