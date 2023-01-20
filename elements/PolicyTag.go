package elements

// The PolicyTag element specifies the retention identifier on an item or folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/policytag
import "encoding/xml"

type PolicyTag struct {
	XMLName xml.Name

	// Indicates whether a policy tag was explicitly set on an item or folder.   A text value of true for the IsExplicit attribute indicates that the policy tag was explicitly set on the item or folder. A text value of false indicates that the policy tag was implicitly set on the item or folder based on the parent folder policy tag.
	IsExplicit *string `xml:"IsExplicit,attr"`
}

func (P *PolicyTag) SetForMarshal() {
	P.XMLName.Local = "t:PolicyTag"
}

func (P *PolicyTag) GetSchema() *Schema {
	return &SchemaTypes
}
