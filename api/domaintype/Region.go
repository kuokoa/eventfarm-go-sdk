package domaintype

type Region struct {
}

func NewRegion() *Region {
	return &Region{}
}

type RegionType struct {
	Slug           string
	Name           string
	Description    string
	IsAfrica       bool
	IsAntarctica   bool
	IsAsia         bool
	IsEurope       bool
	IsNorthAmerica bool
	IsOceania      bool
	IsSouthAmerica bool
}

func (f *Region) GetRegionType() []RegionType {
	return []RegionType{
		{
			Slug:           `africa`,
			Name:           `Africa`,
			Description:    ``,
			IsAfrica:       true,
			IsAntarctica:   false,
			IsAsia:         false,
			IsEurope:       false,
			IsNorthAmerica: false,
			IsOceania:      false,
			IsSouthAmerica: false,
		},
		{
			Slug:           `antarctica`,
			Name:           `Antarctica`,
			Description:    ``,
			IsAfrica:       false,
			IsAntarctica:   true,
			IsAsia:         false,
			IsEurope:       false,
			IsNorthAmerica: false,
			IsOceania:      false,
			IsSouthAmerica: false,
		},
		{
			Slug:           `asia`,
			Name:           `Asia`,
			Description:    ``,
			IsAfrica:       false,
			IsAntarctica:   false,
			IsAsia:         true,
			IsEurope:       false,
			IsNorthAmerica: false,
			IsOceania:      false,
			IsSouthAmerica: false,
		},
		{
			Slug:           `europe`,
			Name:           `Europe`,
			Description:    ``,
			IsAfrica:       false,
			IsAntarctica:   false,
			IsAsia:         false,
			IsEurope:       true,
			IsNorthAmerica: false,
			IsOceania:      false,
			IsSouthAmerica: false,
		},
		{
			Slug:           `north-america`,
			Name:           `North America`,
			Description:    ``,
			IsAfrica:       false,
			IsAntarctica:   false,
			IsAsia:         false,
			IsEurope:       false,
			IsNorthAmerica: true,
			IsOceania:      false,
			IsSouthAmerica: false,
		},
		{
			Slug:           `oceania`,
			Name:           `Oceania`,
			Description:    ``,
			IsAfrica:       false,
			IsAntarctica:   false,
			IsAsia:         false,
			IsEurope:       false,
			IsNorthAmerica: false,
			IsOceania:      true,
			IsSouthAmerica: false,
		},
		{
			Slug:           `south-america`,
			Name:           `South America`,
			Description:    ``,
			IsAfrica:       false,
			IsAntarctica:   false,
			IsAsia:         false,
			IsEurope:       false,
			IsNorthAmerica: false,
			IsOceania:      false,
			IsSouthAmerica: true,
		},
	}
}
