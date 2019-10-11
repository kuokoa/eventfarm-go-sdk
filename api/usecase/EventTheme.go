/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type EventTheme struct {
	restClient rest.RestClientInterface
}

func NewEventTheme(restClient rest.RestClientInterface) *EventTheme {
	return &EventTheme{restClient}
}

// GET: Queries

type GetEventThemeForEventParameters struct {
	EventId string
}

func (t *EventTheme) GetEventThemeForEvent(p *GetEventThemeForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/EventTheme/UseCase/GetEventThemeForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type SetEventThemeParameters struct {
	EventId             string
	PrimaryColor        string
	PrimaryAltColor     string
	AlertColor          string
	BackgroundColor     string
	FontColor           string
	ThemeNameType       string
	ThemeFontFamilyType string
}

func (t *EventTheme) SetEventTheme(p *SetEventThemeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`primaryColor`, p.PrimaryColor)
	queryParameters.Add(`primaryAltColor`, p.PrimaryAltColor)
	queryParameters.Add(`alertColor`, p.AlertColor)
	queryParameters.Add(`backgroundColor`, p.BackgroundColor)
	queryParameters.Add(`fontColor`, p.FontColor)
	queryParameters.Add(`themeNameType`, p.ThemeNameType)
	queryParameters.Add(`themeFontFamilyType`, p.ThemeFontFamilyType)

	return t.restClient.Post(
		`/v2/EventTheme/UseCase/SetEventTheme`,
		&queryParameters,
		nil,
		nil,
	)
}
