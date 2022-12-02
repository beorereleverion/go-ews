package elements

// The ReferenceAttachment element specifies XXX.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/referenceattachment
type ReferenceAttachment struct {
	// The text value of the ChangeKey attribute is the recurring master item's change key. This is a string value.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The text value of the Id attribute is a recurring master item's unique identifier. This is a string value.
	Id *string `xml:"Id,attr"`
}
