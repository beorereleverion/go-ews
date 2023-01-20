package elements

// The IsManagedFoldersRoot element indicates whether the managed folder is the root for all managed folders.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismanagedfoldersroot
import "encoding/xml"

type IsManagedFoldersRoot struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsManagedFoldersRoot) SetForMarshal() {
	I.XMLName.Local = "t:IsManagedFoldersRoot"
}

func (I *IsManagedFoldersRoot) GetSchema() *Schema {
	return &SchemaTypes
}
