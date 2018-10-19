package domaintype

type AppVersion struct {
}

func NewAppVersion() *AppVersion {
	return &AppVersion{}
}

type AppVersionType struct {
	Slug                           string
	Name                           string
	Description                    string
	IsCheckInIos                   bool
	IsCheckInAndroid               bool
	IsTicketBlockManagementIos     bool
	IsTicketBlockManagementAndroid bool
}

func (f *AppVersion) ListAppVersionTypes() []AppVersionType {
	return []AppVersionType{
		{
			Slug:                           `check-in-ios`,
			Name:                           `Check In iOS`,
			Description:                    ``,
			IsCheckInIos:                   true,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: false,
		},
		{
			Slug:                           `check-in-android`,
			Name:                           `Check In Android`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               true,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: false,
		},
		{
			Slug:                           `ticket-block-mgmt-ios`,
			Name:                           `Ticket Block Management iOS`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     true,
			IsTicketBlockManagementAndroid: false,
		},
		{
			Slug:                           `ticket-block-mgmt-android`,
			Name:                           `Ticket Block Management Android`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: true,
		},
	}
}
