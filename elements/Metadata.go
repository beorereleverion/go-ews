package elements

// The Metadata element contains metadata about the mail app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/metadata-ex15websvcsotherref
type Metadata struct {
	// The ActionUrl element identifies the URL that the user should navigate to, in order to fix an issue indicated by the AppStatus element.
	ActionUrl *ActionUrl `xml:"t:ActionUrl"`
	// The AppStatus element value indicates the status of the mail app.
	AppStatus *AppStatus `xml:"t:AppStatus"`
	// The EndNodeUrl element specifies the URL for the mail app in the Office Store.
	EndNodeUrl *EndNodeUrl `xml:"t:EndNodeUrl"`
}
