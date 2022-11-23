package elements

// The FolderShape element identifies the folder properties to include in a GetFolder, FindFolder, or SyncFolderHierarchy response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/foldershape
type FolderShape struct {
	// Identifies the basic configuration of properties to return in a response.
	BaseShape BaseShape `xml:"t:BaseShape"`
	// Identifies additional properties to return in a response.
	AdditionalProperties *AdditionalProperties `xml:"t:AdditionalProperties"`
}
