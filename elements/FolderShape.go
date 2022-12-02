package elements

// The FolderShape element identifies the folder properties to include in a GetFolder, FindFolder, or SyncFolderHierarchy response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/foldershape
type FolderShape struct {
	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"t:AdditionalProperties"`
	// The BaseShape element identifies the set of properties to return in an item or folder response.
	BaseShape *BaseShape `xml:"t:BaseShape"`
}
