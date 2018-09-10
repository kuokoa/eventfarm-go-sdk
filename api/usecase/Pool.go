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

type Pool struct {
	restClient gosdk.RestClientInterface
}

func NewPool(restClient gosdk.RestClientInterface) *Pool {
	return &Pool{restClient}
}

// GET: Queries
// @param string PoolId
func (t *Pool) GetPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/Pool/UseCase/GetPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
func (t *Pool) GetPoolContract(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/Pool/UseCase/GetPoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
// @param string|null SortBy
// @param string|null SortDirection ascending|descending
func (t *Pool) ListAccessiblePoolsForUser(UserId string, Page *int, ItemsPerPage *int, SortBy *string, SortDirection *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListAccessiblePoolsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string|null SortBy name
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *Pool) ListPoolAllotmentsForPool(PoolId string, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPoolAllotmentsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
func (t *Pool) ListPoolContactsByPoolId(PoolId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPoolContactsByPoolId`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string|null Name
// @param string|null SortBy name
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *Pool) ListPools(Name *string, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	if Name != nil {
		queryParameters.Add(`name`, *Name)
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPools`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string|null SortBy name|slug
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
func (t *Pool) ListTagsForPool(PoolId string, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListTagsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string PoolId
// @param string PoolContractType intro|pro|premier|premierPlus|custom
// @param int StartDate
// @param int EndDate
// @param string|null PoolContractId
func (t *Pool) CreatePoolContract(PoolId string, PoolContractType string, StartDate int, EndDate int, PoolContractId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`poolContractType`, PoolContractType)
	queryParameters.Add(`startDate`, strconv.Itoa(StartDate))
	queryParameters.Add(`endDate`, strconv.Itoa(EndDate))
	if PoolContractId != nil {
		queryParameters.Add(`poolContractId`, *PoolContractId)
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/CreatePoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string PoolWebhookType
// @param string Url
func (t *Pool) CreatePoolWebhook(PoolId string, PoolWebhookType string, Url string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`poolWebhookType`, PoolWebhookType)
	queryParameters.Add(`url`, Url)

	return t.restClient.Post(
		`/v2/Pool/UseCase/CreatePoolWebhook`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
func (t *Pool) RemoveCustomerDisplayNameForPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Post(
		`/v2/Pool/UseCase/RemoveCustomerDisplayNameForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
func (t *Pool) RemovePrivacyPolicyLinkForPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Post(
		`/v2/Pool/UseCase/RemovePrivacyPolicyLinkForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string UserId
// @param string|null SlackUserId
// @param array|null RequestedFeatureSlugs
// @param string|null Other
func (t *Pool) SendUpgradeRequestToCsm(PoolId string, UserId string, SlackUserId *string, RequestedFeatureSlugs *[]string, Other *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	if SlackUserId != nil {
		queryParameters.Add(`slackUserId`, *SlackUserId)
	}
	if RequestedFeatureSlugs != nil {
		for i := range *RequestedFeatureSlugs {
			queryParameters.Add(`requestedFeatureSlugs[]`, (*RequestedFeatureSlugs)[i])
		}
	}
	if Other != nil {
		queryParameters.Add(`other`, *Other)
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/SendUpgradeRequestToCsm`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string CustomerDisplayName
func (t *Pool) SetCustomerDisplayNameForPool(PoolId string, CustomerDisplayName string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`customerDisplayName`, CustomerDisplayName)

	return t.restClient.Post(
		`/v2/Pool/UseCase/SetCustomerDisplayNameForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string PrivacyPolicyLink
func (t *Pool) SetPrivacyPolicyLinkForPool(PoolId string, PrivacyPolicyLink string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`privacyPolicyLink`, PrivacyPolicyLink)

	return t.restClient.Post(
		`/v2/Pool/UseCase/SetPrivacyPolicyLinkForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolContractId
// @param string|null PoolContractType intro|pro|premier|premierPlus|custom
// @param int|null StartDate
// @param int|null EndDate
func (t *Pool) UpdatePoolContract(PoolContractId string, PoolContractType *string, StartDate *int, EndDate *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolContractId`, PoolContractId)
	if PoolContractType != nil {
		queryParameters.Add(`poolContractType`, *PoolContractType)
	}
	if StartDate != nil {
		queryParameters.Add(`startDate`, strconv.Itoa(*StartDate))
	}
	if EndDate != nil {
		queryParameters.Add(`endDate`, strconv.Itoa(*EndDate))
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/UpdatePoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}
