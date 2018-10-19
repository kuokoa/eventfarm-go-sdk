package domaintype

type SitePage struct {
}

func NewSitePage() *SitePage {
	return &SitePage{}
}

type SitePageTemplateDifficultyType struct {
	Slug           string
	Name           string
	Description    string
	IsCustom       bool
	IsBeginner     bool
	IsIntermediate bool
	IsAdvanced     bool
	IsExpert       bool
}

func (f *SitePage) ListSitePageTemplateDifficultyTypes() []SitePageTemplateDifficultyType {
	return []SitePageTemplateDifficultyType{
		{
			Slug:           `custom`,
			Name:           `Custom`,
			Description:    ``,
			IsCustom:       true,
			IsBeginner:     false,
			IsIntermediate: false,
			IsAdvanced:     false,
			IsExpert:       false,
		},
		{
			Slug:           `beginner`,
			Name:           `Beginner`,
			Description:    ``,
			IsCustom:       false,
			IsBeginner:     true,
			IsIntermediate: false,
			IsAdvanced:     false,
			IsExpert:       false,
		},
		{
			Slug:           `intermediate`,
			Name:           `Intermediate`,
			Description:    ``,
			IsCustom:       false,
			IsBeginner:     false,
			IsIntermediate: true,
			IsAdvanced:     false,
			IsExpert:       false,
		},
		{
			Slug:           `advanced`,
			Name:           `Advanced`,
			Description:    ``,
			IsCustom:       false,
			IsBeginner:     false,
			IsIntermediate: false,
			IsAdvanced:     true,
			IsExpert:       false,
		},
		{
			Slug:           `expert`,
			Name:           `Expert`,
			Description:    ``,
			IsCustom:       false,
			IsBeginner:     false,
			IsIntermediate: false,
			IsAdvanced:     false,
			IsExpert:       true,
		},
	}
}
