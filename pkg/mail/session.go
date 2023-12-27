package mail

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/DusanKasan/parsemail"
	"github.com/emersion/go-smtp"

	nm "net/mail"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

// A Session is returned after successful login.
type Session struct {
	authUser string
	authPass string
	hooks    []Hook
	from     string
	to       string
}

// AuthPlain implements authentication using SASL PLAIN.
func (s Session) AuthPlain(username, password string) error {
	if username != s.authUser || password != s.authPass {
		return errors.New("invalid username or password")
	}
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	for _, h := range s.hooks {
		if h.Address == to {
			s.to = to
			return nil
		}
	}

	log.Println("Error: Unknown recipient:", to)

	return fmt.Errorf("unknown recipient: %s", to)
}

func (s *Session) Data(r io.Reader) error {
	email, err := parsemail.Parse(r) // returns Email struct and error
	if err != nil {
		return err
	}

	for _, h := range s.hooks {
		if h.Address != s.to {
			continue
		}

		if err := h.Send(s.from, formatEmail(email, h)); err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func formatEmail(email parsemail.Email, hook Hook) string {
	if hook.HTMLMarkdown && email.HTMLBody != "" {
		return formatHTMLEmail(email, hook)
	}
	return formatTextEmail(email, hook)
}

func formatHTMLEmail(email parsemail.Email, hook Hook) string {
	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(email.HTMLBody)
	if err != nil {
		log.Printf("error in markdown conversion: %v. Processing text body instead.", err)
		log.Println("Processing text body instead.")
		return formatTextEmail(email, hook)
	}

	return format(email, markdown)
}

func formatTextEmail(email parsemail.Email, hook Hook) string {
	return format(email, email.TextBody)
}

func format(email parsemail.Email, text string) string {
	return fmt.Sprintf("From: **%s**:\nSubject: *%s*\n%s", formatFrom(email.From), email.Subject, text)
}

func formatFrom(addresses []*nm.Address) string {
	if len(addresses) == 0 {
		return "NO FROM ADDRESS"
	}
	var sb strings.Builder
	for _, a := range addresses {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%s <%s>", a.Name, a.Address))
	}
	return sb.String()
}
