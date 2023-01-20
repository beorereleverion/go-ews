package elements

// The EncryptedSharedFolderData element contains the encrypted data that a client can use to authorize the sharing of its calendar or contact data with other clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/encryptedsharedfolderdata
import "encoding/xml"

type EncryptedSharedFolderData struct {
	XMLName xml.Name

	// The Data element contains encrypted data that represents the shared data.
	Data *Data `xml:"Data"`
	// The Token element contains encrypted data that represents the identification token for the shared data.
	Token *Token `xml:"Token"`
}

func (E *EncryptedSharedFolderData) SetForMarshal() {
	E.XMLName.Local = "t:EncryptedSharedFolderData"
}

func (E *EncryptedSharedFolderData) GetSchema() *Schema {
	return &SchemaTypes
}
