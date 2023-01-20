package elements

// The UpdateFolderResponse element defines the response to an UpdateFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatefolderresponse
import "encoding/xml"

type UpdateFolderResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (U *UpdateFolderResponse) SetForMarshal() {
	U.XMLName.Local = "m:UpdateFolderResponse"
}

func (U *UpdateFolderResponse) GetSchema() *Schema {
	return &SchemaMessages
}
