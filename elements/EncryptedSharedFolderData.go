package elements

// The EncryptedSharedFolderData element contains the encrypted data that a client can use to authorize the sharing of its calendar or contact data with other clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/encryptedsharedfolderdata
type EncryptedSharedFolderData struct {
	// The Data element contains encrypted data that represents the shared data.
	Data *Data `xml:"t:Data"`
	// The Token element contains encrypted data that represents the identification token for the shared data.
	Token *Token `xml:"t:Token"`
}
