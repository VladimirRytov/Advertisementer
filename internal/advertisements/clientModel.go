package advertisements

import (
	"errors"
	"net/mail"
	"regexp"
)

var (
	errClientName         = errors.New("имя должно содержать минимум 2 символа")
	errContactInformation = errors.New("нужно ввести номер телефона или адрес электронной почты")
	errPhonePattern       = errors.New("введеный номер телефона не соотвествует шаблону")
	errEmail              = errors.New("электронная почта введена неверно")

	numberPattern []*regexp.Regexp = []*regexp.Regexp{
		regexp.MustCompile(`(?:\+7|8)(?:\s|-)?\(?\d{3}[)]?[\s|-]?\d{3}[\s|-]?(?:\d{2}[\s|-]?)\d{2}\s?$`),
		regexp.MustCompile(`(?:\+7|8)(?:\s|-)?\(?\d{5}[)]?[\s|-]?\d{1}[\s|-]?(?:\d{2}[\s|-]?)\d{2}\s?$`),
		regexp.MustCompile(`\d{4,10}$`),
	}
)

func ErrClientName() error { return errClientName }

func ErrContactInformation() error { return errContactInformation }

func ErrPhonePattern() error { return errPhonePattern }

func ErrEmail() error { return errEmail }

type Client struct {
	name                  string
	contactNumbers        string
	email                 string
	additionalInformation string
}

func NewClient(name string) (Client, error) {
	client := Client{}
	err := client.SetName(name)
	if err != nil {
		return Client{}, err
	}
	return client, nil
}

func (c *Client) SetName(name string) error {
	if len(name) < 2 {
		return errClientName
	}
	c.name = name
	return nil
}

func (c *Client) SetContactInformation(number, email string) error {
	var err error
	numberLen := len(number)
	emailLen := len(email)
	if numberLen == 0 && emailLen == 0 {
		return errContactInformation
	}
	if numberLen != 0 {
		err = c.setContactNumber(number)
		if err != nil {
			return err
		}
	}
	if emailLen != 0 {
		err = c.setEmail(email)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) setContactNumber(number string) error {
	for _, v := range numberPattern {
		got := v.FindAllString(number, 1)
		if len(got) != 0 {
			c.contactNumbers = got[0]
			return nil
		}
	}
	return errPhonePattern
}

func (c *Client) SetAdditionalInformation(info string) {
	c.additionalInformation = info
}

func (c *Client) setEmail(email string) error {
	e, err := mail.ParseAddress(email)
	if err != nil {
		return errEmail
	}
	c.email = e.Address
	return nil
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) ContactNumber() string {
	return c.contactNumbers
}

func (c *Client) Email() string {
	return c.email
}
func (c *Client) AdditionalInformation() string {
	return c.additionalInformation
}
