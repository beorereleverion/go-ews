# GO-EWS

golang library for interacion with [EWS Exchange Web Service](https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/web-services-reference-for-exchange) 


## Elements

All elements from [EWS elements](https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ews-xml-elements-in-exchange) has been described and are  **accessible!**
*(to the ReadMe update time=, and there are several shortcomings))*

## Operations

Some operations from [EWS operations](https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ews-operations-in-exchange) has been described and are  **accessible!**.  100% of fields are mapped.

*But you can use all of operation, which you need by the interaction with elements. [Example](https://github.com/beorereleverion/go-ews/blob/main/examples/getCalendarItemProperty/main.go)*

### Described operations

 - GetFolder
 - FindItem
 - FindPeople
 - CreteItem

## Usage

some usable examples you can find in [Examples](https://github.com/beorereleverion/go-ews/blob/main/examples/) folder
in this example you can create draft with operation(in examples you can find how to do anything without predefined operation)

    package main

    import (
        "fmt"
        "os"

        goews "github.com/beorereleverion/go-ews"
        "github.com/beorereleverion/go-ews/elements"
        log "github.com/sirupsen/logrus"
    )

    var (
        url, user, password string
    )

    func main() {
        setOSEnvs()
        client := goews.NewClient(url, user, password, goews.Config{
            Dump:    true,
            NTLM:    true,
            SkipTLS: false,
        })
        createItemResponse, err := client.CreateItem(&elements.CreateItem{
            MessageDisposition: getPTR("SendAndSaveCopy"),
            Items: &elements.ItemsNonEmptyArrayOfAllItemsType{
                Message: &elements.Message{
                    
                    ItemClass: &elements.ItemClass{
                        TEXT: "IPM.Note",
                    },
                    Subject: &elements.Subject{
                        TEXT: "Project Action",
                    },
                    Body: &elements.Body{
                        BodyType: getPTR("Text"),
                        TEXT:     "Priority - Update specification",
                    },
                    ToRecipients: &elements.ToRecipients{
                        Mailbox: &elements.Mailbox{
                            EmailAddress: &elements.EmailAddressNonEmptyStringType{
                                TEXT: "sschmidt@example.com",
                            },
                        },
                    },
                    IsRead: &elements.IsRead{
                        TEXT: false,
                    },
                },
            },
            SavedItemFolderId: &elements.SavedItemFolderId{DistinguishedFolderId: &elements.DistinguishedFolderId{
                Id: getPTR(elements.DistinguishedFolderIddrafts),
            }},
        })
        if err != nil {
            panic(err)
        }
        fmt.Printf("%#v\n", createItemResponse)
    }

    func setOSEnvs() {
        url = os.Getenv("URL")
        if url == "" {
            log.Fatal("url can not be empty")
        }
        user = os.Getenv("USER")
        if user == "" {
            log.Fatal("user can not be empty")
        }
        password = os.Getenv("PASSWORD")
        if password == "" {
            log.Fatal("password can not be empty")
        }
    }

    func getPTR[T comparable](t T) *T {
        return &t
    }

