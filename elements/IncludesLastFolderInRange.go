package elements

// The IncludesLastFolderInRange element indicates whether the last item to synchronize has been included in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includeslastfolderinrange
import "encoding/xml"

type IncludesLastFolderInRange struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IncludesLastFolderInRange) SetForMarshal() {
	I.XMLName.Local = "m:IncludesLastFolderInRange"
}

func (I *IncludesLastFolderInRange) GetSchema() *Schema {
	return &SchemaMessages
}
