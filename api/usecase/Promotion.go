/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"github.com/eventfarm/go-sdk/rest"
	"net/http"
	"net/url"
	"strconv"
)

type Promotion struct {
	restClient rest.RestClientInterface
}

func NewPromotion(restClient rest.RestClientInterface) *Promotion {
	return &Promotion{restClient}
}

// GET: Queries

type ListPromotionsForEventParameters struct {
	EventId       string
	WithData      *[]interface{} // StackAndTicketType
	SortBy        *string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
}

func (t *Promotion) ListPromotionsForEvent(p *ListPromotionsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Promotion/UseCase/ListPromotionsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePromotionParameters struct {
	EventId       string
	PromotionType string
	Code          string
	StartTime     string
	EndTime       string
	Amount        float64
	Used          int64
	Message       string
	Maximum       *int64
	IsEnabled     *bool
	PromotionId   *string
	IsDeleted     *bool
}

func (t *Promotion) CreatePromotion(p *CreatePromotionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`promotionType`, p.PromotionType)
	queryParameters.Add(`code`, p.Code)
	queryParameters.Add(`startTime`, p.StartTime)
	queryParameters.Add(`endTime`, p.EndTime)
	queryParameters.Add(`amount`, fmt.Sprintf("%f", p.Amount))
	queryParameters.Add(`used`, strconv.FormatInt(p.Used, 10))
	queryParameters.Add(`message`, p.Message)
	if p.Maximum != nil {
		queryParameters.Add(`maximum`, strconv.FormatInt(*p.Maximum, 10))
	}
	if p.IsEnabled != nil {
		queryParameters.Add(`isEnabled`, strconv.FormatBool(*p.IsEnabled))
	}
	if p.PromotionId != nil {
		queryParameters.Add(`promotionId`, *p.PromotionId)
	}
	if p.IsDeleted != nil {
		queryParameters.Add(`isDeleted`, strconv.FormatBool(*p.IsDeleted))
	}

	return t.restClient.Post(
		`/v2/Promotion/UseCase/CreatePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

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

type SetEndTimeParameters struct {
	PromotionId string
	EndTime     int64
}

func (t *Promotion) SetEndTime(p *SetEndTimeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`endTime`, strconv.FormatInt(p.EndTime, 10))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetEndTime`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetMaximumParameters struct {
	PromotionId string
	Maximum     int64
}

func (t *Promotion) SetMaximum(p *SetMaximumParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`maximum`, strconv.FormatInt(p.Maximum, 10))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetMaximum`,
		&queryParameters,
		nil,
		nil,
	)
}

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

type SetStacksForPromotionParameters struct {
	PromotionId string
	StackIds    []interface{}
}

func (t *Promotion) SetStacksForPromotion(p *SetStacksForPromotionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	for i := range p.StackIds {
		queryParameters.Add(`stackIds[]`, p.StackIds[i])
	}

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetStacksForPromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetStartTimeParameters struct {
	PromotionId string
	StartTime   int64
}

func (t *Promotion) SetStartTime(p *SetStartTimeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	queryParameters.Add(`startTime`, strconv.FormatInt(p.StartTime, 10))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetStartTime`,
		&queryParameters,
		nil,
		nil,
	)
}

type UpdatePromotionParameters struct {
	PromotionId   string
	PromotionType *string
	Code          *string
	StartTime     *string
	EndTime       *string
	Amount        *float64
	Used          *int64
	Message       *string
	Maximum       *int64
	IsEnabled     *bool
}

func (t *Promotion) UpdatePromotion(p *UpdatePromotionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	if p.PromotionType != nil {
		queryParameters.Add(`promotionType`, *p.PromotionType)
	}
	if p.Code != nil {
		queryParameters.Add(`code`, *p.Code)
	}
	if p.StartTime != nil {
		queryParameters.Add(`startTime`, *p.StartTime)
	}
	if p.EndTime != nil {
		queryParameters.Add(`endTime`, *p.EndTime)
	}
	if p.Amount != nil {
		queryParameters.Add(`amount`, fmt.Sprintf("%f", *p.Amount))
	}
	if p.Used != nil {
		queryParameters.Add(`used`, strconv.FormatInt(*p.Used, 10))
	}
	if p.Message != nil {
		queryParameters.Add(`message`, *p.Message)
	}
	if p.Maximum != nil {
		queryParameters.Add(`maximum`, strconv.FormatInt(*p.Maximum, 10))
	}
	if p.IsEnabled != nil {
		queryParameters.Add(`isEnabled`, strconv.FormatBool(*p.IsEnabled))
	}

	return t.restClient.Post(
		`/v2/Promotion/UseCase/UpdatePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}
