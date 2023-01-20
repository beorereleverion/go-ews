package elements

// The ReadFlagChange element is returned in SyncFolderItems operation responses when an item has been read. This property is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/readflagchange
import "encoding/xml"

type ReadFlagChange struct {
	XMLName xml.Name

	// The IsRead element indicates whether a message has been read.
	IsRead *IsRead `xml:"IsRead"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
}

func (R *ReadFlagChange) SetForMarshal() {
	R.XMLName.Local = "t:ReadFlagChange"
}

func (R *ReadFlagChange) GetSchema() *Schema {
	return &SchemaTypes
}
