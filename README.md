
<h1 align="center">
      	<img height="100" src="https://i.imgur.com/yiv5YMX.png" alt="Event Farm Logo" /> 
      	<br />Official Event Farm Golang SDK
</h1>

### Installation

```bash
go get https://github.com/eventfarm/go-sdk
```
https://i.imgur.com/yiv5YMX.png
### Usage

OAuth is done using the Password Grant. You are going to need a Client Id and Secret that will be provided by Event Farm.

```
package ef_api

import (
	"encoding/json"

	"github.com/eventfarm/go-sdk/api/usecase"
)

func GetEvent(eventId string) (*domain_model.Event, error) {
    apiDomain := os.Getenv(`EF_API_DOMAIN`)
	restClient := rest.NewHttpRestClient(apiDomain)
	restClient.EnableLogging = true

	authDomain := os.Getenv(`EF_AUTH_DOMAIN`)
	accessTokenRestClient := rest.NewHttpRestClient(authDomain)
	accessTokenRestClient.EnableLogging = true

	c := *rest.NewEventFarmRestClient(
		restClient,
		accessTokenRestClient,
		os.Getenv(`EF_CLIENT_ID`),
		os.Getenv(`EF_CLIENT_SECRET`),
		os.Getenv(`EF_USER`),
		os.Getenv(`EF_PASSWORD`),
		nil,
	)

	uc := usecase.NewUseCaseFactory(&c)
    
	resp, err := uc.Event().GetEvent(
		&usecase.GetEventParameters{
			EventId:  eventId,
			WithData: &[]string{"Pool", "Tags", "Stacks", "StacksWithAvailabilityCounts", "TicketTypes", "EventTexts"},
		},
	)
	if err != nil {
		return nil, err
	}
	if err = ValidateResponse(resp); err != nil {
		return nil, err
	}
	d := &struct {
		Data responses.Event `json:"data"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(d)
	return domain_model.AsEvent(d.Data), nil
}
```

For a more detailed usage example, see [test/sdk.go](test/sdk.go)


### Responses

Responses are following the [Json API Spec](http://jsonapi.org/). Please see the [format/spec guide](http://jsonapi.org/format/) for further details.

```json
{
  "data": {
    "type": "Event",
    "id": "7fff5387-0000-45a5-ae48-1d05a7175b93",
    "attributes": {
      "name": "Awesome Test Party",
      "email": "party-organizer@example.com",
      "contactEmail": "MyPartySupport@example.com",
      "guestLimit": -1,
      "type": {
        "slug": "eventfarm",
        "name": "Event Farm",
        "description": null,
        "isEventFarm": true,
        "isCio": false,
        "isListly": false,
        "isDnc": false,
        "isRnc": false,
        "isRslc": false,
        "isSundance": false
      },
      "eventAttributes": {
        "hasDistribute": true,
        "hasDonate": true,
        "hasFee": true,
        "hasEditName": true,
        "hasReveal": true,
        "hasAllowNotes": true,
        "hasDuplicateEmails": true,
        "hasNavigation": true,
        "hasSocialMedia": true,
        "hasSocialMediaBar": true,
        "hasMapLocation": true,
        "hasShowDescription": true,
        "hasIPadPurchase": false,
        "hasSimpleLayout": true,
        "hasLabelPrint": false,
        "hasSkipEventAllocateDisplay": true,
        "hasGeoRestrict": true,
        "hasVisaCheckout": true,
        "hasArchived": false,
        "hasGuestCanChangeResponse": true,
        "hasEfxEnabled": false
      },
      "userFields": {
        "hasAddress": false,
        "hasCompany": false,
        "hasPhone": false,
        "hasTitle": false,
        "hasCountry": false,
        "hasPosition": false
      },
      "startTime": {
        "timestamp": 1643400000,
        "timezone": "America/Los_Angeles",
        "locale": "en_US",
        "date": {
          "short": "1/28/22",
          "medium": "January 28, 2022",
          "long": "Friday, January 28, 2022"
        },
        "time": {
          "short": "12:00 PM",
          "medium": "12:00 PM PST",
          "long": "12:00:00 PM PST"
        },
        "short": "1/28/22, 12:00 PM PST",
        "medium": "January 28, 2022 at 12:00 PM PST",
        "long": "Friday, January 28, 2022 at 12:00 PM PST"
      },
      "endTime": {
        "timestamp": 1643486400,
        "timezone": "America/Los_Angeles",
        "locale": "en_US",
        "date": {
          "short": "1/29/22",
          "medium": "January 29, 2022",
          "long": "Saturday, January 29, 2022"
        },
        "time": {
          "short": "12:00 PM",
          "medium": "12:00 PM PST",
          "long": "12:00:00 PM PST"
        },
        "short": "1/29/22, 12:00 PM PST",
        "medium": "January 29, 2022 at 12:00 PM PST",
        "long": "Saturday, January 29, 2022 at 12:00 PM PST"
      },
      "altKeyword": "mediumeventp1",
      "venue": {
        "name": null,
        "address": null,
        "artistName": null,
        "artistDeposit": 0,
        "checkName": null,
        "flatPercentage": 0,
        "flatFee": 0,
        "memo": null
      },
      "isEFxEnabled": false,
      "defaultSitePage": null,
      "created": {
        "createdAt": 1516128284,
        "ip4Created": "127.0.0.1"
      },
      "modified": {
        "modifiedAt": 1516128284,
        "ip4Modified": null
      },
      "mapSourceType": {
        "slug": "google",
        "name": "Google",
        "description": "Google Maps",
        "isGoogle": true,
        "isYahoo": false,
        "isBing": false
      }
    },
    "meta": {
      "isInitialized": true
    },
    "links": {
      "self": "https://eventfarm.com/api/v2/Event/UseCase/GetEvent?eventId=7fff5387-0000-45a5-ae48-1d05a7175b93"
    },
    "relationships": {
      "pool": {
        "type": "Pool",
        "id": "7fff1f7a-0000-45a5-a6b7-9123da6f4aff",
        "meta": {
          "isInitialized": false
        },
        "links": {
          "self": "https://eventfarm.com/api/v2/Pool/UseCase/GetPool?poolId=7fff1f7a-0000-45a5-a6b7-9123da6f4aff"
        }
      },
      "stacks": [
        {
          "type": "Stack",
          "id": "7fffb628-0000-45a5-ae48-4a001dbd6d01",
          "attributes": {
            "id": "7fffb628-0000-45a5-ae48-4a001dbd6d01",
            "send": "invite,confirm,decline",
            "inviteEmailId": "7fff3d1a-0000-45a5-ae48-4e08d86c122b",
            "confirmEmailId": "7fff3d1a-0000-45a5-ae48-4e0f60015a83",
            "declineEmailId": "7fff3d1a-0000-45a5-ae48-4f03e0c52a9f",
            "serviceFee": 0,
            "price": 0,
            "quantity": 300,
            "maxQuantity": 10,
            "openingAt": 1516041900,
            "closingAt": 1643485500,
            "isTransferable": true,
            "isDeleted": false,
            "deletedAt": null,
            "stackMethodType": {
              "slug": "invite-to-rsvp-fcfs",
              "name": "Invite to RSVP (FCFS)",
              "description": null,
              "isPublicRegistration": false,
              "isPublicPurchase": false,
              "isInviteToRegister": false,
              "isInviteToPurchase": false,
              "isInviteToRSVP": false,
              "isInviteToRegisterFCFS": false,
              "isInviteToPurchaseFCFS": false,
              "isInviteToRSVPFCFS": true,
              "isInvitation": true,
              "isFirstComeFirstServe": true,
              "isAnyInviteToPurchase": false,
              "isAnyInviteToRegister": false,
              "isAnyInviteToRSVP": true
            },
            "created": {
              "createdAt": 1516128329,
              "ip4Created": "172.18.0.1"
            },
            "modified": {
              "modifiedAt": 1516128337,
              "ip4Modified": "172.18.0.1"
            }
          },
          "extraAttributes": {
            "availabilityCounts": {
              "totalQuantity": 300,
              "totalUsed": 0,
              "totalUnused": 300,
              "totalAvailable": 300,
              "percentUsed": 0,
              "percentUnused": 100,
              "totalAffirmed": 0,
              "totalUnaffirmed": 300,
              "percentAffirmed": 0,
              "percentUnaffirmed": 100,
              "total": 0,
              "assigned": 0,
              "purchased": 0,
              "confirmed-by-rsvp": 0,
              "declined-by-rsvp": 0,
              "left-behind": 0,
              "not-yet-purchased": 0,
              "registered": 0,
              "unconfirmed": 0,
              "recycled": 0,
              "not-yet-registered": 0
            }
          },
          "meta": {
            "isInitialized": true
          },
          "links": {
            "self": "https://eventfarm.com/api/v2/Stack/UseCase/GetStack?stackId=7fffb628-0000-45a5-ae48-4a001dbd6d01"
          },
          "relationships": {
            "ticketType": {
              "type": "TicketType",
              "id": "7fff36ce-0000-45a5-ae48-3918d3759d9d",
              "attributes": {
                "id": "7fff36ce-0000-45a5-ae48-3918d3759d9d",
                "name": "VIP",
                "code": "",
                "description": null,
                "quantity": 300,
                "sortOrder": 1,
                "isDeleted": false,
                "deletedAt": null,
                "created": {
                  "createdAt": 1516128313,
                  "ip4Created": "172.18.0.1"
                },
                "modified": {
                  "modifiedAt": 1516128313,
                  "ip4Modified": null
                }
              },
              "meta": {
                "isInitialized": true
              },
              "links": {
                "self": "https://eventfarm.com/api/v2/Stack/UseCase/GetTicketType?ticketTypeId=7fff36ce-0000-45a5-ae48-3918d3759d9d"
              }
            }
          }
        }
      ],
      "ticketBlocks": [
        {
          "type": "TicketBlock",
          "id": "7fff6ce5-0000-45a5-ae49-390fe1d99c88",
          "attributes": {
            "name": "Medium Event P1 - TB2",
            "emailText": null,
            "isBlacklistEnabled": false,
            "isDeleted": false,
            "deletedAt": null,
            "created": {
              "createdAt": 1516128569,
              "ip4Created": "127.0.0.1"
            },
            "modified": {
              "modifiedAt": 1516128569,
              "ip4Modified": null
            }
          },
          "extraAttributes": {
            "allotmentCounts": [
              {
                "allotmentId": "7fffa0bd-0000-45a5-ae49-3918e32c0a59",
                "stackId": "7fffb628-0000-45a5-ae48-4a001dbd6d01",
                "ticketBlockId": "7fff6ce5-0000-45a5-ae49-390fe1d99c88",
                "ticketTypeName": "VIP",
                "stackMethodType": {
                  "slug": "invite-to-rsvp-fcfs",
                  "name": "Invite to RSVP (FCFS)",
                  "description": null,
                  "isPublicRegistration": false,
                  "isPublicPurchase": false,
                  "isInviteToRegister": false,
                  "isInviteToPurchase": false,
                  "isInviteToRSVP": false,
                  "isInviteToRegisterFCFS": false,
                  "isInviteToPurchaseFCFS": false,
                  "isInviteToRSVPFCFS": true,
                  "isInvitation": true,
                  "isFirstComeFirstServe": true,
                  "isAnyInviteToPurchase": false,
                  "isAnyInviteToRegister": false,
                  "isAnyInviteToRSVP": true
                },
                "invitationCounts": {
                  "totalUsed": 0,
                  "firstComeFirstServeOutstanding": 0,
                  "totalQuantity": 10,
                  "totalRemaining": 10
                }
              },
              {
                "allotmentId": "7fffa0bd-0000-45a5-ae49-3920c5604384",
                "stackId": "7fffb628-0000-45a5-ae48-6c08880cbd00",
                "ticketBlockId": "7fff6ce5-0000-45a5-ae49-390fe1d99c88",
                "ticketTypeName": "GA",
                "stackMethodType": {
                  "slug": "invite-to-rsvp-fcfs",
                  "name": "Invite to RSVP (FCFS)",
                  "description": null,
                  "isPublicRegistration": false,
                  "isPublicPurchase": false,
                  "isInviteToRegister": false,
                  "isInviteToPurchase": false,
                  "isInviteToRSVP": false,
                  "isInviteToRegisterFCFS": false,
                  "isInviteToPurchaseFCFS": false,
                  "isInviteToRSVPFCFS": true,
                  "isInvitation": true,
                  "isFirstComeFirstServe": true,
                  "isAnyInviteToPurchase": false,
                  "isAnyInviteToRegister": false,
                  "isAnyInviteToRSVP": true
                },
                "invitationCounts": {
                  "totalUsed": 0,
                  "firstComeFirstServeOutstanding": 0,
                  "totalQuantity": 300,
                  "totalRemaining": 300
                }
              }
            ]
          },
          "meta": {
            "isInitialized": true
          },
          "links": {
            "self": "https://eventfarm.com/api/v2/TicketBlock/UseCase/GetTicketBlock?ticketBlockId=7fff6ce5-0000-45a5-ae49-390fe1d99c88"
          }
        }
      ]
    }
  }
}
```

#### Paginated Response

```json
{
  "meta": {
    "totalPages": 2,
    "totalResults": 2,
    "totalResultsReturned": 1,
    "currentPage": 1,
    "itemsPerPage": 1,
    "executedAt": 1526599136
  },
  "links": {
    "self": "https://eventfarm.com/api/v2/Event/UseCase/ListEventsForUser?eventDateFilterType=current-future\u0026itemsPerPage=1\u0026userId=7fff1483-0000-4578-ad31-f6114a033eb7",
    "first": "https://eventfarm.com/api/v2/Event/UseCase/ListEventsForUser?eventDateFilterType=current-future\u0026itemsPerPage=1\u0026userId=7fff1483-0000-4578-ad31-f6114a033eb7\u0026page=1",
    "last": "https://eventfarm.com/api/v2/Event/UseCase/ListEventsForUser?eventDateFilterType=current-future\u0026itemsPerPage=1\u0026userId=7fff1483-0000-4578-ad31-f6114a033eb7\u0026page=2",
    "next": "https://eventfarm.com/api/v2/Event/UseCase/ListEventsForUser?eventDateFilterType=current-future\u0026itemsPerPage=1\u0026userId=7fff1483-0000-4578-ad31-f6114a033eb7\u0026page=1"
  },
  "data": [
    {
      "type": "Event",
      "id": "7fff5387-0000-4578-ae9c-921df4f75d90",
      "attributes": {
        "name": "Awesome Test Party",
        "email": "party-organizer@example.com",
        "contactEmail": "MyPartySupport@example.com",
        "guestLimit": -1,
        "type": {
          "slug": "eventfarm",
          "name": "Event Farm",
          "description": null,
          "isEventFarm": true,
          "isCio": false,
          "isListly": false,
          "isDnc": false,
          "isRnc": false,
          "isRslc": false,
          "isSundance": false
        },
        "eventAttributes": {
          "hasDistribute": false,
          "hasDonate": false,
          "hasFee": false,
          "hasEditName": false,
          "hasReveal": false,
          "hasAllowNotes": false,
          "hasDuplicateEmails": true,
          "hasNavigation": true,
          "hasSocialMedia": true,
          "hasSocialMediaBar": true,
          "hasMapLocation": false,
          "hasShowDescription": true,
          "hasIPadPurchase": false,
          "hasSimpleLayout": true,
          "hasLabelPrint": false,
          "hasSkipEventAllocateDisplay": true,
          "hasGeoRestrict": false,
          "hasVisaCheckout": false,
          "hasArchived": false,
          "hasGuestCanChangeResponse": true,
          "hasEfxEnabled": false
        },
        "userFields": {
          "hasAddress": false,
          "hasCompany": false,
          "hasPhone": false,
          "hasTitle": false,
          "hasCountry": false,
          "hasPosition": false
        },
        "startTime": {
          "timestamp": 1659043920,
          "timezone": "America/Los_Angeles",
          "locale": "en_US",
          "date": {
            "short": "7/28/22",
            "medium": "July 28, 2022",
            "long": "Thursday, July 28, 2022"
          },
          "time": {
            "short": "2:32 PM",
            "medium": "2:32 PM PST",
            "long": "2:32:00 PM PDT"
          },
          "short": "7/28/22, 2:32 PM PST",
          "medium": "July 28, 2022 at 2:32 PM PST",
          "long": "Thursday, July 28, 2022 at 2:32 PM PST"
        },
        "endTime": {
          "timestamp": 1659130320,
          "timezone": "America/Los_Angeles",
          "locale": "en_US",
          "date": {
            "short": "7/29/22",
            "medium": "July 29, 2022",
            "long": "Friday, July 29, 2022"
          },
          "time": {
            "short": "2:32 PM",
            "medium": "2:32 PM PST",
            "long": "2:32:00 PM PDT"
          },
          "short": "7/29/22, 2:32 PM PST",
          "medium": "July 29, 2022 at 2:32 PM PST",
          "long": "Friday, July 29, 2022 at 2:32 PM PST"
        },
        "altKeyword": "awesome-test-party",
        "venue": {
          "name": "Event Farm HQ",
          "address": "2403 Main Street, Santa Monica, CA",
          "artistName": null,
          "artistDeposit": 0,
          "checkName": null,
          "flatPercentage": 0,
          "flatFee": 0,
          "memo": null
        },
        "isEFxEnabled": false,
        "defaultSitePage": null,
        "created": {
          "createdAt": 1468963985,
          "ip4Created": "172.18.0.1"
        },
        "modified": {
          "modifiedAt": 1516139911,
          "ip4Modified": "172.18.0.1"
        },
        "mapSourceType": {
          "slug": "google",
          "name": "Google",
          "description": "Google Maps",
          "isGoogle": true,
          "isYahoo": false,
          "isBing": false
        }
      },
      "meta": {
        "isInitialized": true
      },
      "links": {
        "self": "https://eventfarm.com/api/v2/Event/UseCase/GetEvent?eventId=7fff5387-0000-4578-ae9c-921df4f75d90"
      },
      "relationships": {
        "pool": {
          "type": "Pool",
          "id": "7fff1f7a-0000-4578-ad49-681e5374350c",
          "meta": {
            "isInitialized": false
          },
          "links": {
            "self": "https://eventfarm.com/api/v2/Pool/UseCase/GetPool?poolId=7fff1f7a-0000-4578-ad49-681e5374350c"
          }
        }
      }
    }
  ]
}
```

