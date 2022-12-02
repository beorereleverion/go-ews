package elements

// The App element contains information about an XML manifest file for a mail app that is installed in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/app
type App struct {
	// The Manifest element contains the base64-encoded app manifest file.
	Manifest *Manifest `xml:"Manifest"`
	// The Metadata element contains metadata about the mail app.
	Metadata *Metadata `xml:"t:Metadata"`
}
