package elements

// The CreateFolderPathResponse element is used to return a folder path.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createfolderpathresponse
import "encoding/xml"

type CreateFolderPathResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (C *CreateFolderPathResponse) SetForMarshal() {
	C.XMLName.Local = "m:CreateFolderPathResponse"
}

func (C *CreateFolderPathResponse) GetSchema() *Schema {
	return &SchemaMessages
}
