package goews

import (
	"github.com/MihaylovNikitos/go-ews/elements"
	"github.com/MihaylovNikitos/go-ews/operations"
)

// Find information about the FindItem EWS operation.
// The FindItem operation searches for items that are located in a user's mailbox. This operation provides many ways to filter and format how search results are returned to the caller.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/finditem-operation
func (c *client) FindItem(eItem *elements.FindItem) (*elements.FindItemResponse, error) {
	oem, err := operations.NewEnvelopeMarshal(eItem)
	if err != nil {
		return nil, err
	}
	eResponse := &elements.FindItemResponse{}
	err = c.DoRequest(oem, eResponse)
	if err != nil {
		return nil, err
	}
	return eResponse, nil
}
