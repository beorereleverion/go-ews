package elements

// The SIPResponseCode element specifies the SIP response code.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sipresponsecode
import "encoding/xml"

type SIPResponseCode struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (S *SIPResponseCode) SetForMarshal() {
	S.XMLName.Local = "t:SIPResponseCode"
}

func (S *SIPResponseCode) GetSchema() *Schema {
	return &SchemaTypes
}
