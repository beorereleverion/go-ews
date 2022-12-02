package elements

// The Apps element contains information about all the XML manifest files for apps installed in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/apps
type Apps struct {
	// The App element contains information about an XML manifest file for a mail app that is installed in a mailbox.
	App *App `xml:"t:App"`
}
