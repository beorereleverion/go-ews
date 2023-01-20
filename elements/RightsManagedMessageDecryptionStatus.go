package elements

// The RightsManagedMessageDecryptionStatus element specifies the rights management decryption status of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rightsmanagedmessagedecryptionstatus
import "encoding/xml"

type RightsManagedMessageDecryptionStatus struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (R *RightsManagedMessageDecryptionStatus) SetForMarshal() {
	R.XMLName.Local = "t:RightsManagedMessageDecryptionStatus"
}

func (R *RightsManagedMessageDecryptionStatus) GetSchema() *Schema {
	return &SchemaTypes
}
