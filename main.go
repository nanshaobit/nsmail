// Package main
// DateTime: 2023-03-01 18:54
// Author: CN
// Mail: Nanshao@n-s.fun
// Description:

package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/emersion/go-smtp"
)

// The Backend implements SMTP server methods.
type Backend struct{}

func (bkd *Backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

// A Session is returned after EHLO.
type Session struct{}

func (s *Session) AuthPlain(username, password string) error {
	if username != "username" || password != "password" {
		return errors.New("Invalid username or password")
	}
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	fmt.Printf("\n%#v\n\n", opts)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		log.Println("-------------------------")
		f := fmt.Sprintf("%s.eml", time.Now())
		err = os.WriteFile(f, b, 777)
		if err != nil {
			log.Println("ERROR --------------")
			return err
		}
		log.Println("Data:", string(b))
		log.Println("-------------------------")
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func main() {
	be := &Backend{}

	s := smtp.NewServer(be)
	// TODO: read config from file.
	s.Addr = ":25"
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
