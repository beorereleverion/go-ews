package elements

// The MessageTrackingSearchResult element contains a single message result for a FindMessageTrackingReportResponse element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagetrackingsearchresult
import "encoding/xml"

type MessageTrackingSearchResult struct {
	XMLName xml.Name

	// The FirstHopServer element contains the name of the server in the forest that first accepted the message.
	FirstHopServer *FirstHopServer `xml:"FirstHopServer"`
	// The MessageTrackingReportId element represents the message by its message ID, the organization where the message was found, the server on which the message was submitted, and an internal ID that uniquely identifies the message.
	MessageTrackingReportId *MessageTrackingReportId `xml:"MessageTrackingReportId"`
	// The PreviousHopServer element represents the previous server name that accepted the message.
	PreviousHopServer *PreviousHopServer `xml:"PreviousHopServer"`
	// The Properties element contains a list of one or more tracking properties.
	Properties *PropertiesArrayOfTrackingPropertiesType `xml:"Properties"`
	// The PurportedSender element contains contact information for the alleged sender of an e-mail message.
	PurportedSender *PurportedSender `xml:"PurportedSender"`
	// The Recipients element represents a collection of recipients that receive a copy of the message.
	Recipients *RecipientsArrayOfRecipientsType `xml:"Recipients"`
	// The Sender element represents the e-mail address for the sender of the message.
	Sender *SenderEmailAddressType `xml:"Sender"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"Subject"`
	// The SubmittedTime element represents the time that the message entered the server.
	SubmittedTime *SubmittedTime `xml:"SubmittedTime"`
}

func (M *MessageTrackingSearchResult) SetForMarshal() {
	M.XMLName.Local = "t:MessageTrackingSearchResult"
}

func (M *MessageTrackingSearchResult) GetSchema() *Schema {
	return &SchemaTypes
}
