package main

import (
	"github.com/markdicksonjr/nibbler"
	"github.com/markdicksonjr/nibbler-mail-outbound/mandrill"
	"log"
)

func main() {

	// allocate configuration (from env vars, files, etc)
	config, err := nibbler.LoadConfiguration(nil)

	if err != nil {
		log.Fatal(err)
	}

	// allocate the sparkpost extension
	mandrillExtension := mandrill.Extension{}

	// initialize the application, provide config, logger, extensions
	app := nibbler.Application{}
	if err = app.Init(config, nibbler.DefaultLogger{}, []nibbler.Extension{
		&mandrillExtension,
	}); err != nil {
		log.Fatal(err.Error())
	}

	var toList []*nibbler.EmailAddress
	toList = append(toList, &nibbler.EmailAddress{Address: "mark@example.com", Name: "MD"})

	_, err = mandrillExtension.SendMail(
		&nibbler.EmailAddress{Address: "test@example.com", Name: "Example User"},
		"test email",
		toList,
		"test plain",
		"<strong>test</strong> plain",
	)
}
