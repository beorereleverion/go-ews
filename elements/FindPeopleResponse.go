package elements

// The FindPeopleResponse element specifies the response to a FindPeople request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findpeopleresponse
import "encoding/xml"

type FindPeopleResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The People element specifies an array of persona data returned as the result of a FindPeople request.
	People *People `xml:"People"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// The TotalNumberOfPeopleInView element specifies the total number of personas returned in a FindPeople response.
	TotalNumberOfPeopleInView *TotalNumberOfPeopleInView `xml:"TotalNumberOfPeopleInView"`
	// Specifies the response class.
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (F *FindPeopleResponse) SetForMarshal() {
	F.XMLName.Local = "m:FindPeopleResponse"
}

func (F *FindPeopleResponse) GetSchema() *Schema {
	return &SchemaMessages
}
