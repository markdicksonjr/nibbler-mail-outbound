package main

import (
	"github.com/markdicksonjr/nibbler"
	"github.com/markdicksonjr/nibbler-mail-outbound/sendgrid"
	"log"
)

func main() {

	// allocate configuration (from env vars, files, etc)
	config, err := nibbler.LoadConfiguration()

	if err != nil {
		log.Fatal(err)
	}

	// allocate the sendgrid extension
	sendgridExtension := sendgrid.Extension{}

	// initialize the application, provide config, logger, extensions
	app := nibbler.Application{}
	if err = app.Init(config, nibbler.DefaultLogger{}, []nibbler.Extension{
		&sendgridExtension,
	}); err != nil {
		log.Fatal(err.Error())
	}

	var toList []*nibbler.EmailAddress
	toList = append(toList, &nibbler.EmailAddress{Address: "mark@example.com", Name: "MD"})

	response, err := sendgridExtension.SendMail(
		&nibbler.EmailAddress{Address: "test@example.com", Name: "Example User"},
		"test email",
		toList,
		"test plain",
		"<strong>test</strong> plain",
	)

	log.Println(response)

	if err != nil {
		log.Fatal(err.Error())
	}

	// start the app
	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
