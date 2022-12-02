package elements

// The Bodies element specifies an array of BodyContentAttributedValue elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bodies
type Bodies struct {
	// The BodyContentAttributedValue element specifies the body content of an item.
	BodyContentAttributedValue *BodyContentAttributedValue `xml:"t:BodyContentAttributedValue"`
}
