package elements

// The DisableApp element specifies a request to disable an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/disableapp
type DisableApp struct {
	// The DisableReason element specifies the reason for disabling an app.
	DisableReason *DisableReason `xml:"t:DisableReason"`
	// The ID element specifies the identifier of an app.
	ID *IDString `xml:"m:ID"`
}
