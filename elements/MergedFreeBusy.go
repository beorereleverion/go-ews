package elements

// The MergedFreeBusy element contains the merged free/busy stream of data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mergedfreebusy
import "encoding/xml"

type MergedFreeBusy struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MergedFreeBusy) SetForMarshal() {
	M.XMLName.Local = "t:MergedFreeBusy"
}

func (M *MergedFreeBusy) GetSchema() *Schema {
	return &SchemaTypes
}
