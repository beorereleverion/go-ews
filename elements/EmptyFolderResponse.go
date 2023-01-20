package elements

// The EmptyFolderResponse element defines a response to an EmptyFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emptyfolderresponse
import "encoding/xml"

type EmptyFolderResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (E *EmptyFolderResponse) SetForMarshal() {
	E.XMLName.Local = "m:EmptyFolderResponse"
}

func (E *EmptyFolderResponse) GetSchema() *Schema {
	return &SchemaMessages
}
