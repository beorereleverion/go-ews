package elements

// The FolderShape element identifies the folder properties to include in a GetFolder, FindFolder, or SyncFolderHierarchy response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/foldershape
import "encoding/xml"

type FolderShape struct {
	XMLName xml.Name

	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"AdditionalProperties"`
	// The BaseShape element identifies the set of properties to return in an item or folder response.
	BaseShape *BaseShape `xml:"BaseShape"`
}

func (F *FolderShape) SetForMarshal() {
	F.XMLName.Local = "m:FolderShape"
}

func (F *FolderShape) GetSchema() *Schema {
	return &SchemaMessages
}
