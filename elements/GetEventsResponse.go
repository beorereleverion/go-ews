package elements

// The GetEventsResponse element represents a response to a GetEvents request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/geteventsresponse
import "encoding/xml"

type GetEventsResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (G *GetEventsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetEventsResponse"
}

func (G *GetEventsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
