/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type EmailDesign struct {
}

func NewEmailDesign() *EmailDesign {
	return &EmailDesign{}
}

type EmailDesignLayoutType struct {
	Slug             string
	Name             string
	Description      string
	IsBlank          bool
	IsAltEmailLayout bool
}

func (f *EmailDesign) ListEmailDesignLayoutTypes() []EmailDesignLayoutType {
	return []EmailDesignLayoutType{
		{
			Slug:             `blank`,
			Name:             `Blank`,
			Description:      ``,
			IsBlank:          true,
			IsAltEmailLayout: false,
		},
		{
			Slug:             `alt-email-layout`,
			Name:             `Alt Email Layout`,
			Description:      ``,
			IsBlank:          false,
			IsAltEmailLayout: true,
		},
	}
}
