package elements

// The Changes element contains a sequenced array of change types that represent the type of differences between the folders on the client and the folders on the computer that is running Microsoft Exchange Server 2007.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/changes-hierarchy
import "encoding/xml"

type ChangesHierarchy struct {
	XMLName xml.Name

	// The Create element identifies a single folder to create in the local client store.
	Create *CreateFolderSync `xml:"Create"`
	// The Delete element identifies a single folder to delete in the local client store.
	Delete *DeleteFolderSync `xml:"Delete"`
	// The Update element identifies a single folder to update in the local client store.
	Update *UpdateFolderSync `xml:"Update"`
}

func (C *ChangesHierarchy) SetForMarshal() {
	C.XMLName.Local = "m:Changes"
}

func (C *ChangesHierarchy) GetSchema() *Schema {
	return &SchemaMessages
}
