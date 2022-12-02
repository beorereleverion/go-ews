package elements

// The WithinSizeRange element specifies the minimum and maximum sizes that incoming messages must be in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/withinsizerange
type WithinSizeRange struct {
	// The MaximumSize element represents the maximum size that a message must be in order for the condition or exception to apply.
	MaximumSize *MaximumSize `xml:"m:MaximumSize"`
	// The MinimumSize element represents the minimum size that a message must be in order for the condition or exception to apply.
	MinimumSize *MinimumSize `xml:"m:MinimumSize"`
}
