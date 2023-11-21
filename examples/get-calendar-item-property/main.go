package main

import (
	"fmt"
	"os"

	goews "github.com/MihaylovNikitos/go-ews"
	"github.com/MihaylovNikitos/go-ews/elements"
	"github.com/MihaylovNikitos/go-ews/operations"
	log "github.com/sirupsen/logrus"
)

var (
	url, user, password, envName string
)

func main() {
	setOSEnvs()
	client := goews.NewClient(url, user, password, goews.Config{
		//TODO true
		Dump:    false,
		NTLM:    true,
		SkipTLS: false,
	})
	gfe, err := operations.NewEnvelopeMarshal(&elements.GetFolder{
		FolderShape: &elements.FolderShape{
			BaseShape: &elements.BaseShape{
				TEXT: elements.BaseShapeAllProperties,
			},
		},
		FolderIds: &elements.FolderIds{
			DistinguishedFolderId: &elements.DistinguishedFolderId{
				Id: getPTR(elements.DistinguishedFolderIdcalendar),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	gfr := &elements.GetFolderResponse{}
	err = client.DoRequest(gfe, gfr)
	if err != nil {
		panic(err)
	}

	fie, err := operations.NewEnvelopeMarshal(&elements.FindItem{
		ItemShape: &elements.ItemShape{
			BaseShape: &elements.BaseShape{
				TEXT: elements.BaseShapeAllProperties,
			},
		},
		ParentFolderIds: &elements.ParentFolderIds{
			DistinguishedFolderId: &elements.DistinguishedFolderId{
				Id: getPTR(elements.DistinguishedFolderIdcalendar),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	eir := &elements.FindItemResponse{}
	err = client.DoRequest(fie, eir)
	if err != nil {
		panic(err)
	}
	var k *elements.CalendarItem
	for _, i := range eir.ResponseMessages.FindItemResponseMessage.RootFolder.Items.CalendarItem {
		if i.Subject.TEXT == envName {
			k = i
		}
	}
	gier, err := operations.NewEnvelopeMarshal(&elements.GetItem{
		ItemShape: &elements.ItemShape{
			BaseShape: &elements.BaseShape{
				TEXT: elements.BaseShapeAllProperties,
			},
		},
		ItemIds: &elements.ItemIds{
			ItemId: &elements.ItemId{
				ChangeKey: k.ItemId.ChangeKey,
				Id:        k.ItemId.Id,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	gierResp := &elements.GetItemResponse{}
	err = client.DoRequest(gier, gierResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", gierResp.ResponseMessages.GetItemResponseMessage.Items.CalendarItem[0].Recurrence.WeeklyRecurrence.DaysOfWeek.TEXT)
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
	envName = os.Getenv("ENVNAME")
	if password == "" {
		log.Fatal("Event name can not be empty")
	}
}

func getPTR[T comparable](t T) *T {
	return &t
}
