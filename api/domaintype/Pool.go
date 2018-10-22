/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Pool struct {
}

func NewPool() *Pool {
	return &Pool{}
}

type FeatureType struct {
	Slug            string
	Name            string
	Description     string
	IsSalesforce    bool
	IsApiAccess     bool
	IsMarketo       bool
	IsEmailMasking  bool
	IsWhiteLabeling bool
	IsAgents        bool
	IsVisaCheckout  bool
	IsCanvas        bool
}

type PoolAllotmentType struct {
	Slug        string
	Name        string
	Description string
	IsListly    bool
	IsCIO       bool
	IsEventFarm bool
}

type PoolContactType struct {
	Slug        string
	Name        string
	Description string
	IsFull      bool
	IsCreate    bool
}

type PoolContractType struct {
	Slug          string
	Name          string
	Description   string
	IsIntro       bool
	IsPro         bool
	IsPremier     bool
	IsPremierPlus bool
	IsCustom      bool
}

type PoolWebhookType struct {
	Slug        string
	Name        string
	Description string
	IsCheckin   bool
}

func (f *Pool) ListFeatureTypes() []FeatureType {
	return []FeatureType{
		{
			Slug:            `salesforce`,
			Name:            `Salesforce`,
			Description:     ``,
			IsSalesforce:    true,
			IsApiAccess:     false,
			IsMarketo:       false,
			IsEmailMasking:  false,
			IsWhiteLabeling: false,
			IsAgents:        false,
			IsVisaCheckout:  false,
			IsCanvas:        false,
		},
		{
			Slug:            `apiaccess`,
			Name:            `API Access`,
			Description:     ``,
			IsSalesforce:    false,
			IsApiAccess:     true,
			IsMarketo:       false,
			IsEmailMasking:  false,
			IsWhiteLabeling: false,
			IsAgents:        false,
			IsVisaCheckout:  false,
			IsCanvas:        false,
		},
		{
			Slug:            `marketo`,
			Name:            `Marketo`,
			Description:     ``,
			IsSalesforce:    false,
			IsApiAccess:     false,
			IsMarketo:       true,
			IsEmailMasking:  false,
			IsWhiteLabeling: false,
			IsAgents:        false,
			IsVisaCheckout:  false,
			IsCanvas:        false,
		},
		{
			Slug:            `emailmasking`,
			Name:            `Email Masking`,
			Description:     ``,
			IsSalesforce:    false,
			IsApiAccess:     false,
			IsMarketo:       false,
			IsEmailMasking:  true,
			IsWhiteLabeling: false,
			IsAgents:        false,
			IsVisaCheckout:  false,
			IsCanvas:        false,
		},
		{
			Slug:            `whitelabeling`,
			Name:            `White Labeling`,
			Description:     ``,
			IsSalesforce:    false,
			IsApiAccess:     false,
			IsMarketo:       false,
			IsEmailMasking:  false,
			IsWhiteLabeling: true,
			IsAgents:        false,
			IsVisaCheckout:  false,
			IsCanvas:        false,
		},
		{
			Slug:            `agents`,
			Name:            `Agents`,
			Description:     ``,
			IsSalesforce:    false,
			IsApiAccess:     false,
			IsMarketo:       false,
			IsEmailMasking:  false,
			IsWhiteLabeling: false,
			IsAgents:        true,
			IsVisaCheckout:  false,
			IsCanvas:        false,
		},
		{
			Slug:            `visacheckout`,
			Name:            `VisaCheckout`,
			Description:     ``,
			IsSalesforce:    false,
			IsApiAccess:     false,
			IsMarketo:       false,
			IsEmailMasking:  false,
			IsWhiteLabeling: false,
			IsAgents:        false,
			IsVisaCheckout:  true,
			IsCanvas:        false,
		},
		{
			Slug:            `canvas`,
			Name:            `Canvas`,
			Description:     ``,
			IsSalesforce:    false,
			IsApiAccess:     false,
			IsMarketo:       false,
			IsEmailMasking:  false,
			IsWhiteLabeling: false,
			IsAgents:        false,
			IsVisaCheckout:  false,
			IsCanvas:        true,
		},
	}
}

func (f *Pool) ListPoolAllotmentTypes() []PoolAllotmentType {
	return []PoolAllotmentType{
		{
			Slug:        `listly`,
			Name:        `Listly`,
			Description: ``,
			IsListly:    true,
			IsCIO:       false,
			IsEventFarm: false,
		},
		{
			Slug:        `cio`,
			Name:        `CIO`,
			Description: ``,
			IsListly:    false,
			IsCIO:       true,
			IsEventFarm: false,
		},
		{
			Slug:        `eventfarm`,
			Name:        `EventFarm`,
			Description: ``,
			IsListly:    false,
			IsCIO:       false,
			IsEventFarm: true,
		},
	}
}

func (f *Pool) ListPoolContactTypes() []PoolContactType {
	return []PoolContactType{
		{
			Slug:        `full`,
			Name:        `Team Manager`,
			Description: ``,
			IsFull:      true,
			IsCreate:    false,
		},
		{
			Slug:        `create`,
			Name:        `Event Create Only`,
			Description: ``,
			IsFull:      false,
			IsCreate:    true,
		},
	}
}

func (f *Pool) ListPoolContractTypes() []PoolContractType {
	return []PoolContractType{
		{
			Slug:          `intro`,
			Name:          `Intro`,
			Description:   ``,
			IsIntro:       true,
			IsPro:         false,
			IsPremier:     false,
			IsPremierPlus: false,
			IsCustom:      false,
		},
		{
			Slug:          `pro`,
			Name:          `Pro`,
			Description:   ``,
			IsIntro:       false,
			IsPro:         true,
			IsPremier:     false,
			IsPremierPlus: false,
			IsCustom:      false,
		},
		{
			Slug:          `premier`,
			Name:          `Premier`,
			Description:   ``,
			IsIntro:       false,
			IsPro:         false,
			IsPremier:     true,
			IsPremierPlus: false,
			IsCustom:      false,
		},
		{
			Slug:          `premierPlus`,
			Name:          `Premier Plus`,
			Description:   ``,
			IsIntro:       false,
			IsPro:         false,
			IsPremier:     false,
			IsPremierPlus: true,
			IsCustom:      false,
		},
		{
			Slug:          `custom`,
			Name:          `Custom`,
			Description:   ``,
			IsIntro:       false,
			IsPro:         false,
			IsPremier:     false,
			IsPremierPlus: false,
			IsCustom:      true,
		},
	}
}

func (f *Pool) ListPoolWebhookTypes() []PoolWebhookType {
	return []PoolWebhookType{
		{
			Slug:        `checkin`,
			Name:        `Check-In`,
			Description: ``,
			IsCheckin:   true,
		},
	}
}
