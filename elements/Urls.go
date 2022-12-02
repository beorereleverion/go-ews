package elements

// The Urls element specifies an array of URLs for a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/urls
type Urls struct {
	// The Url element represents the location of the client Web service for push notifications.
	Url *Url `xml:"Url"`
}
