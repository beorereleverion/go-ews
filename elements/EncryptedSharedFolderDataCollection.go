package elements

// The EncryptedSharedFolderDataCollection element contains a collection of data structures that a client can use to authorize the sharing of its calendar or contact data with other clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/encryptedsharedfolderdatacollection
type EncryptedSharedFolderDataCollection struct {
	// The EncryptedSharedFolderData element contains the encrypted data that a client can use to authorize the sharing of its calendar or contact data with other clients.
	EncryptedSharedFolderData *EncryptedSharedFolderData `xml:"t:EncryptedSharedFolderData"`
}
