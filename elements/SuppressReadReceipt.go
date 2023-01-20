package elements

// The SuppressReadReceipt element is used to suppress read receipts.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suppressreadreceipt
import "encoding/xml"

type SuppressReadReceipt struct {
	XMLName xml.Name

	// The ReferenceItemId element identifies the item to which the response object refers.
	ReferenceItemId *ReferenceItemId `xml:"ReferenceItemId"`
}

func (S *SuppressReadReceipt) SetForMarshal() {
	S.XMLName.Local = "t:SuppressReadReceipt"
}

func (S *SuppressReadReceipt) GetSchema() *Schema {
	return &SchemaTypes
}
