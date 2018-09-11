package main

import (
	"log"
	"os"

	"bitbucket.ef.network/go/sdk"
	"bitbucket.ef.network/go/sdk/api/domaintype"
	"bitbucket.ef.network/go/sdk/api/usecase"
	"encoding/json"
	"io/ioutil"
)

type EventAttributes struct {
	HasDistribute               bool `json:"hasDistribute"`
	HasDonate                   bool `json:"hasDonate"`
	HasFee                      bool `json:"hasFee"`
	HasEditName                 bool `json:"hasEditName"`
	HasReveal                   bool `json:"hasReveal"`
	HasAllowNotes               bool `json:"hasAllowNotes"`
	HasDuplicateEmails          bool `json:"hasDuplicateEmails"`
	HasNavigation               bool `json:"hasNavigation"`
	HasSocialMedia              bool `json:"hasSocialMedia"`
	HasSocialMediaBar           bool `json:"hasSocialMediaBar "`
	HasMapLocation              bool `json:"hasMapLocation"`
	HasShowDescription          bool `json:"hasShowDescription"`
	HasIPadPurchase             bool `json:"hasIPadPurchase"`
	HasSimpleLayout             bool `json:"hasSimpleLayout"`
	HasLabelPrint               bool `json:"hasLabelPrint"`
	HasFacebookConnect          bool `json:"hasFacebookConnect"`
	HasSkipEventAllocateDisplay bool `json:"hasSkipEventAllocateDisplay"`
	HasGeoRestrict              bool `json:"hasGeoRestrict"`
	HasVisaCheckout             bool `json:"hasVisaCheckout"`
	HasArchived                 bool `json:"hasArchived"`
	HasGuestCanChangeResponse   bool `json:"hasGuestCanChangeResponse"`
}

type Event struct {
	Name            string          `json:"name"`
	Email           string          `json:"email"`
	GuestLimit      int             `json:"guestLimit"`
	Type            string          `json:"domaintype"`
	EventAttributes EventAttributes `json:"eventAttributes"`
	IsEFxEnabled    bool            `json:"isEFxEnabled"`
}

type PaginationMetaData struct {
	TotalPages           int `json:"totalPages"`
	TotalResults         int `json:"totalResults"`
	TotalResultsReturned int `json:"totalResultsReturned"`
	CurrentPage          int `json:"currentpage"`
	ItemsPerPage         int `json:"itemsPerPage"`
	ExecutedAt           int `json:"executedAt"`
}

type PaginationLinks struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
	Last  string `json:"last"`
}

type EventData struct {
	Type       string `json:"domaintype"`
	Id         string `json:"id"`
	Attributes Event  `json:"attributes"`
}

type EventsResponse struct {
	Meta  PaginationMetaData `json:"meta"`
	Links PaginationLinks    `json:"links"`
	Data  []EventData        `json:"data"`
}

func main() {
	testUseCase()
	testType()
}

func testType() {
	typeFactory := domaintype.NewDomainTypeFactory()

	log.Printf("%+v", typeFactory.GetInvitationDomainModel().GetInvitationStatusType()[0])
}

func testUseCase() {
	c := sdk.NewEventFarmRestClientForStage(
		os.Getenv(`EF_STAGE`),
		os.Getenv(`EF_CLIENT_ID`),
		os.Getenv(`EF_CLIENT_SECRET`),
	)
	uc := usecase.NewUseCaseFactory(c)

	page := 1
	itemsPerPage := 20
	sortBy := `event-start`
	sortDirection := `descending`
	var withData []string
	withData = []string{
		`Pool`,
		`Tags`,
		`Stacks`,
		`TicketTypes`,
		`TicketBlocks`,
		`QuestionsAndAnswers`,
	}
	resp, err := uc.Event().ListEventsForUser(&usecase.ListEventsForUserParameters{
		UserId: `7fff1483-0000-4578-ad31-f6114a033eb7`,
		WithData: &withData,
		Page: &page,
		ItemsPerPage: &itemsPerPage,
		SortBy: &sortBy,
		SortDirection: &sortDirection,
	})
	if err != nil {
		logError(err)
		return
	}

	bodyBuff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logError(err)
		return
	}

	eventsResponse := EventsResponse{}
	err = json.Unmarshal(bodyBuff, &eventsResponse)
	if err != nil {
		logError(err)
		return
	}

	log.Printf("%+v", eventsResponse)
}

func logError(err error) {
	log.Print(err)
}
