package main

import (
	"fmt"
	"os"

	goews "github.com/MihaylovNikitos/go-ews"
	"github.com/MihaylovNikitos/go-ews/elements"
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
				ToRecipients: []*elements.ToRecipients{{
					Mailbox: &elements.Mailbox{
						EmailAddress: &elements.EmailAddressNonEmptyStringType{
							TEXT: "sschmidt@example.com",
						},
					},
				}},
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
	fmt.Printf("%#v\n", createItemResponse.ResponseMessages)
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
