package elements

// The UninstallApp element specifies a request to uninstall an app by its identifier.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uninstallapp
type UninstallApp struct {
	// The ID element specifies the identifier of an app.
	ID *IDString `xml:"m:ID"`
}
