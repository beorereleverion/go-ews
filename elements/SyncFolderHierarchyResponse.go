package elements

// The SyncFolderHierarchyResponse element defines a response to a SyncFolderHierarchy request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncfolderhierarchyresponse
import "encoding/xml"

type SyncFolderHierarchyResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}
