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

type SitePage struct {
	restClient gosdk.RestClientInterface
}

func NewSitePage(restClient gosdk.RestClientInterface) *SitePage {
	return &SitePage{restClient}
}

// GET: Queries
// @param string SitePageId
func (t *SitePage) GetSitePage(SitePageId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, SitePageId)

	return t.restClient.Get(
		`/v2/SitePage/UseCase/GetSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *SitePage) ListSitePagesForEvent(EventId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/SitePage/UseCase/ListSitePagesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param bool|null ShouldIncludeSharedTemplates true|false
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
func (t *SitePage) ListTemplatesForPool(PoolId string, ShouldIncludeSharedTemplates *bool, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	if ShouldIncludeSharedTemplates != nil {
		queryParameters.Add(`shouldIncludeSharedTemplates`, strconv.FormatBool(*ShouldIncludeSharedTemplates))
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/SitePage/UseCase/ListTemplatesForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string EventId
// @param string Title
// @param string Content
// @param int|null DisplayOrder
// @param string|null SitePageId
// @param string|null Styles
// @param string|null Scripts
// @param string|null SourceTemplateId
func (t *SitePage) CreateSitePage(EventId string, Title string, Content string, DisplayOrder *int, SitePageId *string, Styles *string, Scripts *string, SourceTemplateId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`title`, Title)
	queryParameters.Add(`content`, Content)
	if DisplayOrder != nil {
		queryParameters.Add(`displayOrder`, strconv.Itoa(*DisplayOrder))
	}
	if SitePageId != nil {
		queryParameters.Add(`sitePageId`, *SitePageId)
	}
	if Styles != nil {
		queryParameters.Add(`styles`, *Styles)
	}
	if Scripts != nil {
		queryParameters.Add(`scripts`, *Scripts)
	}
	if SourceTemplateId != nil {
		queryParameters.Add(`sourceTemplateId`, *SourceTemplateId)
	}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/CreateSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Name
// @param string Content
// @param string Difficulty custom|beginner|intermediate|advanced|expert
// @param string|null PoolId
// @param string|null Styles
// @param string|null Scripts
// @param string|null Description
// @param string|null TemplateId
func (t *SitePage) CreateTemplate(Name string, Content string, Difficulty string, PoolId *string, Styles *string, Scripts *string, Description *string, TemplateId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, Name)
	queryParameters.Add(`content`, Content)
	queryParameters.Add(`difficulty`, Difficulty)
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}
	if Styles != nil {
		queryParameters.Add(`styles`, *Styles)
	}
	if Scripts != nil {
		queryParameters.Add(`scripts`, *Scripts)
	}
	if Description != nil {
		queryParameters.Add(`description`, *Description)
	}
	if TemplateId != nil {
		queryParameters.Add(`templateId`, *TemplateId)
	}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/CreateTemplate`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SitePageId
// @param string|null NewSitePageId
// @param string|null ToEventId
func (t *SitePage) DuplicateSitePage(SitePageId string, NewSitePageId *string, ToEventId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, SitePageId)
	if NewSitePageId != nil {
		queryParameters.Add(`newSitePageId`, *NewSitePageId)
	}
	if ToEventId != nil {
		queryParameters.Add(`toEventId`, *ToEventId)
	}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/DuplicateSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TemplateId
// @param string|null NewTemplateId
// @param string|null ToPoolId
func (t *SitePage) DuplicateTemplate(TemplateId string, NewTemplateId *string, ToPoolId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`templateId`, TemplateId)
	if NewTemplateId != nil {
		queryParameters.Add(`newTemplateId`, *NewTemplateId)
	}
	if ToPoolId != nil {
		queryParameters.Add(`toPoolId`, *ToPoolId)
	}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/DuplicateTemplate`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *SitePage) GenerateSitePageTemplates() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/GenerateSitePageTemplates`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SitePageId
func (t *SitePage) RemoveSitePage(SitePageId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, SitePageId)

	return t.restClient.Post(
		`/v2/SitePage/UseCase/RemoveSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SitePageTemplateId
func (t *SitePage) RemoveTemplate(SitePageTemplateId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageTemplateId`, SitePageTemplateId)

	return t.restClient.Post(
		`/v2/SitePage/UseCase/RemoveTemplate`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SitePageId
// @param string Content
// @param string|null Styles
// @param string|null Scripts
func (t *SitePage) SetContentForSitePage(SitePageId string, Content string, Styles *string, Scripts *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, SitePageId)
	queryParameters.Add(`content`, Content)
	if Styles != nil {
		queryParameters.Add(`styles`, *Styles)
	}
	if Scripts != nil {
		queryParameters.Add(`scripts`, *Scripts)
	}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/SetContentForSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SitePageId
// @param string DisplayOrder
func (t *SitePage) SetDisplayOrderForSitePage(SitePageId string, DisplayOrder string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, SitePageId)
	queryParameters.Add(`displayOrder`, DisplayOrder)

	return t.restClient.Post(
		`/v2/SitePage/UseCase/SetDisplayOrderForSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SitePageId
// @param string|null Title
// @param string|null Description
// @param string|null Keywords
// @param string|null ImageUrl
// @param string|null Name
func (t *SitePage) SetMetaInfoForSitePage(SitePageId string, Title *string, Description *string, Keywords *string, ImageUrl *string, Name *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, SitePageId)
	if Title != nil {
		queryParameters.Add(`title`, *Title)
	}
	if Description != nil {
		queryParameters.Add(`description`, *Description)
	}
	if Keywords != nil {
		queryParameters.Add(`keywords`, *Keywords)
	}
	if ImageUrl != nil {
		queryParameters.Add(`imageUrl`, *ImageUrl)
	}
	if Name != nil {
		queryParameters.Add(`name`, *Name)
	}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/SetMetaInfoForSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SitePageId
// @param string Title
func (t *SitePage) SetTitleForSitePage(SitePageId string, Title string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, SitePageId)
	queryParameters.Add(`title`, Title)

	return t.restClient.Post(
		`/v2/SitePage/UseCase/SetTitleForSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SitePageId
// @param string EventId
// @param string Title
// @param string Content
// @param int|null DisplayOrder
func (t *SitePage) UpdateSitePage(SitePageId string, EventId string, Title string, Content string, DisplayOrder *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, SitePageId)
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`title`, Title)
	queryParameters.Add(`content`, Content)
	if DisplayOrder != nil {
		queryParameters.Add(`displayOrder`, strconv.Itoa(*DisplayOrder))
	}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/UpdateSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TemplateId
// @param string|null Name
// @param string|null Content
// @param string|null PoolId
// @param string|null Styles
// @param string|null Scripts
func (t *SitePage) UpdateTemplate(TemplateId string, Name *string, Content *string, PoolId *string, Styles *string, Scripts *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`templateId`, TemplateId)
	if Name != nil {
		queryParameters.Add(`name`, *Name)
	}
	if Content != nil {
		queryParameters.Add(`content`, *Content)
	}
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}
	if Styles != nil {
		queryParameters.Add(`styles`, *Styles)
	}
	if Scripts != nil {
		queryParameters.Add(`scripts`, *Scripts)
	}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/UpdateTemplate`,
		&queryParameters,
		nil,
		nil,
	)
}
