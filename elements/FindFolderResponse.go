package elements

// The FindFolderResponse element defines a response to a FindFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findfolderresponse
import "encoding/xml"

type FindFolderResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (F *FindFolderResponse) SetForMarshal() {
	F.XMLName.Local = "m:FindFolderResponse"
}

func (F *FindFolderResponse) GetSchema() *Schema {
	return &SchemaMessages
}
