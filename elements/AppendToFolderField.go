package elements

// The AppendToFolderField element is not implemented. Any request that uses this element will always return an error response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/appendtofolderfield
import "encoding/xml"

type AppendToFolderField struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (A *AppendToFolderField) SetForMarshal() {
	A.XMLName.Local = "t:AppendToFolderField"
}

func (A *AppendToFolderField) GetSchema() *Schema {
	return &SchemaTypes
}
