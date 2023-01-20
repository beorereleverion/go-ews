package elements

// The Changes element contains a sequence array of change types that represent the types of differences between the items on the client and the items on the Exchange server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/changes-items
import "encoding/xml"

type ChangesItems struct {
	XMLName xml.Name

	// The Create element identifies a single item to create in the local client store.
	Create *CreateItemSync `xml:"Create"`
	// The Delete element identifies a single item to delete in the local client store.
	Delete *DeleteItemSync `xml:"Delete"`
	// The ReadFlagChange element is returned in SyncFolderItems operation responses when an item has been read. This property is read-only.
	ReadFlagChange *ReadFlagChange `xml:"ReadFlagChange"`
	// The Update element identifies a single item to update in the local client store.
	Update *UpdateItemSync `xml:"Update"`
}

func (C *ChangesItems) SetForMarshal() {
	C.XMLName.Local = "m:Changes"
}

func (C *ChangesItems) GetSchema() *Schema {
	return &SchemaMessages
}
