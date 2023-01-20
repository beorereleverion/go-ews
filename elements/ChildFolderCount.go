package elements

// The ChildFolderCount element represents the number of immediate child folders that are contained within a folder. This property is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/childfoldercount
import "encoding/xml"

type ChildFolderCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (C *ChildFolderCount) SetForMarshal() {
	C.XMLName.Local = "t:ChildFolderCount"
}

func (C *ChildFolderCount) GetSchema() *Schema {
	return &SchemaTypes
}
