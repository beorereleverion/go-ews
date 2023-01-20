package elements

// The SharedFolderId element represents the identifier of the shared folder the local folder identifier for which should be returned by a GetSharingFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sharedfolderid
import "encoding/xml"

type SharedFolderId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SharedFolderId) SetForMarshal() {
	S.XMLName.Local = "m:SharedFolderId"
}

func (S *SharedFolderId) GetSchema() *Schema {
	return &SchemaMessages
}
