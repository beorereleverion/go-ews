package main

import (
	"fmt"
	"os"

	goews "github.com/beorereleverion/go-ews"
	"github.com/beorereleverion/go-ews/elements"
	"github.com/beorereleverion/go-ews/operations"
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
	gfe, err := operations.NewGetFolderEnvelope()
	if err != nil {
		panic(err)
	}
	gfe.Body = &operations.GetFolderBody{
		GetFolder: operations.GetFolder{
			FolderShape: &elements.FolderShape{
				BaseShape: elements.BaseShapeAllProperties,
			},
			FolderIds: &elements.FolderIds{
				DistinguishedFolderId: &elements.DistinguishedFolderId{
					ID: elements.DistinguishedFolderIdAttrIDcalendar,
				},
			},
		},
	}
	bansw, err := client.SendAndReceive(gfe)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bansw))
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
