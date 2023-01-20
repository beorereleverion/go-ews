package elements

// The DeleteFolderResponse element defines a response to a DeleteFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletefolderresponse
import "encoding/xml"

type DeleteFolderResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (D *DeleteFolderResponse) SetForMarshal() {
	D.XMLName.Local = "m:DeleteFolderResponse"
}

func (D *DeleteFolderResponse) GetSchema() *Schema {
	return &SchemaMessages
}
