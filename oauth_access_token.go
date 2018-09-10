package gosdk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type OAuthAccessToken struct {
	tokenType    string
	accessToken  string
	expiresAt    int64
	refreshToken string
	userId       string
}

type OAuthAccessTokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type EventFarmAccessTokenResponse struct {
	OAuthAccessTokenResponse
	ExpiresAt int64  `json:"expires_at"`
	UserId    string `json:"user_id"`
}

func NewOAuthAccessTokenFromResponse(response *http.Response) (oAuthAccessToken *OAuthAccessToken, err error) {
	if response.StatusCode != 200 {
		err = errors.New(`Unable to load access token`)
		LogResponse(response)
		return
	}

	bodyBuff, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	eventFarmAccessTokenResponse := EventFarmAccessTokenResponse{}
	err = json.Unmarshal(bodyBuff, &eventFarmAccessTokenResponse)
	if err != nil {
		return
	}

	oAuthAccessToken = &OAuthAccessToken{
		tokenType:    eventFarmAccessTokenResponse.TokenType,
		expiresAt:    eventFarmAccessTokenResponse.ExpiresAt,
		userId:       eventFarmAccessTokenResponse.UserId,
		accessToken:  eventFarmAccessTokenResponse.AccessToken,
		refreshToken: eventFarmAccessTokenResponse.RefreshToken,
	}

	return
}

func (t OAuthAccessToken) getHeaderString() string {
	return t.tokenType + ` ` + t.accessToken
}

func (t *OAuthAccessToken) isExpired() bool {
	return time.Now().Unix() > t.expiresAt
}
