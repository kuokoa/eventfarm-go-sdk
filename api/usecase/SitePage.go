/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/eventfarm/go-sdk/rest"
)

type SitePage struct {
	restClient rest.RestClientInterface
}

func NewSitePage(restClient rest.RestClientInterface) *SitePage {
	return &SitePage{restClient}
}

// GET: Queries
// @param string SitePageId

type GetSitePageParameters struct {
	SitePageId string
}

func (t *SitePage) GetSitePage(p *GetSitePageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, p.SitePageId)

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

type ListSitePagesForEventParameters struct {
	EventId      string
	Page         *int64
	ItemsPerPage *int64
}

func (t *SitePage) ListSitePagesForEvent(p *ListSitePagesForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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

type ListTemplatesForPoolParameters struct {
	PoolId                       string
	ShouldIncludeSharedTemplates *bool
	Page                         *int64
	ItemsPerPage                 *int64
}

func (t *SitePage) ListTemplatesForPool(p *ListTemplatesForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.ShouldIncludeSharedTemplates != nil {
		queryParameters.Add(`shouldIncludeSharedTemplates`, strconv.FormatBool(*p.ShouldIncludeSharedTemplates))
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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

type CreateSitePageParameters struct {
	EventId          string
	Title            string
	Content          string
	DisplayOrder     *int64
	SitePageId       *string
	Styles           *string
	Scripts          *string
	SourceTemplateId *string
}

func (t *SitePage) CreateSitePage(p *CreateSitePageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`title`, p.Title)
	queryParameters.Add(`content`, p.Content)
	if p.DisplayOrder != nil {
		queryParameters.Add(`displayOrder`, strconv.FormatInt(*p.DisplayOrder, 10))
	}
	if p.SitePageId != nil {
		queryParameters.Add(`sitePageId`, *p.SitePageId)
	}
	if p.Styles != nil {
		queryParameters.Add(`styles`, *p.Styles)
	}
	if p.Scripts != nil {
		queryParameters.Add(`scripts`, *p.Scripts)
	}
	if p.SourceTemplateId != nil {
		queryParameters.Add(`sourceTemplateId`, *p.SourceTemplateId)
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

type CreateTemplateParameters struct {
	Name        string
	Content     string
	Difficulty  string
	PoolId      *string
	Styles      *string
	Scripts     *string
	Description *string
	TemplateId  *string
}

func (t *SitePage) CreateTemplate(p *CreateTemplateParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`content`, p.Content)
	queryParameters.Add(`difficulty`, p.Difficulty)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.Styles != nil {
		queryParameters.Add(`styles`, *p.Styles)
	}
	if p.Scripts != nil {
		queryParameters.Add(`scripts`, *p.Scripts)
	}
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}
	if p.TemplateId != nil {
		queryParameters.Add(`templateId`, *p.TemplateId)
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

type DuplicateSitePageParameters struct {
	SitePageId    string
	NewSitePageId *string
	ToEventId     *string
}

func (t *SitePage) DuplicateSitePage(p *DuplicateSitePageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, p.SitePageId)
	if p.NewSitePageId != nil {
		queryParameters.Add(`newSitePageId`, *p.NewSitePageId)
	}
	if p.ToEventId != nil {
		queryParameters.Add(`toEventId`, *p.ToEventId)
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

type DuplicateTemplateParameters struct {
	TemplateId    string
	NewTemplateId *string
	ToPoolId      *string
}

func (t *SitePage) DuplicateTemplate(p *DuplicateTemplateParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`templateId`, p.TemplateId)
	if p.NewTemplateId != nil {
		queryParameters.Add(`newTemplateId`, *p.NewTemplateId)
	}
	if p.ToPoolId != nil {
		queryParameters.Add(`toPoolId`, *p.ToPoolId)
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

type RemoveSitePageParameters struct {
	SitePageId string
}

func (t *SitePage) RemoveSitePage(p *RemoveSitePageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, p.SitePageId)

	return t.restClient.Post(
		`/v2/SitePage/UseCase/RemoveSitePage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SitePageTemplateId

type RemoveTemplateParameters struct {
	SitePageTemplateId string
}

func (t *SitePage) RemoveTemplate(p *RemoveTemplateParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageTemplateId`, p.SitePageTemplateId)

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

type SetContentForSitePageParameters struct {
	SitePageId string
	Content    string
	Styles     *string
	Scripts    *string
}

func (t *SitePage) SetContentForSitePage(p *SetContentForSitePageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, p.SitePageId)
	queryParameters.Add(`content`, p.Content)
	if p.Styles != nil {
		queryParameters.Add(`styles`, *p.Styles)
	}
	if p.Scripts != nil {
		queryParameters.Add(`scripts`, *p.Scripts)
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

type SetDisplayOrderForSitePageParameters struct {
	SitePageId   string
	DisplayOrder string
}

func (t *SitePage) SetDisplayOrderForSitePage(p *SetDisplayOrderForSitePageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, p.SitePageId)
	queryParameters.Add(`displayOrder`, p.DisplayOrder)

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

type SetMetaInfoForSitePageParameters struct {
	SitePageId  string
	Title       *string
	Description *string
	Keywords    *string
	ImageUrl    *string
	Name        *string
}

func (t *SitePage) SetMetaInfoForSitePage(p *SetMetaInfoForSitePageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, p.SitePageId)
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}
	if p.Keywords != nil {
		queryParameters.Add(`keywords`, *p.Keywords)
	}
	if p.ImageUrl != nil {
		queryParameters.Add(`imageUrl`, *p.ImageUrl)
	}
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
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

type SetTitleForSitePageParameters struct {
	SitePageId string
	Title      string
}

func (t *SitePage) SetTitleForSitePage(p *SetTitleForSitePageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, p.SitePageId)
	queryParameters.Add(`title`, p.Title)

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

type UpdateSitePageParameters struct {
	SitePageId   string
	EventId      string
	Title        string
	Content      string
	DisplayOrder *int64
}

func (t *SitePage) UpdateSitePage(p *UpdateSitePageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`sitePageId`, p.SitePageId)
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`title`, p.Title)
	queryParameters.Add(`content`, p.Content)
	if p.DisplayOrder != nil {
		queryParameters.Add(`displayOrder`, strconv.FormatInt(*p.DisplayOrder, 10))
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

type UpdateTemplateParameters struct {
	TemplateId string
	Name       *string
	Content    *string
	PoolId     *string
	Styles     *string
	Scripts    *string
}

func (t *SitePage) UpdateTemplate(p *UpdateTemplateParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`templateId`, p.TemplateId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.Content != nil {
		queryParameters.Add(`content`, *p.Content)
	}
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.Styles != nil {
		queryParameters.Add(`styles`, *p.Styles)
	}
	if p.Scripts != nil {
		queryParameters.Add(`scripts`, *p.Scripts)
	}

	return t.restClient.Post(
		`/v2/SitePage/UseCase/UpdateTemplate`,
		&queryParameters,
		nil,
		nil,
	)
}
