package elements

// The GetUserAvailabilityRequest element contains the arguments used to obtain user availability information. This is a root element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseravailabilityrequest
import "encoding/xml"

type GetUserAvailabilityRequest struct {
	XMLName xml.Name

	// The FreeBusyViewOptions element specifies the type of free/busy information returned in the response.
	FreeBusyViewOptions *FreeBusyViewOptions `xml:"FreeBusyViewOptions"`
	// The MailboxDataArray element contains a list of mailboxes to query for availability information.
	MailboxDataArray *MailboxDataArray `xml:"MailboxDataArray"`
	// The SuggestionsViewOptions element contains the options for obtaining meeting suggestion information.
	SuggestionsViewOptions *SuggestionsViewOptions `xml:"SuggestionsViewOptions"`
	// The TimeZone element contains elements that identify time zone information. This element also contains information about the transition between standard time and daylight saving time.
	TimeZone *TimeZoneAvailability `xml:"TimeZone"`
}

func (G *GetUserAvailabilityRequest) SetForMarshal() {
	G.XMLName.Local = "m:GetUserAvailabilityRequest"
}

func (G *GetUserAvailabilityRequest) GetSchema() *Schema {
	return &SchemaMessages
}
