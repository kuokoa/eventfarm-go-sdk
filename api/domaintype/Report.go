package domaintype

type Report struct {
}

func NewReport() *Report {
	return &Report{}
}

type ReportFormatType struct {
	Slug        string
	Name        string
	Description string
	IsCsv       bool
	IsHtml      bool
	IsPdf       bool
	IsZip       bool
}

type ReportType struct {
	Slug         string
	Name         string
	Description  string
	IsInvitation bool
}

func (f *Report) GetReportFormatType() []ReportFormatType {
	return []ReportFormatType{
		{
			Slug:        `csv`,
			Name:        `csv`,
			Description: `csv`,
			IsCsv:       true,
			IsHtml:      false,
			IsPdf:       false,
			IsZip:       false,
		},
		{
			Slug:        `html`,
			Name:        `html`,
			Description: `html`,
			IsCsv:       false,
			IsHtml:      true,
			IsPdf:       false,
			IsZip:       false,
		},
		{
			Slug:        `pdf`,
			Name:        `pdf`,
			Description: `pdf`,
			IsCsv:       false,
			IsHtml:      false,
			IsPdf:       true,
			IsZip:       false,
		},
		{
			Slug:        `zip`,
			Name:        `zip`,
			Description: `zip`,
			IsCsv:       false,
			IsHtml:      false,
			IsPdf:       false,
			IsZip:       true,
		},
	}
}

func (f *Report) GetReportType() []ReportType {
	return []ReportType{
		{
			Slug:         `invitation`,
			Name:         `Invitation`,
			Description:  `Guest Summary Report`,
			IsInvitation: true,
		},
	}
}
