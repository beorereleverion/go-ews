package elements

// The SuppressReadReceipts element indicates whether read receipts should be suppressed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suppressreadreceipts
import "encoding/xml"

type SuppressReadReceipts struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	SuppressReadReceiptsfalse bool = false
	// true
	SuppressReadReceiptstrue bool = true
)

func (S *SuppressReadReceipts) SetForMarshal() {
	S.XMLName.Local = "m:SuppressReadReceipts"
}

func (S *SuppressReadReceipts) GetSchema() *Schema {
	return &SchemaMessages
}
