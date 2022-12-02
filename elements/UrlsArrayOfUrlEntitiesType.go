package elements

// The Urls element specifies an array of URLs that are the result of entity extraction from an item in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/urls-arrayofurlentitiestype
type UrlsArrayOfUrlEntitiesType struct {
	// The UrlEntity element identifies a single extracted URL entity.
	UrlEntity *UrlEntity `xml:"t:UrlEntity"`
}
