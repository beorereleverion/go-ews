package elements

// The PlayOnPhoneEnabled element indicates whether the Play-on-Phone feature is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/playonphoneenabled
import "encoding/xml"

type PlayOnPhoneEnabled struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	PlayOnPhoneEnabledfalse bool = false
	// true
	PlayOnPhoneEnabledtrue bool = true
)

func (P *PlayOnPhoneEnabled) SetForMarshal() {
	P.XMLName.Local = "t:PlayOnPhoneEnabled"
}

func (P *PlayOnPhoneEnabled) GetSchema() *Schema {
	return &SchemaTypes
}
