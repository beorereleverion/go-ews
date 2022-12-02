package elements

// The Manifests element contains a collection of base64-encoded app manifests that are installed for the email account.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/manifests
type Manifests struct {
	// The Manifest element contains the base64-encoded app manifest file.
	Manifest *Manifest `xml:"Manifest"`
}
