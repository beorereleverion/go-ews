package elements

// The ProxyPartnerToken element is used by HTTP proxy of the computer that is running Microsoft Exchange Server 2010 that has the Client Access server role installed. This element is not used by Exchange Web Services (EWS) operations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proxypartnertoken
import "encoding/xml"

type ProxyPartnerToken struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (P *ProxyPartnerToken) SetForMarshal() {
	P.XMLName.Local = "t:ProxyPartnerToken"
}

func (P *ProxyPartnerToken) GetSchema() *Schema {
	return &SchemaTypes
}
