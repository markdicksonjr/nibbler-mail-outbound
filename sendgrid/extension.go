package sendgrid

import (
	"errors"
	"github.com/markdicksonjr/nibbler"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Extension struct {
	nibbler.NoOpExtension
	nibbler.MailSender

	apiKey      string
	initialized bool
}

func (s *Extension) Init(app *nibbler.Application) error {
	if app.Config == nil || app.Config.Raw == nil {
		return errors.New("sendgrid extension could not get app config")
	}

	s.apiKey = app.Config.Raw.Get("sendgrid", "api", "key").String("")

	if len(s.apiKey) == 0 {
		return errors.New("sendgrid extension could not get API key")
	}

	s.initialized = true
	return nil
}

func (s *Extension) SendMail(from *nibbler.EmailAddress, subject string, to []*nibbler.EmailAddress, plainTextContent string, htmlContent string) (*nibbler.MailSendResponse, error) {
	if !s.initialized {
		return nil, errors.New("send grid extension used for sending without initialization")
	}

	if from == nil || len((*from).Address) == 0 {
		return nil, errors.New("send grid extension requires 'from' field")
	}

	if len(to) == 0 || to[0] == nil || len((*to[0]).Address) == 0 {
		return nil, errors.New("send grid extension requires at least one recipient")
	}

	fromSg := mail.Email{
		Name:    (*from).Name,
		Address: (*from).Address,
	}

	toSg := mail.Email{
		Name:    (*to[0]).Name,
		Address: (*to[0]).Address,
	}

	var toList []*mail.Email
	for i, v := range to {
		if i > 0 {
			toList = append(toList, &mail.Email{
				Name:    (*v).Name,
				Address: (*v).Address,
			})
		}
	}

	message := mail.NewSingleEmail(&fromSg, subject, &toSg, plainTextContent, htmlContent)

	if len(to) > 1 {
		message.Personalizations = append(message.Personalizations, &mail.Personalization{
			To: toList,
		})
	}
	client := sendgrid.NewSendClient(s.apiKey)
	res, err := client.Send(message)

	if res != nil {
		return &nibbler.MailSendResponse{
			Body:       res.Body,
			Headers:    res.Headers,
			StatusCode: res.StatusCode,
		}, err
	}

	return nil, err
}

func (s *Extension) GetName() string {
	return "sendgrid"
}