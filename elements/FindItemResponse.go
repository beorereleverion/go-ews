package elements

// The FindItemResponse element defines a response to a FindItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/finditemresponse
import "encoding/xml"

type FindItemResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (F *FindItemResponse) SetForMarshal() {
	F.XMLName.Local = "m:FindItemResponse"
}

func (F *FindItemResponse) GetSchema() *Schema {
	return &SchemaMessages
}
