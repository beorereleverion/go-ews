package elements

// The InternetMessageHeaders element contains a collection of some of the Internet message headers that are contained in an item in a mailbox. To get the entire collection of Internet message headers, use the PR_TRANSPORT_MESSAGE_HEADERS property. For more information about EWS and Internet message headers, seeGetting Internet message headersin EWS, MIME, and the missing Internet message headers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/internetmessageheaders
type InternetMessageHeaders struct {
	// The InternetMessageHeader element represents the Internet message header for a given header within the headers collection. To get the entire collection of Internet message headers, use the PR_TRANSPORT_MESSAGE_HEADERS property. For more information about EWS and Internet message headers, seeGetting Internet message headers in EWS, MIME, and the missing Internet message headers.
	InternetMessageHeader *InternetMessageHeader `xml:"t:InternetMessageHeader"`
}
