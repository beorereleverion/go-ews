package elements

// The EncryptedSharedFolderDataCollection element contains a collection of data structures that a client can use to authorize the sharing of its calendar or contact data with other clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/encryptedsharedfolderdatacollection
import "encoding/xml"

type EncryptedSharedFolderDataCollection struct {
	XMLName xml.Name

	// The EncryptedSharedFolderData element contains the encrypted data that a client can use to authorize the sharing of its calendar or contact data with other clients.
	EncryptedSharedFolderData *EncryptedSharedFolderData `xml:"EncryptedSharedFolderData"`
}

func (E *EncryptedSharedFolderDataCollection) SetForMarshal() {
	E.XMLName.Local = "m:EncryptedSharedFolderDataCollection"
}

func (E *EncryptedSharedFolderDataCollection) GetSchema() *Schema {
	return &SchemaMessages
}
