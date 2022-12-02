package elements

// The ConvertId element defines a request to convert item and folder identifiers between supported Exchange formats. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/convertid
type ConvertId struct {
	// The SourceIds element contains the source identifiers to convert.
	SourceIds *SourceIds `xml:"m:SourceIds"`
	// Describes the identifier format that will be returned for all the converted identifiers. The DestinationFormat is described by the IdFormatType.
	DestinationFormat *string `xml:"DestinationFormat,attr"`
}
