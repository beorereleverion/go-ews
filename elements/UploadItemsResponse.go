package elements

// The UploadItemsResponse element represents a response to a single UploadItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uploaditemsresponse
import "encoding/xml"

type UploadItemsResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (U *UploadItemsResponse) SetForMarshal() {
	U.XMLName.Local = "m:UploadItemsResponse"
}

func (U *UploadItemsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
