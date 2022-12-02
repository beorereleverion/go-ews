package elements

// The Updates element contains a set of elements that define append, set, and delete changes to folder properties.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updates-folder
type UpdatesFolder struct {
	// The AppendToFolderField element is not implemented. Any request that uses this element will always return an error response.
	AppendToFolderField *AppendToFolderField `xml:"t:AppendToFolderField"`
	// The DeleteFolderField element represents an operation to delete a given property from a folder during an UpdateFolder call.
	DeleteFolderField *DeleteFolderField `xml:"t:DeleteFolderField"`
	// The SetFolderField element represents an update that sets the value for a single property on a folder in an UpdateFolder operation.
	SetFolderField *SetFolderField `xml:"t:SetFolderField"`
}
