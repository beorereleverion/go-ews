package elements

// The ExportItemsResponse element represents response to a single ExportItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exportitemsresponse
import "encoding/xml"

type ExportItemsResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (E *ExportItemsResponse) SetForMarshal() {
	E.XMLName.Local = "m:ExportItemsResponse"
}

func (E *ExportItemsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
