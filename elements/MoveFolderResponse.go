package elements

// The MoveFolderResponse element defines a response to a MoveFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/movefolderresponse
import "encoding/xml"

type MoveFolderResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (M *MoveFolderResponse) SetForMarshal() {
	M.XMLName.Local = "t:MoveFolderResponse"
}

func (M *MoveFolderResponse) GetSchema() *Schema {
	return &SchemaTypes
}
