package elements

// The MarkAsJunkResponse element specifies the response to a MarkAsJunk request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markasjunkresponse
import "encoding/xml"

type MarkAsJunkResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (M *MarkAsJunkResponse) SetForMarshal() {
	M.XMLName.Local = "m:MarkAsJunkResponse"
}

func (M *MarkAsJunkResponse) GetSchema() *Schema {
	return &SchemaMessages
}
