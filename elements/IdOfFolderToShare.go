package elements

// The IdOfFolderToShare element represents the identifier of the folder on the server that will be shared.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/idoffoldertoshare
type IdOfFolderToShare struct {
	// Contains a string that identifies a version of a folder that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of a folder is used.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Contains a string that identifies a folder in the Exchange store. This attribute is required.
	Id *string `xml:"Id,attr"`
}
