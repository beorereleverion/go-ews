package elements

// The CopyItemResponse element defines a response to a CopyItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/copyitemresponse
import "encoding/xml"

type CopyItemResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (C *CopyItemResponse) SetForMarshal() {
	C.XMLName.Local = "m:CopyItemResponse"
}

func (C *CopyItemResponse) GetSchema() *Schema {
	return &SchemaMessages
}
