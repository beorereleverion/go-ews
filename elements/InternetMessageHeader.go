package elements

// The InternetMessageHeader element represents the Internet message header for a given header within the headers collection. To get the entire collection of Internet message headers, use the PR_TRANSPORT_MESSAGE_HEADERS property. For more information about EWS and Internet message headers, seeGetting Internet message headers in EWS, MIME, and the missing Internet message headers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/internetmessageheader
type InternetMessageHeader struct {
	// Identifies the header name.
	HeaderName *string `xml:"HeaderName,attr"`
}
