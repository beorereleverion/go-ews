package elements

// The GetStreamingEventsResponse element represents a response to a GetStreamingEvents element request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getstreamingeventsresponse
import "encoding/xml"

type GetStreamingEventsResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (G *GetStreamingEventsResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetStreamingEventsResponse"
}

func (G *GetStreamingEventsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
