/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/eventfarm/go-sdk/rest"
)

type Promotion struct {
	restClient rest.RestClientInterface
}

func NewPromotion(restClient rest.RestClientInterface) *Promotion {
	return &Promotion{restClient}
}

// GET: Queries

// POST: Commands
// @param string EventId
// @param string PromotionType
// @param string Code
// @param int StartTime
// @param int EndTime
// @param float Amount
// @param int Used
// @param int Maximum
// @param string Message
// @param bool|null IsEnabled true|false
// @param string|null PromotionId

type CreatePromotionParameters struct {
	EventId       string
	PromotionType string
	Code          string
	StartTime     int
	EndTime       int
	Amount        float64
	Used          int
	Maximum       int
	Message       string
	IsEnabled     *bool
	PromotionId   *string
}

func (t *Promotion) CreatePromotion(p *CreatePromotionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`promotionType`, p.PromotionType)
	queryParameters.Add(`code`, p.Code)
	queryParameters.Add(`startTime`, strconv.Itoa(p.StartTime))
	queryParameters.Add(`endTime`, strconv.Itoa(p.EndTime))
	queryParameters.Add(`amount`, fmt.Sprintf("%f", p.Amount))
	queryParameters.Add(`used`, strconv.Itoa(p.Used))
	queryParameters.Add(`maximum`, strconv.Itoa(p.Maximum))
	queryParameters.Add(`message`, p.Message)
	if p.IsEnabled != nil {
		queryParameters.Add(`isEnabled`, strconv.FormatBool(*p.IsEnabled))
	}
	if p.PromotionId != nil {
		queryParameters.Add(`promotionId`, *p.PromotionId)
	}

	return t.restClient.Post(
		`/v2/Promotion/UseCase/CreatePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId

type DisablePromotionParameters struct {
	PromotionId string
}

func (t *Promotion) DisablePromotion(p *DisablePromotionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/DisablePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId

type EnablePromotionParameters struct {
	PromotionId string
}

func (t *Promotion) EnablePromotion(p *EnablePromotionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/EnablePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId

type RemovePromotionParameters struct {
	PromotionId string
}

func (t *Promotion) RemovePromotion(p *RemovePromotionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/RemovePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param float Amount

type SetAmountParameters struct {
	PromotionId string
	Amount      float64
}

func (t *Promotion) SetAmount(p *SetAmountParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`amount`, fmt.Sprintf("%f", p.Amount))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetAmount`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param string Code

type SetCodeParameters struct {
	PromotionId string
	Code        string
}

func (t *Promotion) SetCode(p *SetCodeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`code`, p.Code)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetCode`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param int EndTime

type SetEndTimeParameters struct {
	PromotionId string
	EndTime     int
}

func (t *Promotion) SetEndTime(p *SetEndTimeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`endTime`, strconv.Itoa(p.EndTime))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetEndTime`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param int Maximum

type SetMaximumParameters struct {
	PromotionId string
	Maximum     int
}

func (t *Promotion) SetMaximum(p *SetMaximumParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`maximum`, strconv.Itoa(p.Maximum))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetMaximum`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param string Message

type SetMessageParameters struct {
	PromotionId string
	Message     string
}

func (t *Promotion) SetMessage(p *SetMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`message`, p.Message)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param string PromotionType

type SetPromotionTypeParameters struct {
	PromotionId   string
	PromotionType string
}

func (t *Promotion) SetPromotionType(p *SetPromotionTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`promotionType`, p.PromotionType)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetPromotionType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param int StartTime

type SetStartTimeParameters struct {
	PromotionId string
	StartTime   int
}

func (t *Promotion) SetStartTime(p *SetStartTimeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`startTime`, strconv.Itoa(p.StartTime))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetStartTime`,
		&queryParameters,
		nil,
		nil,
	)
}
