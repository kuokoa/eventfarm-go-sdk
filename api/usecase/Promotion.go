package usecase

import (
	"gosdk"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type Promotion struct {
	restClient gosdk.RestClientInterface
}

func NewPromotion(restClient gosdk.RestClientInterface) *Promotion {
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
func (t *Promotion) CreatePromotion(EventId string, PromotionType string, Code string, StartTime int, EndTime int, Amount float, Used int, Maximum int, Message string, IsEnabled *bool, PromotionId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`promotionType`, PromotionType)
	queryParameters.Add(`code`, Code)
	queryParameters.Add(`startTime`, strconv.Itoa(StartTime))
	queryParameters.Add(`endTime`, strconv.Itoa(EndTime))
	queryParameters.Add(`amount`, Amount)
	queryParameters.Add(`used`, strconv.Itoa(Used))
	queryParameters.Add(`maximum`, strconv.Itoa(Maximum))
	queryParameters.Add(`message`, Message)
	if IsEnabled != nil {
		queryParameters.Add(`isEnabled`, strconv.FormatBool(*IsEnabled))
	}
	if PromotionId != nil {
		queryParameters.Add(`promotionId`, *PromotionId)
	}

	return t.restClient.Post(
		`/v2/Promotion/UseCase/CreatePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
func (t *Promotion) DisablePromotion(PromotionId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/DisablePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
func (t *Promotion) EnablePromotion(PromotionId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/EnablePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
func (t *Promotion) RemovePromotion(PromotionId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/RemovePromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param float Amount
func (t *Promotion) SetAmount(PromotionId string, Amount float) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	queryParameters.Add(`amount`, Amount)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetAmount`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param string Code
func (t *Promotion) SetCode(PromotionId string, Code string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	queryParameters.Add(`code`, Code)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetCode`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param int EndTime
func (t *Promotion) SetEndTime(PromotionId string, EndTime int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	queryParameters.Add(`endTime`, strconv.Itoa(EndTime))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetEndTime`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param int Maximum
func (t *Promotion) SetMaximum(PromotionId string, Maximum int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	queryParameters.Add(`maximum`, strconv.Itoa(Maximum))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetMaximum`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param string Message
func (t *Promotion) SetMessage(PromotionId string, Message string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	queryParameters.Add(`message`, Message)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param string PromotionType
func (t *Promotion) SetPromotionType(PromotionId string, PromotionType string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	queryParameters.Add(`promotionType`, PromotionType)

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetPromotionType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param int StartTime
func (t *Promotion) SetStartTime(PromotionId string, StartTime int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	queryParameters.Add(`startTime`, strconv.Itoa(StartTime))

	return t.restClient.Post(
		`/v2/Promotion/UseCase/SetStartTime`,
		&queryParameters,
		nil,
		nil,
	)
}
