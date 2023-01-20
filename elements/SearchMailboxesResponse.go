package elements

// The SearchMailboxesResponse element contains the response to a SearchMailboxes WSDL operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchmailboxesresponse
import "encoding/xml"

type SearchMailboxesResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (S *SearchMailboxesResponse) SetForMarshal() {
	S.XMLName.Local = "m:SearchMailboxesResponse"
}

func (S *SearchMailboxesResponse) GetSchema() *Schema {
	return &SchemaMessages
}
