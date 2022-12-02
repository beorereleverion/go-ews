package elements

// The InstallApp element specifies the request to install an app.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/installapp
type InstallApp struct {
	// The Manifest element contains the base64-encoded app manifest file.
	Manifest *Manifest `xml:"Manifest"`
}
