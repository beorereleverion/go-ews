package elements

// Element information for ReferenceAttachmentType complexType
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/referenceattachmenttype-complextype-ews
import "encoding/xml"

type ReferenceAttachmentTypecomplexTypeEWS struct {
	XMLName xml.Name

	// The AttachLongPathName element is intended for internal use only.
	AttachLongPathName *AttachLongPathName `xml:"AttachLongPathName"`
}
